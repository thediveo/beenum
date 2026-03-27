// SPDX-License-Identifier: MIT

//go:build ignore

#include <cilium-ebpf/common.h>
#include <libbpf/bpf_helpers.h>

char __license[] SEC("license") = "MIT";

// For background information, please refer to:
// https://github.com/cilium/ebpf/discussions/1567

typedef enum /* : unsigned long long */ {
    MAX_BEES = 42,
    DEFAULT_FOOBAR = (1ll << 48),
} defines;

// Coerce the compiler to consider the enum as being used so that it emits the
// desired enum type information into the ELF.
static __attribute__((used)) defines __defines;

typedef enum: signed char {
    NOT_MAX = -127,
} shortdefines;
static __attribute__((used)) shortdefines __shortdefines;

// Some other typedef to throw off any enum scents...
typedef struct { int b; } foobar;
static __attribute__((used)) foobar __foobar = { 42 };
