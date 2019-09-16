GOPKG ?=	moul.io/guilhunize
DOCKER_IMAGE ?=	moul/guilhunize
GOBINS ?=	.
NPM_PACKAGES ?=	.

all: test install

-include rules.mk
