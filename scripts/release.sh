#!/usr/bin/env bash

# Usage:
#   hack/release.sh $VERSION
#
# Example:
#   hack/release.sh v0.1.1

VERSION=$1

GO111MODULE=off go get github.com/ahmetb/govvv
govvv build -o build/account-service service/account/main.go service/account/plugin.go -version $VERSION
git tag $VERSION
git push origin $VERSION
