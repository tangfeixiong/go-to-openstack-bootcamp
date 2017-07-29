IMPORT_PATH = $(shell echo `pwd` | sed "s|^$(GOPATH)/src/||g")
APP_NAME = $(shell echo $(IMPORT_PATH) | sed 's:.*/::')
APP_VERSION = 0.1
TARGET = ./bin/$(APP_NAME)-$(APP_VERSION)
GO_FILES = $(shell find . -type f -name "*.go")
BUNDLE = public/bundles
ASSETS = $(shell find assets -type f)
PID = .pid
NODE_BIN = $(shell npm bin)
#go server port
PORT ?= 9000
#webpack-dev-server port
DEV_HOT_PORT ?= 8090

build: clean $(BUNDLE) $(TARGET)

clean:
	@rm -rf public/bundles
	@rm -rf $(TARGET)

$(BUNDLE): $(ASSETS)
	@$(NODE_BIN)/webpack --progress --colors

$(TARGET): $(GO_FILES)
	@printf "Buiding go binary ......"
	@go build -race -o $@

kill:
	@kill `cat $(PID)` || true

dev: clean $(BUNDLE) restart
	@DEV_HOT=true $(NODE_BIN)/webpack-dev-server --config webpack.config.js &
	@printf "\n\nWaiting for the file change\n\n"
	@fswatch --event=Updated $(GO_FILES) | xargs -n1 -I{} make restart || make kill

restart: kill $(TARGET)
	@printf "\n\nrestart the app .........\n\n"
	@$(TARGET) -debug --web=:$(PORT) --devWeb=:$(DEV_HOT_PORT) & echo $$! > $(PID)

dist: clean $(TARGET)
	@NODE_ENV=production $(NODE_BIN)/webpack --progress --colors
	@zip -r -v $(APP_NAME)-$(APP_VERSION).zip $(TARGET) webpack-assets.json public templates

.PHONY: kopos-bin
kopos-bin:
	@CGO_ENABLED=0 go build -v -o bin/kopos --installsuffix cgo .

.PHONY: kopos-docker
IMG_NAME ?= docker.io/tangfeixiong/go-to-openstack-bootcamp
IMG_TAG ?= latest
kopos-docker: kopos-bin
	@sed 's%bin/kopos%kopos%' Dockerfile >./bin/Dockerfile
	@docker build -t $(IMG_NAME):$(IMG_TAG) ./bin
	@docker push $(IMG_NAME):$(IMG_TAG)

.PHONY: kopos-run
OS_PASSWORD ?= secret
MOCK_HOST ?= 127.0.0.1
SSH_USER ?= vagrant
kopos-runtest: 
	$(shell inst=$$(docker ps -qf name=kopos); [[ -z $$inst ]] && docker stop $$inst && docker rm $$inst )
	@docker run -d --name kopos -p 10001:10001 -e OS_PASSWORD=$(OS_PASSWORD) -e MOCK_HOST=$(MOCK_HOST) --restart always docker.io/tangfeixiong/go-to-openstack-bootcamp:0.1 -k /Users/fanhongling/Downloads/tmp/vagrant -u vagrant --loglevel 2
