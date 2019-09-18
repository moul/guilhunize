GOPKG ?=	moul.io/guilhunize
DOCKER_IMAGE ?=	moul/guilhunize
GOBINS ?=	.
NPM_PACKAGES ?=	.
LAMBDA_PORT ?=	8080

all: test install

-include rules.mk

.PHONY: netlify
netlify: lambda-build

.PHONY: lambda-build
lambda-build:
	rm -rf lambda-build
	cd lambda && go get
	cd lambda && GOOS=linux GOARCH=amd64 $(GO) build -o ../lambda-build/guilhunize guilhunize.go

.PHONY: lambda-dev
lambda-dev: lambda-build
	@echo "Open http://localhost:$(LAMBDA_PORT)/guilhunize"
	sam local start-api --host=0.0.0.0 --port=$(LAMBDA_PORT) --static-dir=web
