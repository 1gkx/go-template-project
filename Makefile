.SILENT: help
.DEFAULT_GOAL = help

gen-pb: #? Generate protocol buffers
	docker run --rm \
		-u $(UID):$(GID) \
		-v $(shell pwd):/gen \
		-w /gen \
		images.digitalms.ru/aux/protoc:4.24.4 \
			--proto_path=/gen/pkg/api \
			--go_out=/gen/pkg/api \
			--go-grpc_out=/gen/pkg/api \
			--go_opt=paths=source_relative \
			--go-grpc_opt=paths=source_relative api.proto

lint-pb: #? Protocol buffers file linter
	docker run --rm \
		-u $(UID):$(GID) \
		-v $(shell pwd):/gen \
		-w /gen \
		images.digitalms.ru/aux/protoc:4.24.4 \
			--protolint_out=. \
			pkg/api/api.proto

lint-pb-fix: #? Protocol buffers file linter fix
	docker run --rm \
		-u $(UID):$(GID) \
		-v $(shell pwd):/gen \
		-w /gen \
		images.digitalms.ru/aux/protoc:4.24.4 \
			--protolint_out=. \
			--protolint_opt=fix \
			pkg/api/api.proto

help: #? Show this message
	@printf "\033[34;01mApplication management:\033[0m\n"
	@awk '"/^@?[a-zA-Z-_0-9]+:/" { \
		nb = sub( /^#\? /, "", helpMsg ); \
		if(nb == 0) { \
			helpMsg = $$0; \
			nb = sub( /^[^:]*:.* #\? /, "", helpMsg ); \
		} \
		if (nb) \
			printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg; \
		} \
		{ helpMsg = $$0 }' \
	$(MAKEFILE_LIST) | column -ts:
