#!/usr/bin/env bash

go build -o $(go env GOPATH)/bin/quickpr -ldflags="-s -w" -trimpath .
