override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_USER_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_USER_INTERNAL_PROTO := "app/admin-service/internal/conf/config.conf.proto"
SAAS_USER_PROTO_FILES := ""
ifneq ($(SAAS_USER_INTERNAL_PROTO), "")
	SAAS_USER_PROTO_FILES=$(SAAS_USER_API_PROTO) $(SAAS_USER_INTERNAL_PROTO)
else
	SAAS_USER_PROTO_FILES=$(SAAS_USER_API_PROTO)
endif
.PHONY: protoc-admin-protobuf
# protoc :-->: generate admin service protobuf
protoc-admin-protobuf:
	@echo "# generate admin service protobuf"
	$(call protoc_protobuf,$(SAAS_USER_PROTO_FILES))
