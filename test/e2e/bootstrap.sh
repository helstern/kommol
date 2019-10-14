#!/usr/bin/env bash

SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(readlink -f "${SCRIPT_DIR}")"
PROJECT_DIR="$(cd ${SCRIPT_DIR}/../../ && pwd)"

function bootstrap() {
    set -o allexport
    source ${PROJECT_DIR}/.env
    set +o allexport

    export PROJECT_DIR
    export KOMMOL=${PROJECT_DIR}/target/kommol/kommol
}