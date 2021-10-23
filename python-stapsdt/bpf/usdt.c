#define UUIDSZ 36 + 1

#include "vmlinux.h"
#include "bpf_helpers.h"
#include "bpf_tracing.h"

struct event_t {
    __s32 counter;
    char uuid[UUIDSZ];
};

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 256 * 1024);
} events SEC(".maps");

SEC("uprobe/pyapp/pyprobe")
int handler(struct pt_regs *ctx) {
    struct event_t event = {};

    bpf_probe_read_user_str(event.uuid, UUIDSZ, (void *)PT_REGS_PARM1(ctx));
    event.counter = PT_REGS_PARM2(ctx);
    bpf_ringbuf_output(&events, &event, sizeof(struct event_t), 0);

    return 0;
}

char __license[] SEC("license") = "Dual MIT/GPL";
