.PHONY: store-configuration
# run :-->: run database-migration
store-configuration:
	go run ./app/admin-service/cmd/store-configuration/... -conf=./app/admin-service/configs

.PHONY: run-database-migration
# run :-->: run database-migration
run-database-migration:
	go run ./app/admin-service/cmd/database-migration/... -conf=./app/admin-service/configs

.PHONY: run-admin-service
# run service :-->: run admin-service
run-admin-service:
	go run ./app/admin-service/cmd/admin-service/... -conf=./app/admin-service/configs

.PHONY: testing-admin-service
# testing service :-->: testing admin-service
testing-admin-service:
	@echo "==> testing-admin-service"


.PHONY: run-service
# run service :-->: run admin-service
run-service:
	#@$(MAKE) run-admin-service
	go run ./app/admin-service/cmd/admin-service/... -conf=./app/admin-service/configs

.PHONY: testing-service
# testing service :-->: testing admin-service
testing-service:
	$(MAKE) testing-admin-service
