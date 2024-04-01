#!/bin/bash -e

SRC_DIR="$(dirname "$(realpath "$0")")"

for cmd in "${SRC_DIR}/cmd/"*; do
  tool="$(basename "${cmd}")"
  pushd "${cmd}" > /dev/null
  go build -o "${tool}"
  popd > /dev/null
done
