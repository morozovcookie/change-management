CURRENT_DIR = $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

MIGRATIONS_DIR = $(CURRENT_DIR)/migrations

MIGRATE = $(shell which migrate)
# Verify that migrate is installed.
.PHONY: migrate-check
migrate-check:
	$(call error-if-empty,$(MIGRATE),migrate)

# Creates a new migration.
.PHONY: migration
migration: migrate-check
	@echo "+ $@"
	@$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

define error-if-empty
@if [[ -z "$(1)" ]]; then echo "$(2) not installed"; false; fi
endef
