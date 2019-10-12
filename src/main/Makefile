mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(shell dirname $(mkfile_path))

PROJECT_DIR := $(shell cd ${mkfile_dir}/../../ && pwd)
out_dir := ${PROJECT_DIR}/target

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

build: ARGS=-v
build: GOOS ?= linux
build: GOARCH ?= amd64
build: ${mkfile_dir}/internal ${mkfile_dir}/cmd
	${QUIET} cd ${mkfile_dir}/cmd/kommol && GOOS=${GOOS} GOARCH=${GOARCH}  ${GO} build -o ${out_dir}/kommol/kommol ${ARGS}
	${QUIET} cd ${out_dir} && tar -cz -f ${out_dir}/kommol-${GOOS}-${GOARCH}.tar.gz -C ${out_dir}/kommol .

deps:
	${QUIET} ${GO} mod vendor

.PHONY: ${TEST_TARGETS} check test tests deps
