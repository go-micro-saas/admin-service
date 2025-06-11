# build

统一端口

* http: 8080
* grpc: 50051

```shell
docker build \
    --build-arg BUILD_FROM_IMAGE=golang:1.23.10 \
    --build-arg RUN_SERVICE_IMAGE=debian:stable-20250520 \
    --build-arg APP_DIR=app \
    --build-arg SERVICE_NAME=account-service \
    --build-arg VERSION=latest \
    -t account-service:latest \
    -f ./devops/docker-build/Dockerfile .
```