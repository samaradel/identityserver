#!/usr/bin/env bash
set -e

VERSION="$(git describe)"

echo "Building version $VERSION"

docker build -t itsyouonlinebuilder .
docker run --rm -v "$PWD":/go/src/github.com/itsyouonline/identityserver --entrypoint sh itsyouonlinebuilder -c "cd siteservice/website && go generate && cd ../.. && go build -ldflags '-s' -v -o dist/identityserver"
docker build -t itsyouonline/identityserver:"$VERSION" -f DockerfileMinimal .

docker push itsyouonline/identityserver:"$VERSION"
