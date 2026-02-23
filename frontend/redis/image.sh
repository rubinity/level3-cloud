#!/bin/bash
npm run build
# GOOS=linux GOARCH=amd64 go build -o ./api .
# make
# untag image
# docker image rm api_image
#build image
docker buildx build --platform linux/amd64 -t redisui-image .
# # tag image
# docker tag api_image:latest rubinity/mariia-api:latest
# # push image
# docker push rubinity/mariia-api:latest

# for stackit registry
# tag image
docker tag redisui-image registry.onstackit.cloud/mariia-api/redisui-image:latest
# push image
docker push registry.onstackit.cloud/mariia-api/redisui-image:latest
