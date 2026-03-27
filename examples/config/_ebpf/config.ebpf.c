// SPDX-License-Identifier: MIT

//go:build ignore

#include <cilium-ebpf/common.h>
#include <libbpf/bpf_helpers.h>

char __license[] SEC("license") = "MIT";

typedef enum {
    MAX_BEES = 42,
    DEFAULT_FOOBAR = (1 << 48),
} configuration;
