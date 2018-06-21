APPNAME := plugin-tester
GOPATH := $(CURDIR)
BINPATH := bin
SRCPATH := src
PLUGINPATH := $(BINPATH)/plugins
PKGNAME := plugintester

build:
	rm -rf $(PLUGINPATH)
	mkdir -p $(PLUGINPATH)
	go build -buildmode=plugin -o $(PLUGINPATH)/plugin01.so plugin01/main
	go build -buildmode=plugin -o $(PLUGINPATH)/plugin02.so plugin02/main
	go build -o $(BINPATH)/$(APPNAME) $(PKGNAME)/main

run: build
	@echo
	@bin/$(APPNAME)

clean:
	rm -rf pkg
	rm -rf $(BINPATH)

dep-init:
	cd $(SRCPATH)/$(PKGNAME) && dep init

dep-ensure:
	cd $(SRCPATH)/$(PKGNAME) && dep ensure -v
