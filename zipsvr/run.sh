#!/usr/bin/env bash
docker rm -f zipsvr

docker run -d \
-p 443:443 \
--name zipsvr \
-v $(pwd)/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
evanfrawley/zipsvr
