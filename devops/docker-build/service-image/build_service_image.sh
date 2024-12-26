# /bin/sh

CURRENT_FILE_PATH=$(realpath "$0")
CURRENT_FILE_DIR=$(dirname "${CURRENT_FILE_PATH}")
echo "==> The currently executed script file: ${CURRENT_FILE_PATH}"
echo "==> The currently executed script path: ${CURRENT_FILE_DIR}"

# service image
export BUILD_FROM_IMAGE=go-micro-saas/golang-base-image:latest
export RUN_SERVICE_IMAGE=go-micro-saas/golang-release-image:latest
echo "==> build service image BUILD_FROM_IMAGE : ${BUILD_FROM_IMAGE}"
echo "==> build service image RUN_SERVICE_IMAGE : ${RUN_SERVICE_IMAGE}"
docker build \
    --build-arg BUILD_FROM_IMAGE=${BUILD_FROM_IMAGE} \
    --build-arg RUN_SERVICE_IMAGE=${RUN_SERVICE_IMAGE} \
    --build-arg APP_DIR=app \
    --build-arg SERVICE_NAME=account-service \
    --build-arg VERSION=latest \
    -t account-service:latest \
    -f ${CURRENT_FILE_DIR}/Dockerfile_service_image .