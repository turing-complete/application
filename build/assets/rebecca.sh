#!/bin/bash

set -e

function join() {
    local IFS="${1}"
    shift
    echo "$*"
}

root="$(dirname "$(dirname "${BASH_SOURCE[0]}")")"
problem="$(basename "${BASH_SOURCE[0]}")"
problem="${problem%.*}"

program="${problem}-$(join 'x' $@)"
output="${root}/results/${program}"

echo $@
