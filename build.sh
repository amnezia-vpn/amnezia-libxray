#!/bin/bash

set -e

prepare_go() {
    go mod tidy
}

prepare_gomobile() {
    mobile_version=$(go list -m -f '{{.Version}}' golang.org/x/mobile)
    go install golang.org/x/mobile/cmd/gomobile@"$mobile_version"
    go install golang.org/x/mobile/cmd/gobind@"$mobile_version"
    export PATH=~/go/bin:$PATH
    gopath=$(go env GOPATH)
    mkdir -p "${gopath%%:*}/pkg/gomobile"
}

build_apple() {
    rm -fr *.xcframework
    prepare_gomobile
    gomobile bind -target ios,iossimulator,macos -iosversion 15.0
}

build_android() {
    rm -fr *.jar
    rm -fr *.aar
    prepare_gomobile
    rm -fr assets
    mkdir -p assets/geo
    mv dat/* assets/geo
    gomobile bind -target android -androidapi 24 -javapkg=org.amnezia.vpn.protocol.xray -o libxray.aar -ldflags="-w -s -buildid= -checklinkname=0" -trimpath
}

download_geo() {
    go run main/main.go
}

echo "will build libxray for $1"
prepare_go
download_geo
if [ "$1" != "apple" ]; then
build_android
else
build_apple
fi
echo "build libxray done"
