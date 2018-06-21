APPNAME := plugin-tester
GOPATH := $(CURDIR)

build:
	go build -o $(APPNAME) plugintester/main

clean:
	rm -rf pkg
	rm -f $(APPNAME)

dep-init:
	cd src/plugintester && dep init

dep-ensure:
	cd src/plugintester && dep ensure -v
