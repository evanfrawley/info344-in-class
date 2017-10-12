#!/usr/bin/env bash
set -e
docker build -t evanfrawley/zipclient .
docker run -it --rm -p 5000:5000 --name demo evanfrawley/zipclient