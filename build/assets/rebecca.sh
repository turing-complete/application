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

if [ ! -e "${output}/${program}.sqlite3" ]; then
  query="""
    CREATE TABLE ``dynamic`` (``dynamic_power`` REAL);
    INSERT INTO ``dynamic`` VALUES (${x});
  """
  echo "${query}" | sqlite3 "${output}/${program}.sqlite3"
fi

query='SELECT 1e-3 * SUM(`dynamic_power`) FROM `dynamic`;'
echo "${query}" | sqlite3 "${output}/${program}.sqlite3"
