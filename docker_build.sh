#! /bin/bash

docker build -t blockchain-test:0.1.0 . -f Dockerfile --rm=true

docker-compose up


