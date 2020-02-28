#!/usr/bin/env bash

READLINK=$(which greadlink)
if test -z "${READLINK}"; then
    READLINK=$(which readlink)
fi

if test -z "${READLINK}"; then
    echo "missing readlink command."
    exit 1
fi

SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(${READLINK} -f "${SCRIPT_DIR}")"
PROJECT_DIR="$(cd ${SCRIPT_DIR}/../../ && pwd)"

function bootstrap() {
    set -o allexport
    source ${PROJECT_DIR}/.env
    set +o allexport

    export PROJECT_DIR
    export KOMMOL=${PROJECT_DIR}/target/kommol/kommol
}