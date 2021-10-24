#include "vmlinux.h"
#include "bpf_helpers.h"
#include "bpf_tracing.h"

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 256 * 1024);
} events SEC(".maps");

SEC("uprobe/Capp/Cprobe")
int handler(struct pt_regs *ctx) {
    /*
    Displaying notes found in: .note.stapsdt
    Owner                Data size 	Description
    stapsdt              0x00000030	NT_STAPSDT (SystemTap probe descriptors)
        Provider: Capp
        Name: Cprobe
        Location: 0x0000000000401162, Base: 0x0000000000402023, Semaphore: 0x0000000000000000
        Arguments: -4@12(%rsp)
    */
    __s32 arg = 0;
    bpf_probe_read_user(&arg, sizeof(arg), (void *)(PT_REGS_SP(ctx) + 12));
    bpf_ringbuf_output(&events, &arg, sizeof(__s32), 0);

    return 0;
}

char __license[] SEC("license") = "Dual MIT/GPL";
