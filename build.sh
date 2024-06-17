#!/usr/bin/env bash
cd $(dirname -- $0)

go build -o $(go env GOPATH)/bin/quickpr .
