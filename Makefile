include Configfile

.PHONY: build
build: build-linux build-mac build-windows

build-windows: export GOOS=windows
build-windows: export GOARCH=amd64
build-windows: export GO111MODULE=on
build-windows: export GOPROXY=$(MOD_PROXY_URL)
build-windows:
	$(GO) build -v --ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Build=$(BUILD_DATE)" \
		-o bin/windows/amd64/coc.exe cmd/coc/coc.go  # windows

build-linux: export GOOS=linux
build-linux: export GOARCH=amd64
build-linux: export CGO_ENABLED=0
build-linux: export GO111MODULE=on
build-linux: export GOPROXY=$(MOD_PROXY_URL)
build-linux:
	$(GO) build -v --ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Build=$(BUILD_DATE)" \
		-o bin/linux/amd64/coc cmd/coc/coc.go  # linux

build-mac: export GOOS=darwin
build-mac: export GOARCH=amd64
build-mac: export CGO_ENABLED=0
build-mac: export GO111MODULE=on
build-mac: export GOPROXY=$(MOD_PROXY_URL)
build-mac:
	$(GO) build -v --ldflags="-w -X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Build=$(BUILD_DATE)" \
		-o bin/darwin/amd64/coc cmd/coc/coc.go  # mac osx

.PHONY: clean
clean::
	echo "--> cleaning..."
	rm -rf vendor
	go clean ./...