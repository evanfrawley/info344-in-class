#!/usr/bin/env bash

# you should `source` this script before building the
# Docker container image or trying to run it. Use this
# command to source this into your current shell:
#   source devenv.sh

# docker container image name for our customized MySQL image
export TASKS_MYSQL_IMAGE=evanfrawley/mysqltasks

# database name in which our schema will be created
export MYSQL_DATABASE=tasks

# random MySQL root password
export MYSQL_ROOT_PASSWORD=$(openssl rand -base64 18)
