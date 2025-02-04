.PHONY: all
all: help

default: help

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Code Generate with cwgo

.PHONY: gen
gen: ## gen client code of {svc}. example: make gen svc=product
	@scripts/gen.sh ${svc}

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=product
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/zheyuanf/ecommerce-tiktok/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=product
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/zheyuanf/ecommerce-tiktok/app/${svc} --pass "-use github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/zheyuanf/ecommerce-tiktok/app/frontend --idl ../../idl/frontend/${svc}.proto

.PHONY: tidy
tidy: ## run `go mod tidy` for all go module
	@scripts/tidy.sh

##@ Build Development Env with docker-compose

.PHONY: env-start
env-start:  ## launch all middleware software as the docker
	@docker-compose up -d

.PHONY: env-stop
env-stop: ## stop all docker
	@docker-compose down

.PHONY: clean
clean: ## clern up all the tmp files
	@rm -r app/**/log/ app/**/tmp/

##@ 打开浏览器

.PHONY: open.consul
open-consul: ## 在浏览器打开consul的ui界面，查看注册的服务
	@open "http://localhost:8500/ui/"

##@ 构建docker镜像

.PHONY: build-frontend
build-frontend:
	docker build -f ./deploy/Dockerfile.frontend -t ecommerce-frontend:${v} .

.PHONY: build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ecommerce-${svc}:${v} --build-arg SVC=${svc} .