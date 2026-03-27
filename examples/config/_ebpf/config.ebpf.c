// SPDX-License-Identifier: MIT

//go:build ignore

#include <cilium-ebpf/common.h>
#include <libbpf/bpf_helpers.h>

char __license[] SEC("license") = "MIT";

// For background information, please refer to:
// https://github.com/cilium/ebpf/discussions/1567

typedef enum {
    MAX_BEES = 42,
    DEFAULT_FOOBAR = (1ULL << 48),
} defines;

// Coerce the compiler to consider the enum as being used so that it emits the
// desired enum type information into the ELF.
static __attribute__((used)) defines __defines;

