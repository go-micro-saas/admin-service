override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_ADMIN_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_ADMIN_INTERNAL_PROTO := "app/admin-service/internal/conf/config.conf.proto"
SAAS_ADMIN_PROTO_FILES := ""
ifneq ($(SAAS_ADMIN_INTERNAL_PROTO), "")
	SAAS_ADMIN_PROTO_FILES=$(SAAS_ADMIN_API_PROTO) $(SAAS_ADMIN_INTERNAL_PROTO)
else
	SAAS_ADMIN_PROTO_FILES=$(SAAS_ADMIN_API_PROTO)
endif
.PHONY: protoc-admin-protobuf
# protoc :-->: generate account service protobuf
protoc-admin-protobuf:
	@echo "# generate testdata service protobuf"
	$(call protoc_protobuf,$(SAAS_ADMIN_PROTO_FILES))
