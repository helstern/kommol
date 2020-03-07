#!/usr/bin/env bash

SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(readlink -f "${SCRIPT_DIR}")"
PROJECT_DIR="$(cd ${SCRIPT_DIR}/../../ && pwd)"

function tool_test_e2e_bootstrap() {
    set -o allexport
    source ${PROJECT_DIR}/.env
    set +o allexport

    export PROJECT_DIR
    export KOMMOL=${PROJECT_DIR}/target/kommol/kommol
}

function tool_test_e2e_start() {
    tool_test_e2e_bootstrap

#    docker run --rm --name kommol-haproxy \
#        --network=host \
#        -e E2E_WEBSITE_BUCKET=${E2E_WEBSITE_BUCKET} \
#        -v ${SCRIPT_DIR}/haproxy:/usr/local/etc/haproxy:ro \
#        haproxy:2.0.7-alpine >> ${SCRIPT_DIR}/log.txt

    ${KOMMOL} -bind 127.0.0.1:8180 -log-level debug -gcp.credentials ${GOOGLE_APPLICATION_CREDENTIALS} &

    docker run --detach --rm --name kommol-haproxy \
        --network=host \
        -e E2E_WEBSITE_BUCKET=${E2E_WEBSITE_BUCKET} \
        -v ${SCRIPT_DIR}/haproxy:/usr/local/etc/haproxy:ro \
        haproxy:2.0.7-alpine >> ${SCRIPT_DIR}/log.txt


}

function tool_test_e2e_stop() {
    docker stop kommol-haproxy
    pkill -f ${PROJECT_DIR}/target/kommol/kommol > /dev/null
}

function tool_test_e2e_request()
{
    curl --resolve radu.helstern.org:80:127.0.0.1 \
        -H 'X-KOMMOL-STRATEGY: GCP_WEBSITE' http://radu.helstern.org:80${1} \

}

if [[ $#  -eq 0 ]]; then
    echo "no command given"
    exit 1
fi

CMD="tool_test_e2e_${1}"; shift
"$CMD" "$@"
