#!/bin/bash

set -e

function die() {
  echo "Error: ${1}." 1>&2
  exit 1
}

function has() {
  if hash "${1}" 2> /dev/null; then
    return 0
  else
    return 1
  fi
}

if ! has vips; then
  die 'vips should be installed'
fi

if [ "$#" -ne 1 ] && [ "$#" -ne 2 ]; then
  die 'expected one or two input variables'
fi

x="${1}"
y="${1}"

if [ "$#" -eq 2 ]; then
  y="${2}"
fi

root="$(dirname "$(dirname "${BASH_SOURCE[0]}")")"
input="${root}/assets"
output="${root}/results"

problem="$(basename "${BASH_SOURCE[0]}")"
problem="${problem%.*}"

suffix="${x}-${y}"
program="${problem}-${suffix}"
output="${output}/${program}"

mkdir -p "${output}"

if [ ! -e "${output}/${program}.png" ]; then
  vips resize "${input}/${problem}.png" "${output}/${program}.png" "${x}" --vscale "${y}"
fi

echo $@
