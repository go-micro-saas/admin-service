.PHONY: rm-none-image
# rm none:none images
rm-none-image:
	@echo "rm :-->: rm none images"
	/bin/sh ./devops/docker-build/service-image/rm_none_images.sh

.PHONY: build-base-image
# build :-->: base image
build-base-image:
	@echo "build :-->: build base image"
	/bin/sh ./devops/docker-build/service-image/build_base_image.sh
	/bin/sh ./devops/docker-build/service-image/build_release_image.sh
	$(MAKE) rm-none-image

.PHONY: build-service-image
# build :-->: service image
build-service-image:
	@echo "build :-->: build service image"
	$(MAKE) build-base-image
	/bin/sh ./devops/docker-build/service-image/build_service_image.sh
	$(MAKE) rm-none-image

.PHONY: build
# build :-->: service image
build:
	@echo "build :-->: build service image"
	#docker build -t account-service:v1.0.0 -f ./devops/Dockerfile .
	#docker pull golang:1.23.10
	#docker pull debian:stable-20250520
	docker build \
		--build-arg BUILD_FROM_IMAGE=golang:1.23.10 \
		--build-arg RUN_SERVICE_IMAGE=debian:stable-20250520 \
		--build-arg APP_DIR=app \
		--build-arg SERVICE_NAME=account-service \
		--build-arg VERSION=latest \
		-t account-service:latest \
		-f ./devops/docker-build/Dockerfile .

.PHONY: deploy-general-config
# deploy :-->: store general config
deploy-general-config:
	@echo "deploy :-->: general config"
	go run ./app/account-service/cmd/store-configuration/... \
      -conf=./app/account-service/configs \
      -source_dir ./devops/docker-deploy/general-configs \
      -store_dir go-micro-saas/general-configs/testing

.PHONY: deploy-service-config
# deploy :-->: store service config
deploy-service-config:
	@echo "deploy :-->: service config"
	go run ./app/account-service/cmd/store-configuration/... \
      -conf=./app/account-service/configs \
      -source_dir ./devops/docker-deploy/service-configs \
      -store_dir go-micro-saas/account-service/testing/latest

.PHONY: deploy-database-migration
# deploy :-->: database migration
deploy-database-migration:
	@echo "deploy :-->: database migration"
	$(MAKE) run-database-migration

.PHONY: deploy-on-docker
# deploy :-->: image on docker
deploy-on-docker:
	@echo "deploy-on-docker :-->: deploying on docker"
	docker-compose -f ./devops/docker-deploy/docker-compose.yaml up -d

.PHONY: stop-docker-deploy
# stop :-->: docker container
stop-docker-deploy:
	@echo "deploy-on-docker :-->: deploying on docker"
	docker-compose -f ./devops/docker-deploy/docker-compose.yaml down
