#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t evanfrawley/zipsvr .
docker push evanfrawley/zipsvr
go clean
