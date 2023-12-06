#include <linux/bpf.h>
#include <bpf/bpf_tracing.h>
#include "ctx.h"

char __license[] SEC("license") = "Dual MIT/GPL";

#define GO_PARAM1(x) ((x)->ax)
#define GO_PARAM2(x) ((x)->bx)
#define GO_PARAM3(x) ((x)->cx)

struct greet_event {
    char param[6];
};

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 256 * 1024);
} greet_params SEC(".maps");

SEC("uprobe/go_test_greet")
int BPF_UPROBE(go_test_greet) {
    struct greet_event *e;

    /* reserve sample from BPF ringbuf */
    e = bpf_ringbuf_reserve(&greet_params, sizeof(*e), 0);
    if (!e)
        return 0;
    
    /* fill in event data */
    bpf_probe_read_str(&e->param, sizeof(e->param), (void*)GO_PARAM1(ctx));
    
    bpf_ringbuf_submit(e, 0);
    return 0;
}
