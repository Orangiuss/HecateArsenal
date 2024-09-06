#!/bin/bash

# Run make
make

# Build Docker image
docker build -t Orangius/hecatearsernal_env -f docker-image/Dockerfile .