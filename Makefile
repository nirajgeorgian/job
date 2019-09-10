-include .env

VERSION                 := $(shell git describe --tags)
BUILD                   := $(shell git rev-parse --short HEAD)
PROJECTNAME             := $(shell basename "$(PWD)")
IMAGE_NAME              := oojob/job
STDERR                  := /tmp/.$(PROJECTNAME)-stderr.txt # Redirect error output to a file, so we can show it in development mode.
PID                     := /tmp/.$(PROJECTNAME)-api-server.pid # PID file will store the server process id when it's running on development mode
SERVER_OUT              := "bin/server"
ENTRYPOINT              := "entry-point.sh"
PKG                     := "github.com/nirajgeorgian/job"
SERVER_PKG_BUILD        := "${PKG}"
PROTOC_ZIP              := "protoc-3.9.1-linux-x86_64.zip"
PKG_LIST                := $(shell go list ${PKG}/... | grep -v /vendor/)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

setup-protobuf-linux: ## install protobuff on linux
	@apt-get update && apt-get -y install unzip
	@curl -OL https://github.com/google/protobuf/releases/download/v3.9.1/$(PROTOC_ZIP)
	@unzip -o $(PROTOC_ZIP) -d /usr/local
	@rm -f $(PROTOC_ZIP)

setup-protobuf-mac: ## install protobuff on mac
	@brew install protobuf
	@brew pin protobuf

setup-dep: ## install dep for dependency management
	@curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@dep version

setup-go: ## install go dependency and setup go environment
	@echo "installing grpc..."
	@go get google.golang.org/grpc
	@cd ${GOPATH}/src/google.golang.org/grpc git checkout v1.23.0;
	@echo "installing proto and protoc-gen-go..."
	@go get github.com/golang/protobuf/proto
	@go get github.com/golang/protobuf/protoc-gen-go
	@cd ${GOPATH}/src/github.com/golang/protobuf git checkout v1.3.0;
	@echo "installing protoc-gen-gorm..."
	@go get github.com/infobloxopen/protoc-gen-gorm
	@cd ${GOPATH}/src/github.com/infobloxopen/protoc-gen-gorm; dep ensure;
	@cd ${GOPATH}/src/github.com/infobloxopen/protoc-gen-gorm git checkout v0.17.0;
	@echo "installing atlas-app-toolkit..."
	@go get github.com/infobloxopen/atlas-app-toolkit/gorm
	@cd ${GOPATH}/src/github.com/infobloxopen/atlas-app-toolkit/gorm git checkout v0.17.0;

protos: ## generate the server and client proto defination files
	@for service in job ; do \
		echo "generating $$service model" ; \
		protoc \
			-I src/proto \
			-I${GOPATH}/src \
			--go_out=${GOPATH}/src \
			--gorm_out=${GOPATH}/src \
			src/proto/model.proto ; \
		echo "generating $$service api" ; \
		protoc \
			-I src/proto \
			-I${GOPATH}/src \
			--go_out=plugins=grpc:${GOPATH}/src \
			src/proto/api.proto ; \
  done

dep: protos ## setup all dependencies
	@dep ensure
	@rm -r vendor/github.com/golang
	@rm -r vendor/github.com/infobloxopen
	@rm -r vendor/google.golang.org/grpc

build: ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

clean: ## Remove previous builds
	@echo "cleaning protos"
	@find . -name "*.pb.go" -type f -not -path "./vendor/*" -delete
	@find . -name "*.pb.gw.go" -type f -not -path "./vendor/*" -delete
	@find . -name "*.pb.gorm.go" -type f -not -path "./vendor/*" -delete

help:
	@IFS=$$'\n' ; \
    help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
    for help_line in $${help_lines[@]}; do \
        IFS=$$'#' ; \
        help_split=($$help_line) ; \
        help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
        printf "%-30s %s\n" $$help_command $$help_info ; \
    done
