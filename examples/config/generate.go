//go:generate bpf2go -go-package config config _ebpf/config.ebpf.c -- -I../../_headers -I../../_headers/libbpf

package config
