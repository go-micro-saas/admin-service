override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_USER_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_USER_INTERNAL_PROTO := "app/account-service/internal/conf/config.conf.proto"
SAAS_USER_PROTO_FILES := ""
ifneq ($(SAAS_USER_INTERNAL_PROTO), "")
	SAAS_USER_PROTO_FILES=$(SAAS_USER_API_PROTO) $(SAAS_USER_INTERNAL_PROTO)
else
	SAAS_USER_PROTO_FILES=$(SAAS_USER_API_PROTO)
endif
.PHONY: protoc-account-protobuf
# protoc :-->: generate account service protobuf
protoc-account-protobuf:
	@echo "# generate testdata service protobuf"
	$(call protoc_protobuf,$(SAAS_USER_PROTO_FILES))
