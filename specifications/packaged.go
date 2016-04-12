// Code generated by go-bindata.
// sources:
// specifications/api/companies.raml
// specifications/api/contracts.raml
// specifications/api/organizations.raml
// specifications/api/securitySchemes/oauth_2_0.raml
// specifications/api/userorganizations.raml
// specifications/api/users.raml
// DO NOT EDIT!

package specifications

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _companiesRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x4b\x73\xdb\x36\x10\xbe\xf3\x57\x6c\xdc\xe9\xcd\xa6\x48\xd9\xc9\x78\x78\xaa\xe3\xa6\x89\x5a\x3b\xe9\xd8\xf5\xc1\x93\xf1\x44\x10\xb1\x92\xd0\x90\x00\x03\x80\xb2\x98\x34\xff\xbd\x4b\x82\x12\x1f\xa6\x62\xd9\x3d\x94\x33\x3a\x08\xfb\xc0\xb7\xdf\x3e\xb0\x3f\xfd\x7c\x75\x76\x79\x01\xa1\x1f\x78\x56\xd8\x04\x23\x98\x58\x73\xab\xf2\x0f\x32\x11\x12\xbd\x15\x6a\x23\x94\x8c\x20\xf0\x43\x6f\xc6\x0c\xde\x68\x11\xc1\xc8\x33\x18\xe7\x5a\xd8\xe2\x3a\x5e\x62\x8a\x26\xf2\x00\x8e\x40\xb1\xdc\x2e\x3f\x8d\x3f\x05\xe5\x5f\xf7\xbd\x10\x32\x4e\x72\x8e\xd0\x33\x18\x6d\x75\x7d\xcd\xd2\xc4\xb3\x45\xe6\xbc\x9c\xab\x34\x63\xb2\x70\x1e\x32\xad\x32\xd4\x56\x38\x91\xfb\x16\x89\x9a\xb1\x44\xf0\xe6\xa4\xfc\x4a\xfb\x08\x8c\xd5\x42\x2e\x3a\x82\x54\xc8\x0b\x94\x0b\xbb\x8c\xe0\xb8\x2b\x60\xeb\x8d\x20\x7c\x19\x6c\x45\x59\x3e\x4b\x44\xfc\x07\x16\x66\xf7\x05\x1f\xef\xfa\x9e\x26\x16\x53\x13\xc1\xb8\xf1\x83\xeb\x4c\x68\x1c\xf2\xc1\x99\xc5\xce\xb1\xc6\x2f\x39\xe9\xf2\x08\xe6\x2c\x31\x8d\x4c\xe9\x05\x93\xe2\x2b\xb3\x94\x80\xe7\xa0\x09\x83\x60\xaf\x7b\x84\x9c\xab\xff\x16\xec\x8f\xbc\x5b\xb6\x96\x7a\xcf\x64\x0d\xf9\xc0\x35\x4b\xb3\x04\x07\x0a\x00\x26\x17\xe2\x33\x9e\x33\xc9\x8b\x47\xd3\x77\x04\xe1\xe9\xf5\xdb\x77\xb7\xd7\xe2\xf4\xf7\xf3\xd5\xd9\xea\xaf\x9b\x57\xb7\xe9\x2a\x7c\x77\x95\xa6\xf9\xe9\xab\xcb\x2f\xe3\xdb\xec\x6b\x3f\x77\x14\x61\x78\x7a\x14\x06\x47\xad\x48\xfb\x5c\x1d\x55\x27\xbf\x88\x84\x90\xc4\x25\x12\x3f\x56\x69\x2f\x76\x78\xfd\x26\x1c\x1f\x9f\xb4\x7f\x1e\xa9\xc4\xae\xd2\x57\x02\xef\x1b\x8f\x8e\x9a\xba\x09\x3c\xd7\x66\xc8\x5f\x17\x11\x7c\x6c\xda\x0b\xee\xbc\x91\xb3\xae\x3b\x23\x53\xc6\x3a\x1f\x1c\x4d\xac\x45\x66\xab\x9e\xbd\xc2\x85\x30\x16\x35\x30\x90\x78\xbf\xb9\xb0\xd2\x9b\x29\x5e\x34\xb7\xb2\x2c\x23\xd6\xaa\x42\x1b\xfd\x6d\xc8\x74\x20\x5d\xe7\x5b\xeb\x05\x6e\x2e\x13\x26\x4b\x58\xf1\x9e\xa5\x24\x7f\x8b\xb6\x56\xb9\xa0\x4b\x1f\x82\xf9\xa7\xe5\x93\x74\x61\x1b\x80\x0f\x67\x14\x97\xd2\x75\xa5\x43\x22\x52\x61\x0d\x30\x8d\x0e\x17\x72\xb0\xaa\x2a\x0d\x34\x96\x2a\x06\x72\x83\xda\xf7\x5c\xbd\x98\x8c\x7a\x03\x3b\xc9\x1e\x07\x41\x17\x7f\x37\xd6\x7d\x22\x7e\x10\x75\xd5\x04\xa3\x6f\xae\xf4\x26\xfc\xbb\x53\xdf\x12\x01\xd0\x4e\x54\x33\x06\xe1\x1b\x98\x58\x95\x93\x8d\xd2\x77\x50\xf3\x1f\x69\x64\xfc\xe0\xb0\xf9\xcf\x38\x0d\xa9\x03\xb8\x83\xef\xb0\xe9\xb5\x0e\x73\x25\x5b\xed\x59\x50\x95\x5c\xad\x38\xc0\x40\x2f\xfe\x87\xd1\x3f\x16\x7b\x3f\xdf\xee\x3b\x09\xc2\xb6\x62\x07\xe1\x8d\x64\x75\x0a\x91\xb7\x0c\x4e\x76\x1a\xbc\x57\x16\xe6\x2a\x97\x4e\x3b\xcb\x9f\xc5\xe3\x63\xbc\xdd\x64\xe5\xa8\xa5\x6e\x16\xae\x6e\x6a\x3b\xdf\x09\xca\x93\xe9\x74\x93\xd1\xe9\x14\x84\x01\x49\xb0\x58\x92\xa8\x7b\xe4\xfe\xff\x40\xf0\xf1\x6e\x82\x37\x88\x37\x78\x4b\xb4\xbf\x29\x3d\x13\x9c\xa3\x7c\x51\x99\x8d\xda\xb3\xe9\x29\x3c\x96\x76\x5d\x1a\x5b\x85\xbd\xa3\xc9\x06\xda\x6c\x98\x8b\x7d\x39\xe9\xb3\xd3\x9a\x8e\x2e\xba\x55\x39\xf3\x29\x9f\xd1\x10\x46\x9a\x0d\xba\xf8\x93\xd1\x2a\x81\x34\xef\xfa\x8f\xa5\xfa\x8c\x3b\x1b\xbc\x7e\x85\xdc\x1d\xb1\x92\x56\xb3\xd8\x9a\xc1\x4b\x9e\x5e\x9b\xad\x26\x6f\x5c\x57\xed\xdf\xa1\x7b\xa0\xdf\xed\x12\x61\x6b\x02\xf7\x4b\xa4\x69\x58\x9e\x75\xe7\x80\x81\x10\xd4\xbc\x12\x64\xac\xda\x93\x7c\xf8\xa0\x39\x4d\xfc\xd2\x21\x4a\x5e\x96\xcc\xac\xa8\x56\x0e\x7f\x1f\xb2\xea\x6d\xed\x4d\xf5\x00\xf2\xe1\xba\x9d\x29\x95\x20\x93\x3d\x59\x27\x82\x49\xbd\xf4\x95\xc0\xdc\x63\xca\x9b\x68\x0e\x2b\x44\x38\x67\x79\x42\x73\x4d\x26\x45\xa5\x46\x12\xb1\x6a\xc7\x5c\xce\x7f\x8d\x36\xd7\xb2\x69\xc6\xcd\xb7\x6b\xd7\xa0\x1c\x59\xe2\x61\x18\xb7\x90\x16\x17\xa8\x7f\x84\xfb\xba\xb4\x26\x46\xe7\x06\xed\x61\xf9\xce\xcc\xf3\x84\xe6\x94\x26\x76\x17\x44\xa5\x0f\xbf\xd6\xb0\x89\xf9\x69\x30\xdd\x1f\x16\x6d\x4d\xcf\x06\x75\xc9\xd6\xe5\xfd\xb4\x42\xd3\x78\x7d\x14\xd5\xcb\xa7\xc0\xaa\xa1\x89\x34\x4f\x69\xdb\xa1\x2d\xf8\xdf\x00\x00\x00\xff\xff\x7b\x0f\xa7\x77\x0e\x0c\x00\x00")

func companiesRamlBytes() ([]byte, error) {
	return bindataRead(
		_companiesRaml,
		"companies.raml",
	)
}

func companiesRaml() (*asset, error) {
	bytes, err := companiesRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "companies.raml", size: 3086, mode: os.FileMode(420), modTime: time.Unix(1458554786, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _contractsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x5d\x6f\xeb\x44\x10\x7d\xf7\xaf\x18\x2e\x42\x80\x54\x12\x27\xea\xbd\xba\xf2\x0b\x2a\xa1\x85\x88\x16\x50\x3f\x1e\x50\x55\x55\x9b\xf5\x24\x5e\xea\xec\x9a\xfd\x68\x63\x4a\xff\x3b\xb3\x6b\x7b\x6d\x27\x94\x0a\x90\xc8\x43\x55\x8f\x77\xe6\x9c\x39\x73\x76\xfc\xe9\x67\x97\x27\x17\xe7\x30\x9b\xa4\x89\x15\xb6\xc4\x0c\x96\xd6\xfc\xa2\xdc\x4f\xb2\x14\x12\x93\x47\xd4\x46\x28\x99\x41\x3a\x99\x25\x2b\x66\xf0\x46\x8b\x0c\xa6\x89\x41\xee\xb4\xb0\xf5\x15\x2f\x70\x8b\x26\x4b\x00\xbe\x02\xc5\x9c\x2d\xee\xe7\xf7\xa9\x7f\x6c\x7e\x9f\x08\xc9\x4b\x97\x23\xec\x25\x4c\xe3\xd9\x89\x66\xdb\x32\x49\x6c\x5d\x35\x65\xae\xc4\x46\x32\xeb\x34\x36\x45\x2a\xad\x2a\xd4\x56\x34\x2f\xfd\xcf\xd0\x01\xcc\xbf\xa9\x33\x30\x56\x0b\xb9\x69\xc3\x39\xb3\x44\xde\xff\x6d\x03\x95\x5b\x95\x82\xff\x80\x75\xcf\xc6\x83\xec\xa5\x51\x22\x1a\xae\x45\x65\x43\x9b\xd7\x05\xb6\x89\xf0\x80\x35\x38\x83\xf9\x00\x76\xc0\xeb\xa0\x1c\x45\x17\x4a\x5a\xcd\xb8\x7d\x8d\x79\xc5\x46\x8f\xe3\x02\xb7\x77\x31\xbc\x15\x72\x69\x71\x6b\x32\x98\xf7\x31\xb6\xeb\x62\x69\x0c\x3a\x29\x7e\x73\xd8\xc6\xad\x76\x5d\xeb\x9c\x88\xa0\xb4\x7b\xad\xf2\x96\xde\xb5\x47\xfd\x7b\x4d\x08\xed\x1c\xe5\xc6\x16\x19\x1c\xa7\x7f\xad\xd4\x99\x46\x84\xb5\xc0\x32\x07\xab\x60\x85\x41\x2b\x58\x2b\x0d\x96\x6d\x0c\xac\x6a\xb0\x5e\x4b\x6a\xb9\x06\xae\x91\x59\xaa\x0f\x2c\x92\x68\x8b\x0a\xf9\xc8\x4a\xe1\xa7\x66\xbe\x7e\x5b\x97\xa8\xc1\xec\x15\x52\xa5\x30\x16\xd4\x3a\xa2\x2c\xf3\xcf\x0d\xf1\x10\x26\x46\x86\x88\x93\xb6\x08\xee\x48\xad\xfc\x7f\x23\x60\x80\xc9\x06\xd3\x5f\x2d\x70\x95\x92\x91\x48\x25\x34\x19\x64\x68\xe3\xbe\xd2\x3f\xf1\xf1\x1f\x31\x0c\xc1\xd3\x7d\x15\x8f\xef\x27\x53\x30\x53\x78\xa6\xdd\x94\xc8\x98\x47\xe1\xa1\x35\x4f\xff\xd0\x79\xe6\x08\x88\x78\x1e\xc2\x81\x28\xf3\x50\x81\xea\x64\x88\xa6\x80\xb3\x92\xbb\x92\xe2\x11\x88\x8a\xb1\x87\xe6\xf1\x57\x43\x49\x1a\x2b\xea\x93\x50\x9a\x1a\x2d\x8d\x28\xd1\x93\xb0\x85\x72\x36\x04\x45\x1e\x50\xe3\xf5\xf3\x2c\xb5\xd8\x86\x77\x12\x9f\xfc\x96\x32\x40\xae\x33\x15\xe3\xf4\xdf\x17\xb8\xe3\x58\x51\x89\x02\xa5\x3f\x53\x03\xab\x2a\x64\x9a\xe6\x1e\x52\x68\xf8\x0e\xcd\x97\xa1\x26\x4a\xae\x68\x35\xd1\x1b\x70\x76\xfd\x71\xd8\xc5\x99\x56\xdb\x66\x6e\x84\xe8\x4a\xdb\x36\xc0\xe0\xea\xfb\x93\xf9\xfb\x0f\xe4\xf3\xb2\x54\x4f\x64\x78\xf2\x39\x83\xcb\xe5\xcf\xa7\x17\xdf\xce\x3e\xa4\x47\x44\x69\x67\xc1\x2f\xca\xf7\x1f\x17\x05\xf2\x87\xd3\x06\xc3\x43\x37\x95\x42\x73\x94\x93\xee\xd2\x94\x76\x04\xae\xc5\x6e\x92\xec\xef\x18\xb2\x40\xdc\x83\xe4\xbf\x66\xdd\x36\x4b\xef\xb6\x5f\xb3\x70\x97\x4c\x3b\xcd\xc2\x56\xa9\x94\x69\x97\xcf\xc8\x0c\x0b\x7f\xfd\x3c\x79\xd2\x2b\x8a\xdc\x74\x4b\x50\x64\x3f\xd3\x2f\xa5\x79\x3a\xeb\x7d\xb6\x52\xf9\x60\x7b\x82\x97\x92\x36\x63\x98\xd9\xd4\xcf\x71\xf8\xae\x73\xe5\x62\x7c\xc1\x8f\x87\xf5\x46\xac\x6e\xa4\xef\x43\x69\xf1\x3b\xed\x58\x3a\x32\x7d\xee\x4d\xfa\xd2\xe4\x6c\xd0\xc6\xad\x3f\x50\xa0\xff\xce\xc0\x33\x18\xae\xfc\x97\x83\x74\x79\xd7\xe5\x67\xd4\x6f\xfe\x0e\xee\xe0\x05\xba\xbb\x3b\x42\xfe\x0e\xed\xe1\x26\x3a\x50\xc2\x6b\x91\x0e\x3b\xdc\x57\xe3\x2d\x3d\x5e\x51\xc4\x6b\x72\x3c\x3c\x39\xe2\xf6\xa3\xb2\x64\x2e\x27\x9b\xcf\xce\x74\xe0\x88\x98\xd1\x4f\xf9\x5f\x48\x13\xae\x3a\x17\x15\x93\x76\xac\xd0\x01\x13\x6f\xc0\x43\x99\xfe\x8b\x10\xd1\xd2\xc9\x9f\x01\x00\x00\xff\xff\x1f\x14\xe6\x48\x76\x08\x00\x00")

func contractsRamlBytes() ([]byte, error) {
	return bindataRead(
		_contractsRaml,
		"contracts.raml",
	)
}

func contractsRaml() (*asset, error) {
	bytes, err := contractsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contracts.raml", size: 2166, mode: os.FileMode(420), modTime: time.Unix(1458557466, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _organizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x59\x5d\x73\xd4\xb8\x12\x7d\xf7\xaf\xd0\xe5\xd6\x7d\x4b\x3c\x9e\x30\x81\x9b\x79\xcb\x02\x45\xa5\x36\x09\x14\x01\x52\x0b\x45\x11\xd9\xd6\xcc\x08\x64\xc9\x58\x72\x32\x03\xcb\x7f\xdf\x96\x64\x8f\x6d\xf9\x63\x9c\x8f\xe5\x01\xbf\xa4\xc6\x96\xd4\xad\xd3\xa7\x8f\xba\x95\xff\xfe\xef\xcd\xf1\xd9\x29\x9a\xfa\x81\xa7\xa8\x62\x64\x8e\x4e\x94\xfc\x4b\xe4\xaf\x38\xa3\x9c\x78\xd7\x24\x93\x54\xf0\x39\x0a\xfc\xa9\x17\x62\x49\xde\x65\x74\x8e\x26\x9e\x24\x51\x9e\x51\xb5\xb9\x88\x56\x24\x21\x72\xee\x21\xb4\x8f\x04\xce\xd5\xea\xf3\xc1\xe7\x40\xff\xb4\xcf\x7f\x28\x8f\x58\x1e\x13\xe4\x4c\x98\x6c\xc7\xfa\x19\x4e\x98\xe7\xa9\x4d\x6a\x97\x79\x95\x2d\x31\xa7\xdf\xb1\xd2\x66\xcd\x3a\x69\x26\x52\x92\x29\x6a\xbf\xeb\x67\xc9\x44\x88\x19\x8d\x2b\x3b\x7a\xfa\x1c\x49\x95\x51\xbe\xdc\xbe\x4c\x28\x3f\x25\x7c\xa9\x56\x73\xf4\xb8\x7a\x89\xd7\xe5\xcb\xe9\x61\x50\xbc\x4e\xf3\x90\xd1\xe8\x4f\xb2\x91\xdd\x4b\x7e\xfc\x54\x9f\x7f\xa2\x48\x22\xe7\xe8\xa0\x9c\x1d\xf3\xdb\x4c\x9b\x06\xc1\xf6\x6d\x4c\x64\x94\xd1\xd4\xec\x15\x9d\x52\xa9\x90\x58\x20\x51\x43\x00\x3d\x3f\xbf\xf0\x8b\xe1\xe2\x86\x43\x34\xee\xe2\x60\x8f\xa1\x5c\xc2\x7a\xe8\x4a\xff\xe1\x18\x82\x72\x55\x5a\x4a\x48\x12\xde\xd6\xd4\xae\x5d\x59\x63\x37\x2b\x81\x70\x46\x0a\x0b\x88\x72\xa4\x56\x54\x36\xb6\x5c\x3a\x51\x30\xe7\x41\xa0\xb5\x84\x39\x89\xb5\x23\x32\x0f\x1b\xf6\xa4\x67\xe6\x90\x35\x4e\x52\xe0\xbf\x4b\x31\xb4\xcc\x08\xe1\x54\xe9\x17\x64\x80\x2e\xfb\x68\x7a\x7c\x74\x79\xb9\x7a\x42\x8f\x5f\xcc\xde\x7c\x78\x79\xfe\x14\xaf\x37\x4f\xd6\x1f\xc2\xfc\x32\x3f\xfd\xca\xbf\x5d\x9e\x66\xef\xbb\xf8\xb2\xdf\x30\xe0\x47\x22\xe9\x09\xf7\xbe\x81\x70\xea\xfc\x3e\x70\x7e\x3f\xee\x0b\xa1\xfd\x3c\xeb\x05\x77\x1f\xe1\x18\x12\xc6\x6f\x79\xe3\x95\x8b\xf5\x65\x63\xc9\x9f\x1d\xd9\xd8\x08\xc8\x3b\x49\x62\x20\x03\xe1\x08\x4b\x49\x97\x1c\x06\x22\x5c\xb2\x42\x01\x47\xb8\xc3\x89\xae\x18\x6d\x0d\xa3\x50\x84\x7a\xc4\x09\xbf\xa6\x6a\x50\x39\xec\xa4\x96\x6f\x99\x60\x35\xf7\x7b\x36\x00\xf6\x79\x9e\xcc\xd1\x47\x13\x99\xbd\xc2\xdb\x8a\x87\x51\x46\xb0\x22\xc0\x98\x18\xfe\xf4\x3a\x6c\x9d\xad\x99\xb5\x81\xf6\x9c\x35\x2e\x72\xbe\x87\x82\x27\xe8\x5c\x5c\xa3\xe9\xd1\xd1\x0c\x05\xff\x9f\xcf\x8e\xe6\x8f\x9f\xa2\x97\x67\x6f\x3d\xab\xbf\x24\xfe\x63\x03\xee\x54\xba\x8b\x3e\x79\x93\x06\xb5\xb5\xe1\x54\x48\x65\x1d\x88\xa9\x4c\x19\xde\x9c\x1b\xc8\x9e\x19\x4b\xe7\xe4\xa6\x2e\xb8\x5e\x2b\x50\x76\x18\xc4\x86\x93\x9b\x66\x48\xd0\xd4\xec\x07\xc9\x95\xc8\x59\x8c\x42\x62\xb3\x99\x14\xc4\x45\x0c\x12\xdf\x47\xef\x75\x12\x59\x29\x83\x3c\x87\x68\x2c\x44\x96\x40\xec\x21\xc6\x70\x12\x44\x5f\x11\x5d\x98\x49\xcd\xf3\x01\x61\xc6\xc4\x8d\x84\xd4\xe6\x78\x09\xbf\x39\x48\x48\x21\x15\xda\xa6\x95\x88\x50\xc4\x9b\x12\x59\x9c\xa6\x90\x91\xc6\xce\xe4\x8b\x2c\xe3\x5f\x85\xb2\xb5\xc9\x8c\xc8\x14\x00\xaa\x98\x71\x10\x4c\xab\x49\xf5\xa5\x87\x97\x1f\x30\x81\xd0\xac\xbe\x66\x93\xff\x5c\xc7\x4c\x64\xf4\x3b\x89\x61\xc4\x92\x94\x21\xaa\x0f\xfa\xbb\x98\xfb\x92\xa8\xa6\x62\xf9\xe8\xb8\x98\x6d\x91\x65\x34\xa1\x4a\x1a\x61\x35\x9e\x5a\x7c\x33\xf2\x2d\x27\x52\xe9\xd4\xaa\x40\xeb\xd8\x77\xf0\x70\xfb\x06\x55\x86\x11\x93\x1f\xa5\x7a\xfe\xb4\x13\xb6\xdb\x43\xa8\x4e\xdc\xaa\x5e\x40\x3f\x90\x8c\x84\xae\x00\x80\xce\x8f\xea\x7b\x9d\xdb\x2c\x7b\xb4\xe7\xbc\x36\x24\x7b\x84\x3e\xa1\x9f\xa8\xcc\xc0\x06\x76\x2e\x66\x40\xce\x85\x28\xd3\xce\xc5\xc0\x41\xa1\x8d\xc3\x2e\x24\x06\x38\xe0\xb0\x60\x98\x07\xe5\x84\x59\xef\x84\x73\xa1\xd0\x42\xe4\xdc\x8e\x4e\xf3\xbb\x03\xbb\x0b\xc1\x77\xa9\xd6\xb1\x5e\x10\x9b\x10\x0d\xc1\xd3\x0b\xcd\xef\x11\x07\x33\x62\xe2\x9c\xb7\x95\xe8\x3e\x54\x6c\x1c\xfb\xc7\xe6\xd8\x6c\x9c\x99\x5d\x45\xd4\xdd\x52\xda\xae\x59\x9d\x8e\xed\x38\x39\x7a\xd9\x72\xef\xcc\x7a\x65\x0f\x77\x90\x23\x99\x47\x11\x91\x72\x91\x33\xb6\x69\xcc\x6a\x47\x78\x77\x8c\x7b\xdc\x6c\x45\x78\x4c\x8c\x5b\x51\x1e\x8e\xb3\x16\xb7\xb2\xec\xf8\x59\xd7\x76\x46\x54\xa3\x82\x78\x98\x90\x3b\xce\xbc\x21\x89\xb8\x26\x55\xd0\x17\x99\x48\x1a\x61\xaf\x4d\xec\x0c\x9a\x0e\xdb\xcc\x85\xb3\x2b\x70\x76\x43\x03\x71\x6b\x41\x3d\x0e\xec\x0e\xb8\x9d\x89\x6f\xa1\x18\x30\x75\x85\xc8\x6c\x35\x51\xd7\x9f\x58\x10\x89\x38\x84\x84\xac\x75\x75\xe1\x2c\x7c\x34\x6e\x61\xcc\xa0\xa6\x89\x37\xba\x1e\x29\x91\x2c\xea\xcb\x49\xb3\xe6\xfe\xd7\x93\xd8\xd4\xab\x3a\x9e\xc6\x2f\x48\xe1\x90\x40\xc5\x5d\x54\x50\xba\x57\x69\x95\xc1\xe5\x32\x77\x15\xc6\x31\x99\xdd\x91\xdb\xdd\x7e\x17\xc5\x6a\x3f\x49\xba\xd3\x7b\x4c\x82\xf7\xa6\xf8\x6f\xc3\x3c\x6e\xa3\xec\x0f\xea\x4a\x97\xb2\x3c\x9c\xb6\xf4\xa8\x4b\xe1\xd9\xa0\xba\xf4\x52\xa7\x53\x61\x1c\x43\xaf\xcc\xf2\x99\x31\x37\xc8\x9e\x8e\x50\x8f\x0d\x76\x67\xb8\x77\x9c\xe1\x91\xe0\x2a\xc3\x91\x92\xf7\xac\xa9\xdc\x5a\xb5\x5a\x57\x87\xbf\x19\x87\x5a\x69\xec\x34\x68\x50\xc1\x3e\x2b\x27\x76\xcb\x87\xae\x71\x35\x51\xb7\xeb\xeb\x8e\x1a\x7a\x80\x16\x79\x81\x6f\x53\x2d\x27\xfa\x43\x8a\x4d\x43\xec\x43\x8d\x14\x1b\xa1\x97\x11\xe1\xb1\x6e\x12\xc2\x8d\xe9\x5c\x2b\x6e\x43\xfb\x90\x6d\x5e\xe3\x0c\xdc\x51\x8d\xbb\x84\xed\xfd\xc1\x8b\x75\x4a\x01\xa0\x6e\xad\x09\x05\xb4\xb7\x98\x0f\x0b\x89\xbd\x1d\xd4\x8e\x11\xbb\x56\xb5\x9b\x3d\xe3\x11\x59\xe0\x9c\xe9\xf6\x8f\x6d\xcc\x30\xf8\x42\xaf\xeb\x7b\xd6\x5d\x4f\x46\x54\x9e\x41\x9d\xe1\x3b\xc6\x74\x0b\x64\x1c\x44\x0b\xcc\x24\xa9\x9f\xcf\x0a\x70\xe8\xf6\x9b\x72\x45\x96\x8e\xe8\x38\x7e\x5f\xe8\xd9\x80\xe8\x42\x12\xb5\xa7\x73\x1b\xd8\x0b\x54\xca\x00\xdd\x25\x40\xe9\xa3\xe7\x85\xdb\x80\xfc\x55\x70\x35\xde\xad\x04\xaf\xef\xec\xd4\x19\x5e\x6b\xfb\xd0\x4b\x43\x36\xec\xf4\xea\xf0\x36\x6e\x15\xae\xd1\x44\xdf\x7f\x1c\x1c\x06\x45\xc2\xd0\xed\x7d\xcb\x7d\x53\x66\x7c\x4e\xbc\xb6\x6c\xad\xae\x7a\x76\x24\x07\x2b\xee\x1f\xd3\x82\xe5\x35\x9f\x0d\x36\xf6\x5e\x12\xce\xde\x2f\x62\xe8\x36\x72\xe0\xac\x0c\xdc\x88\xdd\xf7\xd0\xab\xf6\x66\xfb\x69\x83\xf5\xb8\xb2\xd3\xc1\xcb\x4a\x7a\x0b\xb2\xfe\x33\xe0\x19\xe6\x11\x61\x50\x91\xb4\xe1\xf2\xc7\xc9\x7f\xcf\x01\xd0\x55\x41\x58\x71\x8a\x8c\x49\x46\x4a\x15\xc6\x29\x05\x12\x41\x46\x6f\x97\x6e\x36\x3e\xaf\x4f\x50\xf1\xdd\xa4\xbe\x11\x3b\x4d\xb4\x03\x14\x31\xaa\x6f\x8a\xca\xaf\x9c\x90\xd8\xde\x84\x60\x73\xc2\x58\xf9\x48\xa9\xff\xab\xa8\x0a\xae\x5e\x18\x5f\x4e\x71\x48\xd8\x48\x9e\x32\x33\x16\x5e\x61\x65\xf6\x07\x02\x48\x75\x07\xa5\xb9\x5a\x68\x1f\xec\xa1\xdc\xa4\x6f\x8a\x8b\x72\xc7\xb0\x4e\x22\x09\xbb\x26\x16\x1a\x5d\xae\x14\x42\x1d\xff\x5a\x1e\xf7\xdc\xd5\x37\x0b\xea\xee\xeb\xc8\x2d\x68\xdd\x70\x35\xee\x23\x35\x17\x9c\xc1\xb7\x6b\x77\xbb\x2e\x88\xed\x63\xc2\xd0\xd7\x78\xb6\x2e\x86\xed\x53\xfb\xd7\xd2\x61\xd0\xfe\x5a\xfd\x37\xea\xe0\x0e\x05\xf8\x7d\x82\xd1\xbf\x4d\x67\xbf\x7d\x3b\x2b\x1e\xcb\xb3\xce\x51\x1d\x35\xb0\xfb\x8f\x20\x6d\xc1\x94\xbe\x45\x15\x0c\x82\xb6\xa5\xe5\xe4\x87\x71\xa0\xa6\x6d\x8d\xcc\x1a\xc8\xad\x31\x2d\x6f\x8b\xd9\x7d\x70\xea\x67\x1c\xa4\x63\x30\x1d\x07\xea\x30\xaa\xe3\x6b\xd8\xad\x63\x79\x3f\x70\xf6\x4e\xaf\xa9\x4b\x7d\x97\x0d\x76\xac\x55\x4d\xb3\x91\xb2\x82\x94\x2e\xee\xb7\x6f\x46\x87\xc0\xeb\xc9\xbc\x1d\xb9\xb7\x2b\xfb\x7a\xf2\xef\x36\x19\xd8\x6a\x97\xf4\x69\xec\xb6\x98\xbb\xb2\x00\x14\xae\x40\xd3\xc9\x84\x81\x43\xbd\x11\xc2\xe7\xe6\x73\x17\xfd\x3b\x9a\x39\xd3\x67\x56\x27\xe6\xee\x6d\x8f\x69\xdf\xaa\xf5\xca\x1e\xce\xfb\x27\x00\x00\xff\xff\x54\xe6\xe4\x72\x46\x20\x00\x00")

func organizationsRamlBytes() ([]byte, error) {
	return bindataRead(
		_organizationsRaml,
		"organizations.raml",
	)
}

func organizationsRaml() (*asset, error) {
	bytes, err := organizationsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "organizations.raml", size: 8262, mode: os.FileMode(420), modTime: time.Unix(1459339257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _securityschemesOauth_2_0Raml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x92\x41\x8f\xd3\x40\x0c\x85\xef\xf9\x15\x56\x2e\x9c\x48\x96\xd5\x9e\x72\x5b\x84\xc4\x11\x0e\xf4\x84\x10\x9a\x9d\x31\xcd\x88\x64\x3c\xd8\x4e\x4b\x10\x3f\x1e\x4f\xd3\x54\x6d\xd5\x4a\x9c\xb6\xa7\xea\x79\xde\xf3\x67\xc7\x01\xc5\x73\xcc\x1a\x29\x75\xf0\xb7\x02\xf8\xf4\x3c\x69\xff\x08\x51\xc0\x41\x66\x52\xf2\x34\x80\xf6\x4e\x61\x40\x15\xc0\xdf\x8a\x9c\xdc\x00\x2e\x67\x01\xc6\x5f\x13\x8a\x82\x33\x0f\x71\xfc\xe3\x4a\x0e\x28\x99\x33\xee\x9c\xa2\xe5\x05\x54\x17\x07\x81\x98\x2c\x70\x12\xe4\x37\x96\xec\x3d\x4d\x49\x61\x1f\xcd\x36\x29\x6c\x51\x35\xa6\xad\xb5\xc1\xc8\x90\x9d\xc8\x9e\x38\x34\x95\xce\x19\xbb\x85\x08\x1e\x9b\x87\x6a\x81\x7d\xc1\xf0\x7e\xee\x2c\xba\x47\x17\x90\xa5\xfc\x05\x78\x3e\x47\x58\xa4\xd2\xfc\x6a\xba\xe5\xb7\x11\x0c\x85\x52\x30\x05\xa3\xda\xb9\x21\x86\xb5\x4d\x81\x43\x11\x2b\xff\xc4\xd4\xc0\x07\x82\x44\x5a\xc0\x4d\x31\xce\x1e\xf9\x80\x7d\xca\x32\x05\xea\xc5\xf3\xfd\xe0\xa9\xc1\x76\xc2\x33\x88\x72\x99\x29\x3b\x76\x23\xda\xd2\x1a\xb3\x1c\x2a\x9f\x57\xe5\x48\x7e\x6e\x7e\x6d\xf0\x8b\xad\xd5\xc7\x8d\xda\x13\x46\xc9\x94\x04\x8f\x88\x4f\x0f\x4f\x37\xc9\x36\x69\xfd\xf2\x18\x2a\x59\xbe\xe2\xc1\x72\x71\x10\x1b\x8e\x1d\xf4\xaa\x59\xba\xb6\x8d\x2a\x33\x4d\x0d\xa5\x21\x26\x6c\x77\xef\x5a\x2a\x6f\xdb\x53\x4e\xb5\x2e\xe4\x4b\x99\xe3\xbf\xac\x67\xfb\xbb\x6e\xfd\x91\x5d\x52\xe9\xe0\x2b\x78\x0a\x08\xdf\xac\x2e\x9e\xf2\x3a\xd7\x5b\xa8\xcb\x49\x76\x2e\x8c\x31\xd5\x17\x52\x4c\x3f\xe8\xa4\x10\x6f\x5d\x5a\x6f\x8b\xf6\x09\xf9\x76\x69\xc4\xf1\xe5\x5e\xed\x7e\xa0\xa7\xa4\xec\xbc\x71\xb2\xed\xff\xf4\xc6\xd3\x98\x5d\x9a\xaf\xe0\x56\xf5\xe6\xd3\x8b\x1e\xab\x78\x37\x7e\x91\xef\xa8\x76\xb7\x1a\x7d\xb4\x04\xad\xab\x7f\x01\x00\x00\xff\xff\x62\x8e\x61\x93\x26\x04\x00\x00")

func securityschemesOauth_2_0RamlBytes() ([]byte, error) {
	return bindataRead(
		_securityschemesOauth_2_0Raml,
		"securitySchemes/oauth_2_0.raml",
	)
}

func securityschemesOauth_2_0Raml() (*asset, error) {
	bytes, err := securityschemesOauth_2_0RamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "securitySchemes/oauth_2_0.raml", size: 1062, mode: os.FileMode(420), modTime: time.Unix(1459339155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userorganizationsRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\xae\xc3\x6e\x5b\x9c\x14\x3b\xf9\xd6\x5d\x86\x0e\x1b\x0a\xb4\xd8\x61\x28\x82\x42\x91\xd9\x84\x85\x2c\x09\xa2\xdc\xc1\xcb\xf2\xee\x93\xac\xa4\x96\xb3\x34\x1b\x2f\x4e\x4c\xf2\xff\xc9\x8f\x7e\xfb\xee\xf6\xea\xdb\x57\x58\xcc\xe6\x22\x50\xd0\x58\xc3\x75\xe0\x1f\xb6\xbb\x31\x9a\x0c\x8a\x67\xf4\x4c\xd6\xd4\x30\x9f\x2d\xc4\x4a\x32\x7e\xf7\x54\x43\x25\x18\x55\xe7\x29\xf4\x77\x6a\x83\x2d\x72\x2d\x00\x3e\x80\x95\x5d\xd8\x3c\x5c\x3e\xcc\xd3\xdf\x1c\x6f\xc8\x28\xdd\x35\x08\x47\x0d\xd5\x4b\xed\xcc\xcb\x56\x8b\xd0\xbb\xa4\xb2\xef\xfb\x62\xc9\xdc\xf8\xb5\x34\xf4\x4b\x86\x68\x7f\x6d\x9e\x29\x0c\xbf\x46\x65\xe7\xad\x43\x1f\x28\x9b\x8f\x61\x8b\xbe\x1a\x38\x78\x32\xeb\x49\x41\xc7\xe8\x4f\x26\xbc\x8d\xeb\x4f\xde\xa4\x48\x93\x9d\x2c\x4f\x81\xa6\x6b\x6b\xb8\xb7\x3f\x0d\xfa\xf7\xd0\x62\xbb\x42\xbf\x14\x99\x0e\x36\x9f\xfa\x98\x1b\xa9\xc0\x52\x54\xc9\x9c\xab\x6d\x7a\x18\xd9\xe2\xae\x2a\xc7\x1d\x36\x29\x7b\x47\xa0\xb0\x05\x56\x36\x31\x8a\x8a\x17\xc3\x0a\xb2\x69\xc9\x5c\xc0\x12\x76\x51\x18\x60\x8d\x21\x0f\xdf\x20\x2b\x4f\x2e\xef\xff\x19\x03\x84\x0d\x82\x26\x0e\x13\x34\x0c\x72\x20\x01\xc4\x30\x8c\x1f\xb3\xfb\x05\xc0\x3e\x0e\x42\x1e\xd9\xc5\xc2\x11\xf0\xe5\xbc\xb8\xec\xca\x36\x7d\x49\x4b\x3a\xa7\x49\x0d\xda\xd5\x13\x97\x97\x3a\x77\x2d\xc8\xe6\x07\xc0\xf7\xcb\xa3\x6c\x9e\xa8\x48\xc7\x7c\xb5\x5d\x6b\xbb\x92\x9a\x9a\x5d\x95\x6e\x16\x79\xa6\xc7\x2e\x2b\x3b\xcb\xe1\xe0\x31\x41\x71\xa5\x14\xba\xb0\x57\xe4\x0d\x39\x20\x33\x41\x22\x4e\x2d\x76\x6e\xad\xfc\x69\xbc\xfe\xb1\xee\x4b\xff\xe2\x98\x48\x2e\x4a\xa1\x63\x96\xff\xa2\xf9\x9f\xd6\x0d\x6a\x0c\x78\x12\xc6\xef\x17\xbd\x5b\x7c\x42\x75\x84\xe5\xa0\x92\x08\xc9\x29\xa4\xd9\xb9\xa5\x3e\x96\x63\xbe\xe2\x97\xe2\xae\x8b\xb7\xe0\xc7\x4e\xeb\x3e\x0a\x25\x7f\x6c\x0a\xdb\x99\xf8\x13\x00\x00\xff\xff\xc7\x4d\xbb\x01\x96\x04\x00\x00")

func userorganizationsRamlBytes() ([]byte, error) {
	return bindataRead(
		_userorganizationsRaml,
		"userorganizations.raml",
	)
}

func userorganizationsRaml() (*asset, error) {
	bytes, err := userorganizationsRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "userorganizations.raml", size: 1174, mode: os.FileMode(420), modTime: time.Unix(1458554765, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _usersRaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\x6d\x73\xdb\xb8\x11\xfe\xae\x5f\x81\xba\xed\xf4\xc3\xd9\x7a\x8b\x7d\xbd\xe8\x4b\x2e\xf1\x25\xa9\x53\x3b\xc9\xc4\x4e\x3b\x1d\xd7\xbd\x83\x48\x48\x42\x43\x12\x2c\x00\xd9\x56\x32\xf9\xef\xdd\x05\x49\x11\x20\x41\x91\xb2\x75\xf6\xa5\x53\xce\xb4\x17\x93\xc0\x62\x5f\x9f\x7d\x00\x52\xbf\xff\xe3\x87\xe7\x67\xa7\x64\xd4\x1f\xf6\x34\xd7\x11\x9b\x90\x13\xad\xfe\x21\x96\xef\x92\x88\x27\xac\x77\xcd\xa4\xe2\x22\x99\x90\x61\x7f\xd4\x9b\x52\xc5\x3e\x4a\x3e\x21\x83\x9e\x62\xc1\x52\x72\xbd\x3a\x0f\x16\x2c\x66\x6a\xd2\x23\xe4\x80\x08\xba\xd4\x8b\x9f\xc7\x3f\x0f\xf1\xcf\xec\xfa\x1d\x4f\x82\x68\x19\x32\x52\x99\x30\x58\x8f\xed\x4b\x1a\x47\x3d\xbd\x4a\x33\x29\xa7\x74\xca\xa2\x6c\x3e\xde\x9b\x10\xa5\x25\x4f\xe6\xe6\x46\x4c\x6f\x4f\x59\x32\xd7\x8b\x09\x39\x1a\x66\x77\x78\x52\xdc\x19\xf7\xe0\xce\x79\x20\x60\x8e\x79\x14\x32\x15\x48\x9e\x6a\xa3\xfd\x2b\x21\x09\x4d\x08\xbb\x4d\x23\x9a\x50\xbc\x47\xe8\x54\x2c\x35\x51\x38\x41\xc1\xb3\x30\xfb\x67\x4c\xd3\x14\x96\xdb\x07\x7d\x19\x59\x68\x9d\xaa\xc9\x60\x30\xe7\x7a\xb1\x9c\xf6\x03\x11\x0f\xb8\x56\x2b\xb1\x14\xc6\x39\x03\x1e\xb2\x04\x9c\xb6\x52\x4c\x82\x9f\x06\xd3\x48\x4c\x07\x31\x55\x1a\xfe\x1d\x8a\x20\xb7\x71\x3c\xc8\xd6\xe8\xc7\xa1\xd1\x2b\x95\xf0\x97\xd4\x3c\x33\x37\xbb\x96\x20\x21\xa1\xb1\x6b\x2d\x5e\x73\x49\x13\xcd\xc2\x0b\x51\x8e\xf5\xf8\x25\xbb\x1c\x83\x2f\x16\x8c\x08\x39\xa7\x09\xff\x9c\x99\xab\x17\x54\x13\xae\x0a\x89\x44\x32\x1a\x12\x1a\x04\x4c\x29\xa2\x05\x51\x22\x86\x09\x33\x18\xc6\x8c\x36\x7f\x52\x84\x27\x33\x21\x63\x33\x7b\xbd\x0a\x0d\x43\x09\x33\x9e\x4d\x9a\xd6\x3d\xcb\xfc\x57\x88\x8a\x30\x9a\xaa\xf8\xab\x58\x5b\x24\xcc\x2c\x6a\x8d\xd0\x55\x7d\x25\xfb\xcf\x92\x81\x2b\xc3\xbe\xb5\x94\xcf\x77\x78\xed\x5d\x5e\xed\xb9\x77\x1a\xbc\x94\x2e\xa7\x11\x0f\xfe\xca\x56\x60\x42\xfe\xec\xf2\x6a\xfd\x14\xa2\xcf\xa3\x6f\xd7\xb6\x05\xac\xfd\xcd\x6a\x3f\xa5\xc9\xa7\x6f\x56\xf9\x19\x0d\xd8\x54\x08\x30\x80\xc0\x7f\x22\x46\xcb\x8a\xc9\xb0\xc3\xf3\xc0\x56\x4b\x35\x95\xb7\x95\x9c\x15\x8f\x9c\x72\xa5\xd1\x01\x8e\x18\x63\x6f\x6e\xa2\x71\x97\x6d\x7a\x00\x00\x88\xa8\x06\xf8\x25\x49\xcc\xe2\x29\x40\xfb\x82\xa7\x20\xa4\x8f\xd0\xf9\x02\x02\xf0\x3c\x08\xc4\x32\xd1\x93\x46\xa0\xe2\x10\xa6\x0e\x50\x64\x01\xf5\x93\x61\x19\x62\x1e\x6c\x37\x77\x54\xce\x35\x7a\xc9\xd5\x76\xf3\x0f\x87\x68\xd8\x7b\xac\x8b\x64\x89\xf6\xe6\x86\x51\x0d\x10\x0d\x3e\xfc\xd7\x3f\xbf\x7b\x76\x39\x3c\x78\x7a\xf5\xdd\x1f\x70\xe0\xf3\x0c\xdc\x9a\xad\x0f\x00\xeb\xdd\x7c\x68\xd0\xa1\xc9\x03\x30\x90\x31\xbd\xbd\x88\xa3\x52\x44\x22\xb7\x9f\x6e\xf9\x51\x40\x86\xc8\x67\xf7\x32\xc2\x13\x8a\x6e\x32\x0e\x4b\x19\xa9\x50\x9a\x46\x81\x08\xd9\xf6\x62\xc6\x28\xe6\xa3\x5a\x47\x73\x53\x3f\xbd\x8f\x9d\xe6\x89\xcd\x2f\xd6\xca\xaf\x9b\x88\xaf\x87\xdc\xa6\x5c\x02\x0c\x93\x90\x6a\xb6\xa1\xb3\xfc\x0a\x70\xbf\xbd\x48\xab\x34\x36\xb6\xf8\xed\x25\xe7\xb5\xd4\x0c\xef\xdb\x8b\xb4\x00\xca\x03\xbc\xf7\xc9\xc6\x02\xa3\xef\x52\x97\xe6\x11\xbb\xa5\x71\x1a\x31\x1f\xa1\x9b\x8a\xa9\x2f\x71\x1c\x89\x07\x64\xf4\xfa\x6f\xe7\x7f\x3e\x3a\x3c\xd4\xd7\x7f\x3f\x1b\x5f\x9c\x8d\x9e\x5c\xbf\xfd\xcb\x0f\x6a\x2a\x68\xfc\xe6\xf3\xe9\xf7\xf1\x9b\x27\xdf\x57\x12\x0c\xcb\x60\xf4\xc3\xc1\x68\x78\x30\x1e\xba\x59\xe6\x8a\xbe\x11\xf2\x93\x51\xe2\x47\x60\xaf\x29\x4d\x56\xc8\x62\x9d\x11\x0b\x91\xab\xf9\x63\x6e\x45\x6d\xc4\x5e\x48\x93\x80\x11\xe0\xf1\xd3\xbd\x6c\x24\xb2\x5f\x73\x13\xef\xf5\x8b\x79\xd0\x6d\xdc\xec\xb4\x55\x99\xc2\x22\xdf\x3d\x19\x8f\xc6\x4f\x0e\xf3\xff\x59\x0f\xd9\x1c\x1e\x82\x41\x9e\x87\xd4\x06\x65\x47\xe7\x4a\x92\xe4\xf0\x4c\xce\x53\x8c\xd8\x8c\xb3\x28\xac\x8d\xc8\xf1\x17\xe8\x04\x4f\xf2\x3f\x6a\x63\x00\x60\xc9\xe1\xd1\x8b\xba\xf4\x1c\xf6\xc8\x85\x88\x85\x94\xe2\x06\x76\x12\xf5\x15\x2c\x60\x23\x17\x07\x4f\x0f\x8f\x86\xf5\x70\x34\x28\x7e\x4c\x53\x0e\x73\xc9\x31\xfc\xd5\xa8\x79\x31\x68\x83\xf2\xa3\x9d\xa8\x3e\x1a\x0e\x87\x4e\xf5\xba\x6a\x7f\x9a\x06\x75\x3b\x0c\x3b\x20\x17\xa7\x18\xc3\xa3\xf2\xff\x6a\xe3\x90\x08\x90\xe7\x2f\x8e\x7f\x7a\xf9\xaa\xa3\xae\xbd\xac\xa4\xae\x39\xbb\xd9\x09\xea\x6f\xbb\x0f\xf8\x3f\x5a\xff\x6f\xa0\x75\x71\xd7\x25\xdf\x65\x06\xc0\x80\x37\x82\x27\xef\xac\xc7\x27\xc9\x35\x54\x9c\x21\xdd\x8d\x99\x67\x8b\xab\x69\x81\x69\x59\xbb\x29\x45\xd4\x95\x9d\x60\xc8\x27\xe4\x52\xdc\x24\x4c\xee\xe7\xc4\xdd\x68\x7a\x2c\xa0\x52\x68\xa0\xcf\xf9\x3c\x81\x49\x1f\x32\xde\xbf\x81\xbe\xe6\xe3\x4f\xc2\x7a\x7a\x52\x89\x10\x94\xdf\xcd\x4e\x77\x58\xf8\x02\x6e\x5d\x96\xa7\x3a\xe4\xaa\x37\x40\x63\x8c\x48\x44\x0b\xcf\x41\xcb\xb1\x64\xc0\x7b\x08\x25\x09\xbb\x31\x96\x9b\x21\x53\x11\x5a\x7c\x11\xf6\x72\x50\x7c\xc6\x59\x83\x7f\x2b\x91\xf8\xfc\x80\xfc\x0e\x6d\x1c\x7c\x29\xaa\xfa\x6b\x36\x6c\x5e\xf2\x67\x5b\xcb\xf2\xe8\x89\x7c\xc9\x0f\x75\x50\xf7\x3d\xe3\x7c\x1a\x02\x8f\xdb\x23\x57\xe4\x2b\x29\x8a\x1c\xd2\x3d\x85\xd0\x57\xf3\x79\x3c\x1c\xd6\x61\xcd\xd5\xde\xbe\x36\x5b\xd2\x6c\x15\xd8\x65\x30\x26\xaf\xde\x52\x8b\x2e\x26\x39\xc9\xe6\x31\xad\x8c\x4c\x76\x85\x5c\xa5\x11\x5d\xbd\x35\x74\xe4\x03\x9b\x73\x3c\xa0\x7a\xcb\x6e\x5e\xa2\x06\xd5\xca\xaf\x84\xb3\x18\x9e\x07\xd4\x28\x5d\x60\x4e\x6f\x93\x87\xda\x3c\xd3\x04\x28\x78\x99\xbd\xfb\x24\x3b\x09\xac\x3d\xb4\xfd\xe6\x29\x97\x86\xb8\x62\x64\x47\xf5\x95\xbc\xb6\xb2\xb0\xc5\xda\x26\x9b\xbb\x59\xde\xc5\x03\xdd\x5c\xb1\xbe\x5a\x7c\x82\xd7\xe1\xf0\xa9\x6f\x19\xf7\x48\x01\x57\xc1\x33\x42\x1a\xe1\xe9\xe0\x0a\xeb\x77\x7d\x4e\x32\xf8\x62\x74\xf9\x3a\xb1\x7a\x68\x65\x27\xeb\x24\xda\xc7\x14\x37\x40\x0d\x39\x56\x59\x38\x1b\xab\xca\x83\x1b\x3c\x8f\x1d\x08\x49\xae\x69\xb4\x34\xc7\x92\x78\x76\xdb\x18\x0d\x5f\x24\xda\xa3\xb0\xd9\xff\x2d\x8e\x6f\xf5\xf8\x76\x79\xe8\xf5\x47\x95\xa1\x75\x09\x21\x1e\xfb\x62\xea\x46\xbe\x50\xf6\xca\x29\x11\xd3\x6c\x43\xf0\x7e\x32\x03\xba\x05\xef\x03\x8b\xc5\xb5\x39\x42\xdf\x10\xa2\x46\x77\xa0\x43\x0e\x5b\xcd\x7a\x69\x0b\x06\x61\xb8\xa2\x73\x82\xb7\x9d\x87\x22\xaa\xb4\xab\xac\x39\x1c\x4b\x84\x86\x6d\x4a\x29\x3e\x03\xea\x8c\x6e\x14\x72\x9d\xa4\xdf\x05\x5a\x7b\xcb\xe6\xb5\x59\xb3\x4a\x9f\x3c\x39\x62\x4a\x86\x2a\x25\x02\x8e\x19\x93\x73\x23\x3c\xd7\x77\x66\x6e\x8f\xcf\x5e\x36\xe2\x8d\xa2\x27\xa1\xbd\xc9\x9c\x79\xb3\xa0\x80\x0f\xeb\xcf\x57\xf9\xaa\x77\xf0\x68\xa1\x70\x07\x9f\x76\x41\x9d\xee\x7e\xed\xd2\xb2\x2c\xdf\x5a\x44\x69\x80\x6f\x6e\xee\xce\x96\x70\xb6\xeb\x53\xc7\xa3\xaf\x99\x46\x1a\x73\x52\x7b\x3b\xf4\x38\x94\xaa\xd8\x16\x56\x5c\x00\x9d\x83\xa3\x6f\x6a\x6e\x00\x8e\x2c\x57\xef\xa9\x84\x51\x3a\xa7\xb2\x96\x48\xf1\x89\xb5\x05\xad\xb2\x50\x85\xc2\xdd\x9d\x93\x5a\x3a\x3e\xb4\x27\x3b\x51\x11\xef\x5e\xcf\xe7\x27\xbb\x63\xd8\x5c\xb4\x89\x87\xba\x1d\x66\x13\xfd\x74\x3b\xcb\x36\x9b\x8a\x3b\xd3\xcd\x75\x8f\x77\xb5\xf4\x44\xa8\x52\xad\x9b\xa9\xe5\x2e\x28\xf4\x3d\x29\x4c\x83\x65\xf5\x2e\xda\x95\x23\x56\x19\xe2\xdc\x7e\xd3\xf1\xb0\x29\xed\x3b\xd5\xa8\xd0\x55\x4f\x97\x68\xd9\x0b\x59\xfd\xa1\x99\xa4\xde\xf2\xec\xcd\x5b\xee\xdd\xfe\xe3\x6c\x92\x52\xeb\x95\x57\x2d\xbe\x3b\xe8\x37\xf6\x00\x2f\xe9\xda\x9a\x94\xd6\x29\xa9\x87\x90\xb6\x6e\x56\xd7\x4c\xb4\x5e\x5f\x1b\x08\xb9\x87\x7f\x3a\x72\x9f\x57\x79\x67\xa5\x05\x58\xde\xfe\x95\xbb\x40\xa5\x5c\x76\x8f\x19\xcd\x18\xef\x3f\x82\xec\x82\xef\xf5\x63\xcb\x4d\x18\x9f\xd6\x46\x77\xc7\xf9\x26\xdb\x1a\x6b\xc5\xa9\x93\xba\x9e\xdd\xf8\xee\x66\xa4\x37\x4b\x90\xca\xa9\x6d\xdd\xac\x76\xe3\xda\x8c\x6c\xb7\xb6\xab\xd9\x78\x79\xaa\xfa\x1b\xec\x03\xcd\x27\xe7\x4e\x37\xf0\xf4\x02\x9f\x57\xee\xde\x0d\x2c\x87\xf7\x1b\x37\x10\x77\x65\x30\x1b\xe2\xdd\x1a\xe8\x7b\xed\xe8\xca\xc7\x6d\xd9\x72\x97\x0e\xe0\xc1\xff\xd6\xa0\xac\x3b\x80\x07\x46\x36\xe4\x5e\xad\x01\x38\x42\xdf\x5b\x15\xdc\xd8\x03\xf0\x8d\xc9\xbd\xc1\xdf\x41\x53\xcf\xa9\x3a\xba\x10\x17\xaa\xec\x45\x77\x43\x84\xfd\xd0\xdf\xf4\x32\xa7\x9d\x00\x3f\x7c\x73\xaa\x6a\xf9\xa8\xc0\xb3\xdb\x3d\x95\x6b\x5a\x2d\xf1\x5c\xc8\xdd\x09\xfb\x78\x2c\xa8\xae\x06\xd1\x05\x6a\x0f\x02\x03\xe6\x06\xf9\x4b\x27\x0b\x71\xd1\x2d\x79\x95\x34\x23\x6e\xf7\x03\x9b\xaa\x56\x35\xb0\xb2\x15\xcb\xc0\x0a\x40\xa8\x39\x6a\x89\xd0\x7c\x96\x2f\xbc\x4b\xce\x58\xd1\xe4\x35\xd3\x59\x87\xca\x3f\x16\x74\x96\xdd\xc7\x47\x0a\xf4\x94\x8c\xa4\x2c\x09\xd1\x6d\x7c\xfd\xb2\x53\xa1\x63\xc1\x3d\x52\x40\x47\x53\xbf\xf1\x32\x22\xb6\xe2\x93\x0d\xaf\x70\x9d\xef\x29\xbd\x7a\x65\xf6\xde\x43\x46\xf1\x96\x35\x7f\x1d\x0b\xa2\xfc\xef\x69\xb3\xb7\xcd\x76\x52\x14\x33\x77\x99\x10\xde\x74\x58\x2f\x44\x6e\x16\x40\x55\xd7\x1f\x7d\x63\x67\x1e\x15\x5f\xd5\xe2\x2b\x61\x70\x7b\x9f\xbc\x93\x21\x3c\x42\x41\x79\x92\x4c\x57\xe6\xe3\xb6\xa2\xae\x1a\x4f\xd4\xf2\x4f\xff\x5f\x9a\x0f\x96\xc2\xca\x49\x9b\x29\xab\xea\x37\xb1\x35\x8d\x4f\xf2\x5f\x0f\xa0\x42\xd9\x87\x4f\x61\xa9\xfd\xbe\xd1\x84\xcd\xe8\x32\x82\xec\x4e\xa2\x55\x76\x76\x1b\x68\x7e\x6d\xdb\x88\xf9\x2d\x99\x5e\xca\xa4\xfa\xce\x00\xbf\x94\x35\xaa\x91\x19\xc4\xbc\xfc\x5a\x4f\x69\xb0\xdd\xa7\x2f\x4f\x34\x9b\x57\x78\xbc\xa3\xef\x39\xce\x04\x0f\xce\x14\xd3\xfb\xe8\xd3\xd9\x32\x22\x33\xa8\xa4\x94\xce\xc1\x75\x7d\xc0\x86\x4c\x5d\xf0\xf4\x2f\xc3\x5f\xba\xa9\x13\xd3\xdb\x3b\x29\x73\x46\x6f\x71\x5d\x46\x14\xff\xcc\x5a\xb5\x39\xea\xaa\x4e\xae\x12\x8f\xf1\xb3\x84\x71\xf6\x09\x9c\x9d\xc6\x79\x7a\x3e\x18\xa8\xa1\x58\x21\x8b\xaf\x9c\x33\xc1\x8f\xdc\xc4\xcc\xcf\x50\x0c\x4e\xd4\x1d\x33\xf8\xb2\xfe\x51\xc7\x6e\xfb\x75\xcd\x49\x94\xa8\x94\x05\x08\xf8\xae\x8f\x7e\x0b\xce\x69\xef\xa2\x19\x97\x07\x23\xcc\xf8\x7d\xf7\x4b\x7f\xfb\xbb\xf6\x1b\x1e\x45\xd0\xd9\x48\x24\x12\x28\x07\xb2\xa0\x38\x2b\xff\x61\x8b\xf7\xd7\x2c\xfd\x6e\xc4\x22\x5f\xb9\xf7\xdf\x00\x00\x00\xff\xff\xc0\x37\x68\x3f\x22\x35\x00\x00")

func usersRamlBytes() ([]byte, error) {
	return bindataRead(
		_usersRaml,
		"users.raml",
	)
}

func usersRaml() (*asset, error) {
	bytes, err := usersRamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "users.raml", size: 13602, mode: os.FileMode(420), modTime: time.Unix(1459349556, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"companies.raml": companiesRaml,
	"contracts.raml": contractsRaml,
	"organizations.raml": organizationsRaml,
	"securitySchemes/oauth_2_0.raml": securityschemesOauth_2_0Raml,
	"userorganizations.raml": userorganizationsRaml,
	"users.raml": usersRaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"companies.raml": &bintree{companiesRaml, map[string]*bintree{}},
	"contracts.raml": &bintree{contractsRaml, map[string]*bintree{}},
	"organizations.raml": &bintree{organizationsRaml, map[string]*bintree{}},
	"securitySchemes": &bintree{nil, map[string]*bintree{
		"oauth_2_0.raml": &bintree{securityschemesOauth_2_0Raml, map[string]*bintree{}},
	}},
	"userorganizations.raml": &bintree{userorganizationsRaml, map[string]*bintree{}},
	"users.raml": &bintree{usersRaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

