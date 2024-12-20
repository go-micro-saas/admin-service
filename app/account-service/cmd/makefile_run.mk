.PHONY: store-configuration
# run :-->: run database-migration
store-configuration:
	go run ./app/account-service/cmd/store-configuration/... -conf=./app/account-service/configs

.PHONY: run-database-migration
# run :-->: run database-migration
run-database-migration:
	go run ./app/account-service/cmd/database-migration/... -conf=./app/account-service/configs

.PHONY: run-account-service
# run service :-->: run account-service
run-account-service:
	go run ./app/account-service/cmd/account-service/... -conf=./app/account-service/configs

.PHONY: testing-account-service
# testing service :-->: testing account-service
testing-account-service:
	@echo "==> testing-account-service"


.PHONY: run-service
# run service :-->: run account-service
run-service:
	#@$(MAKE) run-account-service
	go run ./app/account-service/cmd/account-service/... -conf=./app/account-service/configs

.PHONY: testing-service
# testing service :-->: testing account-service
testing-service:
	$(MAKE) testing-account-service
