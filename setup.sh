#!/bin/bash

# # Run make
# make

# # Build Docker image
# docker build -t Orangius/hecatearsernal_env -f docker-image/Dockerfile .

cd ./src
go build -o hecate-cli ./cmd/hecate-cli
mv hecate-cli ../hecate-cli