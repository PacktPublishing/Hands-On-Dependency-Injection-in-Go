#!/usr/bin/env bash

cd $GOPATH/src/github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/

package-coverage -a -prefix $(go list)/ ./acme/
