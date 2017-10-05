#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t evanfrawley/testserver .
docker push evanfrawley/testserver
