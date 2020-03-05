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

@test "returns a file from a website bucket" {
    local local_file=${BATS_TMPDIR}/radu.helstern.pdf

    STATUS=$(curl -s -w '%{http_code}' \
        --resolve radu.helstern.org:8180:127.0.0.1 --output ${local_file} \
        -H 'X-KOMMOL-STRATEGY: GCP_WEBSITE' http://radu.helstern.org:8180/cv/radu.helstern.pdf
    )
    [ "${STATUS}" = "200" ]
    rm ${local_file}
}

@test "returns the index file from a website bucket" {
    local local_file="${BATS_TMPDIR}/index-kommol-test"

    STATUS=$(curl -s -w '%{http_code}' \
        --resolve radu.helstern.org:8180:127.0.0.1 --output ${local_file} \
        -H 'X-KOMMOL-STRATEGY: GCP_WEBSITE' http://radu.helstern.org:8180/
    )
    [ "${STATUS}" = "200" ]
    rm ${local_file}
}
