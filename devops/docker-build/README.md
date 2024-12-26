# build

统一端口

* http: 8080
* grpc: 50051

```shell
docker build \
    --build-arg BUILD_FROM_IMAGE=golang:1.22.8 \
    --build-arg RUN_SERVICE_IMAGE=debian:stable-20240926-slim \
    --build-arg APP_DIR=app \
    --build-arg SERVICE_NAME=admin-service \
    --build-arg VERSION=latest \
    -t admin-service:latest \
    -f ./devops/docker-build/Dockerfile .
```