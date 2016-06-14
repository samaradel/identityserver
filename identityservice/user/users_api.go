package user

import (
	"encoding/json"
	"net/http"

	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/itsyouonline/identityserver/communication"
	"github.com/itsyouonline/identityserver/credentials/password"
	"github.com/itsyouonline/identityserver/credentials/totp"
	contractdb "github.com/itsyouonline/identityserver/db/contract"
	"github.com/itsyouonline/identityserver/db/user"
	"github.com/itsyouonline/identityserver/db/user/apikey"
	validationdb "github.com/itsyouonline/identityserver/db/validation"
	"github.com/itsyouonline/identityserver/identityservice/contract"
	"github.com/itsyouonline/identityserver/identityservice/invitations"
	"github.com/itsyouonline/identityserver/validation"
)

type UsersAPI struct {
	SmsService                    communication.SMSService
	PhonenumberValidationService  *validation.IYOPhonenumberValidationService
	EmailService                  communication.EmailService
	EmailAddressValidationService *validation.IYOEmailAddressValidationService
}

func isUniquePhonenumber(user *user.User, number string, label string) (unique bool) {
	unique = true
	for _, phonenumber := range user.Phonenumbers {
		if phonenumber.Label != label && phonenumber.Phonenumber == number {
			unique = false
			return
		}
	}
	return
}

func isLastVerifiedPhoneNumber(user *user.User, number string, label string, r *http.Request) (last bool, err error) {
	last = false
	valMgr := validationdb.NewManager(r)
	validated, err := valMgr.IsPhonenumberValidated(user.Username, string(number))
	if err != nil {
		return
	}
	if validated {
		// check if this phone number is the last verified one
		uniquelabel := isUniquePhonenumber(user, number, label)
		hasotherverifiednumbers := false
		verifiednumbers, err := valMgr.GetByUsernameValidatedPhonenumbers(user.Username)
		if err != nil {
			return false, err

		}
		for _, verifiednumber := range verifiednumbers {
			if verifiednumber.Phonenumber != string(number) {
				hasotherverifiednumbers = true
				break

			}
		}
		if uniquelabel && !hasotherverifiednumbers {
			return true, nil
		}

	}
	return
}

// It is handler for POST /users
func (api UsersAPI) Post(w http.ResponseWriter, r *http.Request) {

	var u user.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)
	if err := userMgr.Save(&u); err != nil {
		log.Error("ERROR while saving user:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&u)
}

// It is handler for GET /users/{username}
func (api UsersAPI) usernameGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	userMgr := user.NewManager(r)

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func isValidLabel(label string) (valid bool) {
	valid = true
	labelLength := len(label)
	valid = valid && labelLength > 2 && labelLength < 51

	if !valid {
		log.Debug("Invalid label: ", label)
	}
	return valid
}

// RegisterNewEmailAddress is the handler for POST /users/{username}/emailaddresses
// Register a new email address
func (api UsersAPI) RegisterNewEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	body := user.EmailAddress{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(body.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)
	u, err := userMgr.GetByName(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if _, err := u.GetEmailAddressByLabel(body.Label); err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err = userMgr.SaveEmail(username, body); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	valMgr := validationdb.NewManager(r)
	validated, err := valMgr.IsEmailAddressValidated(username, body.EmailAddress)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !validated {
		_, err = api.EmailAddressValidationService.RequestValidation(r, username, body.EmailAddress, fmt.Sprintf("https://%s/emailvalidation", r.Host))
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

// UpdateEmailAddress is the handler for PUT /users/{username}/emailaddresses/{label}
// Updates the label and/or value of an email address
func (api UsersAPI) UpdateEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]

	body := user.EmailAddress{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(body.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)

	if oldlabel != body.Label {
		u, err := userMgr.GetByName(username)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if _, err := u.GetEmailAddressByLabel(body.Label); err != nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	if err := userMgr.SaveEmail(username, body); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != body.Label {
		if err := userMgr.RemoveEmail(username, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	valMgr := validationdb.NewManager(r)
	validated, err := valMgr.IsEmailAddressValidated(username, body.EmailAddress)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !validated {
		_, err = api.EmailAddressValidationService.RequestValidation(r, username, body.EmailAddress, fmt.Sprintf("https://%s/emailvalidation", r.Host))
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

// Validate email address is the handler for GET /users/{username}/emailaddress/{label}/validate
func (api UsersAPI) ValidateEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	email, err := userobj.GetEmailAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = api.EmailAddressValidationService.RequestValidation(r, username, email.EmailAddress, fmt.Sprintf("https://%s/emailvalidation", r.Host))
	w.WriteHeader(http.StatusNoContent)
}

// ListEmailAddresses is the handler for GET /users/{username}/emailaddresses
func (api UsersAPI) ListEmailAddresses(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	validated := strings.Contains(r.URL.RawQuery, "validated")
	userMgr := user.NewManager(r)

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	emails := user.EmailAddresses
	if validated {
		valMngr := validationdb.NewManager(r)
		validatedemails, err := valMngr.GetByUsernameValidatedEmailAddress(username)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		for index, email := range emails {
			found := false
			for _, validatedemail := range validatedemails {
				if email.EmailAddress == validatedemail.EmailAddress {
					found = true
					break
				}
			}
			if !found {
				emails = append(emails[:index], emails[index+1:]...)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

// DeleteEmailAddress is the handler for DELETE /users/{username}/emailaddresses/{label}
// Removes an email address
func (api UsersAPI) DeleteEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]

	userMgr := user.NewManager(r)
	valMgr := validationdb.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	email, err := u.GetEmailAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if len(u.EmailAddresses) == 1 {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err = userMgr.RemoveEmail(username, label); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err = valMgr.RemoveValidatedEmailAddress(username, email.EmailAddress); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeleteGithubAccount is the handler for DELETE /users/{username}/github
// Delete the associated Github account.
func (api UsersAPI) DeleteGithubAccount(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)
	err := userMgr.DeleteGithubAccount(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete FacebookAccount is the handler for DELETE /users/{username}/facebook
// Delete the associated facebook account
func (api UsersAPI) DeleteFacebookAccount(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	userMgr := user.NewManager(r)
	err := userMgr.DeleteFacebookAccount(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdatePassword handler
func (api UsersAPI) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	body := struct {
		Currentpassword string `json:"currentpassword"`
		Newpassword     string `json:"newpassword"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userMgr := user.NewManager(r)
	exists, err := userMgr.Exists(username)
	if !exists || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	passwordMgr := password.NewManager(r)
	passwordok, err := passwordMgr.Validate(username, body.Currentpassword)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !passwordok {
		writeErrorResponse(w, 422, "incorrect_password")
		return
	}
	err = passwordMgr.Save(username, body.Newpassword)
	if err != nil {
		writeErrorResponse(w, 422, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetUserInformation is the handler for GET /users/{username}/info
func (api UsersAPI) GetUserInformation(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	requestingClient, validClient := context.Get(r, "client_id").(string)
	if !validClient {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	authorization, err := userMgr.GetAuthorization(username, requestingClient)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	respBody := &Userview{
		Username: userobj.Username,
	}

	if authorization.Name {
		respBody.Firstname = userobj.Firstname
		respBody.Lastname = userobj.Lastname
	}
	if authorization.Github {
		respBody.Github = userobj.Github.Name
	}
	if authorization.Facebook {
		respBody.Facebook = userobj.Facebook.Name
	}
	if authorization.Address != nil {
		respBody.Addresses = make([]user.Address, 0)

		for requestedLabel, realLabel := range authorization.Address {
			address, err := userobj.GetAddressByLabel(realLabel)
			if err != nil {
				newaddress := user.Address{
					Label:      requestedLabel,
					City:       address.City,
					Country:    address.Country,
					Nr:         address.Nr,
					Other:      address.Other,
					Postalcode: address.Postalcode,
					Street:     address.Street,
				}
				respBody.Addresses = append(respBody.Addresses, newaddress)
			}
		}
	}

	if authorization.Email != nil {
		respBody.EmailAddresses = make([]user.EmailAddress, 0)

		for requestedLabel, realLabel := range authorization.Email {
			email, err := userobj.GetEmailAddressByLabel(realLabel)
			if err != nil {
				newemail := user.EmailAddress{
					Label:        requestedLabel,
					EmailAddress: email.EmailAddress,
				}
				respBody.EmailAddresses = append(respBody.EmailAddresses, newemail)
			}
		}
	}

	if authorization.Phone != nil {
		respBody.Phonenumbers = make([]user.Phonenumber, 0)
		for requestedLabel, realLabel := range authorization.Phone {
			phonenumber, err := userobj.GetPhonenumberByLabel(realLabel)
			if err != nil {
				newnumber := user.Phonenumber{
					Label:       requestedLabel,
					Phonenumber: phonenumber.Phonenumber,
				}
				respBody.Phonenumbers = append(respBody.Phonenumbers, newnumber)
			}
		}
	}

	if authorization.Bank != nil {
		respBody.BankAccounts = make([]user.BankAccount, 0)

		for requestedLabel, realLabel := range authorization.Bank {
			bank, err := userobj.GetBankAccountByLabel(realLabel)
			if err != nil {
				newbank := user.BankAccount{
					Label:   requestedLabel,
					Bic:     bank.Bic,
					Country: bank.Country,
					Iban:    bank.Iban,
				}
				respBody.BankAccounts = append(respBody.BankAccounts, newbank)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

// usernamevalidateGet is the handler for GET /users/{username}/validate
func (api UsersAPI) usernamevalidateGet(w http.ResponseWriter, r *http.Request) {

	// token := req.FormValue("token")

	// uncomment below line to add header
	// w.Header.Set("key","value")
}

// RegisterNewPhonenumber is the handler for POST /users/{username}/phonenumbers
// Register a new phonenumber
func (api UsersAPI) RegisterNewPhonenumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	body := user.Phonenumber{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(body.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !body.IsValid() {
		log.Debug("Invalid phonenumber: ", body.Phonenumber)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Check if this label is already used
	_, err = u.GetPhonenumberByLabel(body.Label)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err := userMgr.SavePhone(username, body); err != nil {
		log.Error("ERROR while saving a phonenumber - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// respond with created phone number.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(body)
}

// usernamephonenumbersGet is the handler for GET /users/{username}/phonenumbers
func (api UsersAPI) usernamephonenumbersGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	validated := strings.Contains(r.URL.RawQuery, "validated")
	userMgr := user.NewManager(r)

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	phonenumbers := user.Phonenumbers
	if validated {
		valMngr := validationdb.NewManager(r)
		validatednumbers, err := valMngr.GetByUsernameValidatedPhonenumbers(username)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		for index, number := range phonenumbers {
			found := false
			for _, validatednumber := range validatednumbers {
				if number.Phonenumber == validatednumber.Phonenumber {
					found = true
					break
				}
			}
			if !found {
				phonenumbers = append(phonenumbers[:index], phonenumbers[index+1:]...)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phonenumbers)
}

// usernamephonenumberslabelGet is the handler for GET /users/{username}/phonenumbers/{label}
func (api UsersAPI) usernamephonenumberslabelGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	phonenumber, err := userobj.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phonenumber)
}

// Validate phone number is the handler for POST /users/{username}/phonenumbers/{label}/validate
func (api UsersAPI) ValidatePhoneNumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	phonenumber, err := userobj.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	validationKey := ""
	validationKey, err = api.PhonenumberValidationService.RequestValidation(r, username, phonenumber, fmt.Sprintf("https://%s/phonevalidation", r.Host))
	response := struct {
		ValidationKey string `json:"validationkey"`
	}{
		validationKey,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

// Validate phone number is the handler for PUT /users/{username}/phonenumbers/{label}/validate
func (api UsersAPI) VerifyPhoneNumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	values := struct {
		Smscode       string `json:"smscode"`
		ValidationKey string `json:"validationkey"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		log.Debug("Error decoding the ProcessPhonenumberConfirmation request:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = userobj.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err = api.PhonenumberValidationService.ConfirmValidation(r, values.ValidationKey, values.Smscode)
	if err != nil {
		log.Debug(err)
		if err == validation.ErrInvalidCode || err == validation.ErrInvalidOrExpiredKey {
			writeErrorResponse(w, 422, err.Error())
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdatePhonenumber is the handler for PUT /users/{username}/phonenumbers/{label}
// Update the label and/or value of an existing phonenumber.
func (api UsersAPI) UpdatePhonenumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]

	body := user.Phonenumber{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(body.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !body.IsValid() {
		http.Error(w, "Invalid phone number", http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	oldnumber, err := u.GetPhonenumberByLabel(oldlabel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if oldlabel != body.Label {
		// Check if there already is another phone number with the new label
		_, err := u.GetPhonenumberByLabel(body.Label)
		if err != nil {
			writeErrorResponse(w, http.StatusConflict, "duplicate_label")
			return
		}
	}

	if oldnumber.Phonenumber != body.Phonenumber {
		last, err := isLastVerifiedPhoneNumber(u, oldnumber.Phonenumber, oldlabel, r)
		if err != nil {
			log.Error("ERROR while verifying last verified number - ", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if last {
			writeErrorResponse(w, http.StatusConflict, "cannot_delete_last_verified_phone_number")
			return
		}
	}

	if err = userMgr.SavePhone(username, body); err != nil {
		log.Error("ERROR while saving phonenumber - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != body.Label {
		if err := userMgr.RemovePhone(username, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	valMgr := validationdb.NewManager(r)
	if oldnumber.Phonenumber != body.Phonenumber && isUniquePhonenumber(u, oldnumber.Phonenumber, oldlabel) {
		valMgr.RemoveValidatedPhonenumber(username, oldnumber.Phonenumber)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)

}

// DeletePhonenumber is the handler for DELETE /users/{username}/phonenumbers/{label}
// Removes a phonenumber
func (api UsersAPI) DeletePhonenumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)
	valMgr := validationdb.NewManager(r)
	force := r.URL.Query().Get("force") == "true"

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	number, err := user.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	last, err := isLastVerifiedPhoneNumber(user, number.Phonenumber, label, r)
	if err != nil {
		log.Error("ERROR while checking if number can be deleted:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	if last {
		hasTOTP := false
		if !force {
			writeErrorResponse(w, http.StatusConflict, "warning_delete_last_verified_phone_number")
			return
		} else {
			totpMgr := totp.NewManager(r)
			hasTOTP, err = totpMgr.HasTOTP(username)
		}
		if !hasTOTP {
			writeErrorResponse(w, http.StatusConflict, "cannot_delete_last_verified_phone_number")
			return
		}
	}

	if err := userMgr.RemovePhone(username, label); err != nil {
		log.Error("ERROR while saving user:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := valMgr.RemoveValidatedPhonenumber(username, number.Phonenumber); err != nil {
		log.Error("ERROR while saving user:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Create new bank account
// It is handler for POST /users/{username}/banks
func (api UsersAPI) usernamebanksPost(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	bank := user.BankAccount{}

	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		log.Error("Error while decoding the body: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(bank.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//Check if this label is already used
	_, err = user.GetBankAccountByLabel(bank.Label)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err := userMgr.SaveBank(user, bank); err != nil {
		log.Error("ERROR while saving address:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// respond with created bank account
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(bank)
}

// It is handler for GET /users/{username}/banks
func (api UsersAPI) usernamebanksGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.BankAccounts)
}

// It is handler for GET /users/{username}/banks/{label}
func (api UsersAPI) usernamebankslabelGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	bank, err := userobj.GetBankAccountByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bank)
}

// Update an existing bankaccount and label.
// It is handler for PUT /users/{username}/banks/{label}
func (api UsersAPI) usernamebankslabelPut(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	newbank := user.BankAccount{}

	if err := json.NewDecoder(r.Body).Decode(&newbank); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(newbank.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	oldbank, err := user.GetBankAccountByLabel(oldlabel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if oldlabel != newbank.Label {
		_, err := user.GetBankAccountByLabel(oldbank.Label)
		if err == nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	if err = userMgr.SaveBank(user, newbank); err != nil {
		log.Error("ERROR while saving bank - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != newbank.Label {
		if err := userMgr.RemoveBank(user, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newbank)
}

// Delete a BankAccount
// It is handler for DELETE /users/{username}/banks/{label}
func (api UsersAPI) usernamebankslabelDelete(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = user.GetBankAccountByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err := userMgr.RemoveBank(user, label); err != nil {
		log.Error("ERROR while saving user:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RegisterNewAddress is the handler for POST /users/{username}/addresses
// Register a new address
func (api UsersAPI) RegisterNewAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	address := user.Address{}

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		log.Debug("Error while decoding the body: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(address.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//Check if this label is already used
	_, err = u.GetAddressByLabel(address.Label)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err := userMgr.SaveAddress(username, address); err != nil {
		log.Error("ERROR while saving address:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// respond with created phone number.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(address)
}

// It is handler for GET /users/{username}/addresses
func (api UsersAPI) usernameaddressesGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Addresses)
}

// It is handler for GET /users/{username}/addresses/{label}
func (api UsersAPI) usernameaddresseslabelGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	address, err := userobj.GetAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

// UpdateAddress is the handler for PUT /users/{username}/addresses/{label}
// Update the label and/or value of an existing address.
func (api UsersAPI) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]

	newaddress := user.Address{}
	if err := json.NewDecoder(r.Body).Decode(&newaddress); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !isValidLabel(newaddress.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = u.GetAddressByLabel(oldlabel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if oldlabel != newaddress.Label {
		_, err = u.GetAddressByLabel(newaddress.Label)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	if err = userMgr.SaveAddress(username, newaddress); err != nil {
		log.Error("ERROR while saving address - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != newaddress.Label {
		if err := userMgr.RemoveAddress(username, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newaddress)
}

// DeleteAddress is the handler for DELETE /users/{username}/addresses/{label}
// Removes an address
func (api UsersAPI) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = u.GetAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err = userMgr.RemoveAddress(username, label); err != nil {
		log.Error("ERROR while saving address:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Get the contracts where the user is 1 of the parties. Order descending by date.
// It is handler for GET /users/{username}/contracts
func (api UsersAPI) usernamecontractsGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	includedparty := contractdb.Party{Type: "user", Name: username}
	contract.FindContracts(w, r, includedparty)
}

// RegisterNewContract is handler for GET /users/{username}/contracts
func (api UsersAPI) RegisterNewContract(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	includedparty := contractdb.Party{Type: "user", Name: username}
	contract.CreateContract(w, r, includedparty)

}

// Get the list of notifications, these are pending invitations or approvals
// It is handler for GET /users/{username}/notifications
func (api UsersAPI) usernamenotificationsGet(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	type NotificationList struct {
		Approvals        []invitations.JoinOrganizationInvitation `json:"approvals"`
		ContractRequests []contractdb.ContractSigningRequest      `json:"contractRequests"`
		Invitations      []invitations.JoinOrganizationInvitation `json:"invitations"`
	}
	var notifications NotificationList

	invititationMgr := invitations.NewInvitationManager(r)

	userOrgRequests, err := invititationMgr.GetByUser(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	notifications.Invitations = userOrgRequests

	// TODO: Get Approvals and Contract requests
	notifications.Approvals = []invitations.JoinOrganizationInvitation{}
	notifications.ContractRequests = []contractdb.ContractSigningRequest{}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&notifications)

}

// usernameorganizationsGet is the handler for GET /users/{username}/organizations
// Get the list organizations a user is owner of member of
func (api UsersAPI) usernameorganizationsGet(w http.ResponseWriter, r *http.Request) {

}

// GetAllAuthorizations is the handler for GET /users/{username}/authorizations
// Get the list of authorizations.
func (api UsersAPI) GetAllAuthorizations(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	userMgr := user.NewManager(r)

	authorizations, err := userMgr.GetAuthorizationsByUser(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(authorizations)

}

// GetAuthorization is the handler for GET /users/{username}/authorizations/{grantedTo}
// Get the authorization for a specific organization.
func (api UsersAPI) GetAuthorization(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	grantedTo := mux.Vars(r)["grantedTo"]

	userMgr := user.NewManager(r)

	authorization, err := userMgr.GetAuthorization(username, grantedTo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if authorization == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(authorization)
}

// UpdateAuthorization is the handler for PUT /users/{username}/authorizations/{grantedTo}
// Modify which information an organization is able to see.
func (api UsersAPI) UpdateAuthorization(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	grantedTo := mux.Vars(r)["grantedTo"]

	authorization := &user.Authorization{}

	if err := json.NewDecoder(r.Body).Decode(authorization); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	authorization.Username = username
	authorization.GrantedTo = grantedTo

	userMgr := user.NewManager(r)

	err := userMgr.UpdateAuthorization(authorization)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(authorization)
}

// DeleteAuthorization is the handler for DELETE /users/{username}/authorizations/{grantedTo}
// Remove the authorization for an organization, the granted organization will no longer
// have access the user's information.
func (api UsersAPI) DeleteAuthorization(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	grantedTo := mux.Vars(r)["grantedTo"]

	userMgr := user.NewManager(r)

	err := userMgr.DeleteAuthorization(username, grantedTo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (api UsersAPI) AddAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	body := struct {
		Label string
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	apikeyMgr := apikey.NewManager(r)
	apikey := apikey.NewAPIKey(username, body.Label)
	apikeyMgr.Save(apikey)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apikey)
}

func (api UsersAPI) GetAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	apikeyMgr := apikey.NewManager(r)
	apikey, err := apikeyMgr.GetByUsernameAndLabel(username, label)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apikey)
}

func (api UsersAPI) UpdateAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	apikeyMgr := apikey.NewManager(r)
	body := struct {
		Label string
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	apikey, err := apikeyMgr.GetByUsernameAndLabel(username, label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	apikey.Label = body.Label
	apikeyMgr.Save(apikey)
	w.WriteHeader(http.StatusNoContent)

}
func (api UsersAPI) DeleteAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	apikeyMgr := apikey.NewManager(r)
	apikeyMgr.Delete(username, label)

}
func (api UsersAPI) ListAPIKeys(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	apikeyMgr := apikey.NewManager(r)
	apikeys, err := apikeyMgr.GetByUser(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if apikeys == nil {
		apikeys = []apikey.APIKey{}
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(apikeys)
}

// UpdatePassword handler
func (api UsersAPI) UpdateName(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	values := struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userMgr := user.NewManager(r)
	exists, err := userMgr.Exists(username)
	if !exists || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	err = userMgr.UpdateName(username, values.Firstname, values.Lastname)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeErrorResponse(responseWrite http.ResponseWriter, httpStatusCode int, message string) {
	log.Debug(httpStatusCode, message)
	errorResponse := struct {
		Error string `json:"error"`
	}{
		message,
	}
	responseWrite.WriteHeader(httpStatusCode)
	json.NewEncoder(responseWrite).Encode(&errorResponse)
}
