#https://github.com/webrocket/webrocket/blob/master/Makefile


# .PHONY: build doc fmt lint run test vendor_clean vendor_get vendor_update vet
# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
UNAME=$(shell uname -s)
SOURCE_DIR=$(shell pwd)
SCRIPTS_DIR=$(SOURCE_DIR)/scripts
EXTRA_DEPS=
ORIGNGP = GOPATH

GOPATH = ${PWD}/_buildTMP
export GOPATH

ifeq ($(UNAME),Darwin)
	ECHO=echo
else
	ECHO=echo -e
endif

ASCIIDOC=asciidoc
CAT=cat

# default: build

all: build
	rm -rf _buildTMP
	-@$(ECHO) "\n\033[1;32mCONGRATULATIONS!\033[0;32m\nVerla has been built and tested!\033[0m\n"

check: tools format
	-@$(ECHO) ""

clean: clean-tools
	rm -rf bin

format:
	-@$(ECHO) "\n\033[0;35m%%% Formatting\033[0m"
	@go fmt ./...

build: clean $(BUILD_MAN)
	-@$(ECHO) "\033[0;35m%%% Resolving dependencies\033[0m"
	@mkdir ${GOPATH}
	@mkdir ${GOPATH}/src
	@cp -r verla ${GOPATH}/src/verla
	@cd ${GOPATH}/src/verla && go get -v ./...
	-@$(ECHO) "\n\033[0;35m%%% Building libraries and tools\033[0m"
	-@$(ECHO) "verla-server"
	@go build ./_buildTMP/src/verla/server
	-@$(ECHO) "\n\033[0;35m%%% Running tests\033[0m"
	@cd ./_buildTMP/src/verla && go test ./...
	-@$(ECHO) "\n\033[0;35m%%% \033[0m"
	@mkdir bin && mv server bin/verla-server

clean-tools:
	@go clean ./...
	rm -f server
	rm -rf _buildTMP

man: clean-man
	@$(MAKE) -C docs

install-man:
	-@$(ECHO) "\033[0;36mInstalling documentation\033[0m"
	@$(MAKE) -C docs install

clean-man:
	-@$(MAKE) -C docs clean







# build: vendor_update
# 	go build -v -o ./bin/verla ./src/verla
#
# doc:
# 	godoc -http=:6060 -index
#
# # http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
# fmt:
# 	go fmt ./src/...
#
# # https://github.com/golang/lint
# # go get github.com/golang/lint/golint
# lint:
# 	golint ./src
#
# run: build
# 	./bin/verla
#
# test:
# 	go test ./src/...
#
# vendor_clean:
# 	rm -dRf ./_vendor/src
#
# # We have to set GOPATH to just the _vendor
# # directory to ensure that `go get` doesn't
# # update packages in our primary GOPATH instead.
# # This will happen if you already have the package
# # installed in GOPATH since `go get` will use
# # that existing location as the destination.
# vendor_get: vendor_clean
# 	GOPATH=${PWD}/_vendor go get -d -u -v \
# 	github.com/akrennmair/goconf \
# 	github.com/bradfitz/gomemcache/memcache
#
# vendor_update: vendor_get
# 	rm -rf `find ./_vendor/src -type d -name .git` \
# 	&& rm -rf `find ./_vendor/src -type d -name .hg` \
# 	&& rm -rf `find ./_vendor/src -type d -name .bzr` \
# 	&& rm -rf `find ./_vendor/src -type d -name .svn`
