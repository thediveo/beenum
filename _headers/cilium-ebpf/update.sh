#!/usr/bin/env bash
set -e

# Version of cilium/ebpf to fetch example headers from
CILIUM_EBPF_VERSION=0.21.0 # no(!) "v"

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

# The headers we want
prefix=ebpf-"${CILIUM_EBPF_VERSION}/examples/headers"
headers=(
    "$prefix"/LICENSE.BSD-2-Clause
    "$prefix"/common.h
)

# Fetch libbpf release and extract the desired headers
URL="https://github.com/cilium/ebpf/archive/refs/tags/v${CILIUM_EBPF_VERSION}.tar.gz"
echo "downloading ${URL}..."
curl -sL "${URL}" | tar -xz -C "${SCRIPT_DIR}" --xform='s#.*/##' "${headers[@]}"
echo "updated files"

# Update also version x.y.z in README.md
sed -i -E "/[0-9]+\.[0-9]+\.[0-9]+/s//${CILIUM_EBPF_VERSION}/" "${SCRIPT_DIR}/README.md"
echo "updated README.md"

echo "done"
