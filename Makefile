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
gen: ## gen client code of {svc}. example: make gen svc=order
	@scripts/gen.sh ${svc}

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=order
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=order
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/Vigor-Team/youthcamp-2025-mall-be/app/${svc} --pass "-use github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto

.PHONY: gen-gateway
gen-gateway:
	@cd app/gateway && cwgo server --type HTTP --idl ../../idl/gateway/order_api.proto --service gateway --module github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway -I ../../idl

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
run: ## run {svc} server. example: make run svc=order
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

##@ Open Browser

.PHONY: open.gomall
open-gomall: ## open `gomall` website in the default browser
	@open "http://localhost:8080/"

.PHONY: open.consul
open-consul: ## open `consul ui` in the default browser
	@open "http://localhost:8500/ui/"

.PHONY: open.jaeger
open-jaeger: ## open `jaeger ui` in the default browser
	@open "http://localhost:16686/search"

.PHONY: open.prometheus
open-prometheus: ## open `prometheus ui` in the default browser
	@open "http://localhost:9090"


.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module github.com/Vigor-Team/youthcamp-2025-mall-be/app/order --pass "-use github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto
