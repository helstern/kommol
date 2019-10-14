#!/usr/bin/env bats

source ${BATS_TEST_DIRNAME}/bootstrap.sh
bootstrap

function setup() {

    env > ${BATS_TEST_DIRNAME}/log.txt

    ${KOMMOL} -bind 127.0.0.1:8180 -gcp.credentials ${GOOGLE_APPLICATION_CREDENTIALS} &

    docker run --detach --rm --name kommol-haproxy \
        --network=host \
        -e E2E_WEBSITE_BUCKET=E2E_WEBSITE_BUCKET \
        -v ${BATS_TEST_DIRNAME}/haproxy:/usr/local/etc/haproxy:ro \
        haproxy:2.0.7-alpine >> ${BATS_TEST_DIRNAME}/log.txt
}

function teardown() {

    docker stop kommol-haproxy
    pkill -f ${PROJECT_DIR}/target/kommol/kommol > /dev/null
}

@test "download from a public bucket" {
    local local_file=${BATS_TMPDIR}/"$(basename ${E2E_WEBSITE_FILE})"

    STATUS=$(curl -s -w '%{http_code}' --resolve ${E2E_WEBSITE_BUCKET}:80:127.0.0.1 -o ${local_file} http://${E2E_WEBSITE_BUCKET}${E2E_WEBSITE_FILE})
    [ "${STATUS}" = "200" ]
    rm ${local_file}
}