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

if [ -z "${BENCHMARKS_ROOT}" ]; then
  die 'BENCHMARKS_ROOT should be defined'
fi

if [ -z "${STUDIO_ROOT}" ]; then
  die 'STUDIO_ROOT should be defined'
fi

if ! has bc; then
  die 'expected bc to be installed'
fi

if ! has sqlite3; then
  die 'expected SQLite to be installed'
fi

if ! has vips; then
  die 'expected VIPS to be installed'
fi

if [ "$#" -ne 1 ] && [ "$#" -ne 2 ]; then
  die 'expected one or two input variables'
fi

x="${1}"
y="${1}"

if [ "$#" -eq 2 ]; then
  y="${2}"
fi

x="$(printf '%.10f' "$(echo 1.0 + "${x}" | bc)")"
y="$(printf '%.10f' "$(echo 1.0 + "${y}" | bc)")"

root="$(cd "$(dirname "$(dirname "${BASH_SOURCE[0]}")")" && pwd)"
input="${root}/assets"
output="${root}/results"

problem="$(basename "${BASH_SOURCE[0]}")"
problem="${problem%.*}"
scenario="${problem}-${x}x${y}"

mkdir -p "${output}/${scenario}"

picture="${output}/${scenario}/${problem}.v"
if [ ! -e "${picture}" ]; then
  vips resize "${input}/${problem}.tif" "${picture}.v" "${x}" --vscale "${y}"
  vips im_Lab2LabQ "${picture}.v" "${picture}"
fi

input_size="sim${x}x${y}"
program="parsec-vips-${input_size}"
parsec="${BENCHMARKS_ROOT}/parsec/parsec-2.1"

echo """#!/bin/bash
run_desc='Cursom input for performance analysis with simulators (${input_size})'
""" > "${parsec}/config/${input_size}.runconf"

echo """#!/bin/bash
run_exec='bin/vips'
export IM_CONCURRENCY=\${NTHREADS}
run_args='im_benchmark \"${picture}\" output.v'
""" > "${parsec}/pkgs/apps/vips/parsec/${input_size}.runconf"

if [ ! -e "${output}/${scenario}/.${program}" ]; then
  OUTPUT_DIR="${output}/${scenario}" make -C "${STUDIO_ROOT}" \
    setup "record-${program}" 2>&1 > "${output}/${scenario}/${program}.log"
fi

if [ ! -e "${output}/${scenario}/.${program}" ]; then
  die 'failed to simulate'
fi

query='SELECT 1e-3 * SUM(`dynamic_power`) FROM `dynamic`;'
echo "${query}" | sqlite3 "${output}/${scenario}/${program}.sqlite3"
