mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir_path := $(shell dirname $(mkfile_path))

QUIET_FLAG = $(or ${VERBOSE}, 0)
QUIET = $(if $(filter 1,${VERBOSE}),,@)

VERSION_TOOLS_IMAGE=helstern/version-tools
VERSION_TOOLS_VERSION=v0.1.0

release-major: ARGS=-M
release-minor: ARGS=-m
release-patch: ARGS=-p

RELEASE_TARGETS:= release-major release-minor release-patch
${RELEASE_TARGETS}: release

deps:
	${QUIET} docker pull ${VERSION_TOOLS_IMAGE}:${VERSION_TOOLS_VERSION}
	${QUIET} cd src/main && make deps

changelog:
	${QUIET} docker run --user 1000  \
		--volume ~/.gitconfig:/home/versioneer/.gitconfig \
		--volume ${mkfile_dir_path}:/home/versioneer/${mkfile_dir} \
		--workdir /home/versioneer/${mkfile_dir} \
		-it ${VERSION_TOOLS_IMAGE}:${VERSION_TOOLS_VERSION} \
		/bin/sh -c "kacl init"

release:
	${QUIET} docker run --user 1000  \
		--volume ~/.gitconfig:/home/versioneer/.gitconfig \
		--volume ${mkfile_dir_path}:/home/versioneer/${mkfile_dir} \
		--workdir /home/versioneer/${mkfile_dir} \
		-it ${VERSION_TOOLS_IMAGE}:${VERSION_TOOLS_VERSION} \
		/bin/sh -c "release-simple.sh ${ARGS}"
	git push origin && git push --tags origin

.PHONE: ${RELEASE_TARGETS} deps release