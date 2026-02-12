#!/bin/bash

# GOOS=linux GOARCH=amd64 go build -o ./api .
make
# untag image
# docker image rm api_image
#build image
docker buildx build --platform linux/amd64 -t api_image .
# tag image
docker tag api_image:latest rubinity/mariia-api:latest
# push image
docker push rubinity/mariia-api:latest
