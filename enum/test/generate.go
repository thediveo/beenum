//go:generate bpf2go -go-package test test _ebpf/typedefed-enums.ebpf.c -- -I../../_headers -I../../_headers/libbpf

package test
