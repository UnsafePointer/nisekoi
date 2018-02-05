#!/usr/bin/env bash

docker build -q -t nisekoi . &> /dev/null
docker run --rm nisekoi $1
docker rmi -q -f nisekoi &> /dev/null
