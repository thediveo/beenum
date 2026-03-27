#!/usr/bin/env bash
set -e

# Version of libbpf to fetch headers from
LIBBPF_VERSION=1.7.0

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)

# The headers we want
prefix=libbpf-"${LIBBPF_VERSION}"
headers=(
    "$prefix"/LICENSE.BSD-2-Clause
    "$prefix"/src/bpf_core_read.h
    "$prefix"/src/bpf_endian.h
    "$prefix"/src/bpf_helper_defs.h
    "$prefix"/src/bpf_helpers.h
    "$prefix"/src/bpf_tracing.h
)

# Fetch libbpf release and extract the desired headers
URL="https://github.com/libbpf/libbpf/archive/refs/tags/v${LIBBPF_VERSION}.tar.gz"
echo "downloading ${URL}..."
curl -sL "${URL}" | tar -xz -C "${SCRIPT_DIR}" --xform='s#.*/##' "${headers[@]}"
echo "updated files"

# Update also version x.y.z in README.md
sed -i -E "/[0-9]+\.[0-9]+\.[0-9]+/s//${LIBBPF_VERSION}/" "${SCRIPT_DIR}/README.md"
echo "updated README.md"

echo "done"
