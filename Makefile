.PHONY: all
all: help

default: help

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Initialize Project
.PHONY: init
init: ## Just copy `.env.example` to `.env` with one click, executed once.
	@scripts/copy_env.sh

##@ Build

.PHONY: gen
gen: ## gen client code of {svc}. example: make gen svc=user
	@scripts/gen.sh ${svc}

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=user
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=user
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/Vigor-Team/youthcamp-2025-mall-be/app/${svc} --pass "-use github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto

.PHONY: gen-gateway
gen-gateway:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/product_api.proto --service gateway --module github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway -I ../../idl

##@ Build

.PHONY: watch-gateway
watch-gateway:
	@cd app/gateway && air

.PHONY: tidy
tidy: ## run `go mod tidy` for all go module
	@scripts/tidy.sh

.PHONY: lint
lint: ## run `gofmt` for all go module
	@gofmt -l -w app
	@gofumpt -l -w  app

.PHONY: vet
vet: ## run `go vet` for all go module
	@scripts/vet.sh

.PHONY: lint-fix
lint-fix: ## run `golangci-lint` for all go module
	@scripts/fix.sh

.PHONY: run
run: ## run {svc} server. example: make run svc=user
	@scripts/run.sh ${svc}

##@ Development Env

.PHONY: env-start
env-start:  ## launch all middleware software as the docker
	@docker-compose up -d

.PHONY: env-stop
env-stop: ## stop all docker
	@docker-compose down

.PHONY: clean
clean: ## clern up all the tmp files
	@rm -r app/**/log/ app/**/tmp/

.PHONY: build-all
build-all:
	docker build --no-cache -f ./deploy/k8s/Dockerfile.gateway -t gateway:${v} .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t user:${v} --build-arg SVC=user .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t checkout:${v} --build-arg SVC=checkout .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t email:${v} --build-arg SVC=email .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t cart:${v} --build-arg SVC=cart .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t payment:${v} --build-arg SVC=payment .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t product:${v} --build-arg SVC=product .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t order:${v} --build-arg SVC=order .
	docker build --no-cache -f ./deploy/k8s/Dockerfile.svc -t llm:${v} --build-arg SVC=llm .

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module github.com/Vigor-Team/youthcamp-2025-mall-be/app/user --pass "-use github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto


