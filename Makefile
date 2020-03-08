mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(shell dirname $(mkfile_path))

PROJECT_PATH := $(shell cd ${mkfile_dir} && pwd)
PROJECT_DIR := $(shell basename ${PROJECT_PATH})
out_dir := ${PROJECT_PATH}/target

VERSION_TOOLS_IMAGE=helstern/version-tools
VERSION_TOOLS_VERSION=v0.1.0

GO      = go

MODULE   = $(shell env GO111MODULE=on ${GO} list -m)
ALLPKGS = $(shell ${GO} list ./... )
TESTPKGS = $(shell ${GO} list ${MODULE}/test/... )

BIN      = $(CURDIR)/bin

QUIET_FLAG = $(or ${VERBOSE}, 0)
QUIET = $(if $(filter 1,${VERBOSE}),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

TEST_TIMEOUT = 15
TEST_TARGET_VARIANTS := test-default test-bench test-short test-verbose test-race

test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector

${TEST_TARGET_VARIANTS}: NAME=$(MAKECMDGOALS:test-%=%)
${TEST_TARGET_VARIANTS}: test

check test tests:
	$(info $M running ${NAME:%=% }tests…) ## Run tests
	${QUIET} ${GO} test -timeout ${TEST_TIMEOUT}s ${ARGS} ${TESTPKGS}

E2E_SOURCES = $(shell find ${mkfile_dir}/test/e2e -name '*.bats')
e2e-test: $(E2E_SOURCES)
	@ ERROR=0; for file in $^; do bats --tap $${file} || ERROR=1; done; exit $${ERROR}

build: ARGS=-v
build: GOOS ?= linux
build: GOARCH ?= amd64
build: VERSION ?= latest
build: ${mkfile_dir}/internal ${mkfile_dir}/cmd
	${QUIET} cd ${mkfile_dir}/cmd/kommol && GOOS=${GOOS} GOARCH=${GOARCH}  ${GO} build -o ${out_dir}/kommol/kommol ${ARGS}
	${QUIET} cd ${out_dir} && tar -cz -f ${out_dir}/kommol-${VERSION}-${GOOS}-${GOARCH}.tar.gz -C ${out_dir}/kommol .

deps:
	${QUIET} ${GO} mod vendor

# RELEASE

VERSION_TOOLS_IMAGE=helstern/version-tools
VERSION_TOOLS_VERSION=v0.4.1

release-major: ARGS=-M
release-minor: ARGS=-m
release-patch: ARGS=-p

RELEASE_TARGETS:= release-major release-minor release-patch
${RELEASE_TARGETS}: release

changelog:
	${QUIET} docker run --user 1000  \
		--volume ~/.gitconfig:/home/versioneer/.gitconfig \
		--volume ${mkfile_dir}:/home/versioneer/${PROJECT_DIR} \
		--workdir /home/versioneer/${PROJECT_DIR} \
		-it ${VERSION_TOOLS_IMAGE}:${VERSION_TOOLS_VERSION} \
		/bin/sh -c "kacl init"

release:
	@ if test -z "${ARGS}"; then echo "missing release type"; exit 1; fi;
	${QUIET} docker run --user 1000  \
		--volume ~/.gitconfig:/home/versioneer/.gitconfig \
		--volume ${mkfile_dir}:/home/versioneer/${mkfile_dir} \
		--workdir /home/versioneer/${mkfile_dir} \
		-it ${VERSION_TOOLS_IMAGE}:${VERSION_TOOLS_VERSION} \
		/bin/sh -c "release-simple.sh ${ARGS}"
	${QUIET} git push origin && git push --tags origin

.PHONY: ${TEST_TARGETS} check test tests deps ${RELEASE_TARGETS} release
