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
