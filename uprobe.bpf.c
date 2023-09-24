// SPDX-License-Identifier: GPL-2.0 OR BSD-3-Clause
/* Copyright (c) 2020 Facebook */
/* Copyright (c) 2023 Hengqi Chen */
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

SEC("uprobe/./go-symbols/go-symbols:github.com/cilium/ebpf/internal.(*Deque[go.shape.interface { Format(fmt.State, int32); TypeName() string; github.com/cilium/ebpf/btf.copy() github.com/cilium/ebpf/btf.Type }]).Grow+0")
int BPF_UPROBE(deque_grow)
{
	bpf_printk("deque_grow called");
	return 0;
}

char LICENSE[] SEC("license") = "Dual BSD/GPL";
