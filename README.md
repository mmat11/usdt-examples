# USDT examples using cilium/ebpf

This repo contains examples using a [PoC implementation of USDT probes](https://github.com/cilium/ebpf/compare/master...mmat11:matt/usdt)


### Libstapsdt

[libstapsdt](libstapsdt/) instruments a Python script with [python-stapsdt](https://github.com/sthima/python-stapsdt) and shows how to send the probe args to userspace via a ringbuffer:

```console
(venv) ~/g/m/usdt-examples *main> go run -exec sudo ./libstapsdt/
TRACEE (0): fired 77f47f5f-51b1-4a82-827a-d8d66f590118
2021/10/23 22:32:48 Waiting for events..
TRACEE (1): fired 651e17ab-57ac-460e-a572-d3ee0b286285
2021/10/23 22:32:48 New event (1): 651e17ab-57ac-460e-a572-d3ee0b286285
TRACEE (2): fired acfca6f3-25c1-49f5-9c5e-05e89dd6b40a
2021/10/23 22:32:48 New event (2): acfca6f3-25c1-49f5-9c5e-05e89dd6b40a
TRACEE (3): fired 74465de8-35f1-4cf4-bf7d-19cdd1185504
2021/10/23 22:32:49 New event (3): 74465de8-35f1-4cf4-bf7d-19cdd1185504
TRACEE (4): fired c3256e23-956d-462c-8e2d-c9f042c47c79
2021/10/23 22:32:49 New event (4): c3256e23-956d-462c-8e2d-c9f042c47c79
TRACEE (5): fired ad9359d3-84dc-42c2-89ba-871a2b0e5b38
2021/10/23 22:32:50 New event (5): ad9359d3-84dc-42c2-89ba-871a2b0e5b38
TRACEE (6): fired 3e5eb5c3-0daa-4d25-a970-632f2d1e968c
2021/10/23 22:32:50 New event (6): 3e5eb5c3-0daa-4d25-a970-632f2d1e968c
TRACEE (7): fired 358e5f89-2655-4d60-a7e8-5f72ac928fb5
2021/10/23 22:32:51 New event (7): 358e5f89-2655-4d60-a7e8-5f72ac928fb5
TRACEE (8): fired a7de7480-36fc-4340-af8a-0d82336ab81f
2021/10/23 22:32:51 New event (8): a7de7480-36fc-4340-af8a-0d82336ab81f
TRACEE (9): fired cde0d029-0839-44ae-83c9-b1864fd8d506
2021/10/23 22:32:52 New event (9): cde0d029-0839-44ae-83c9-b1864fd8d506
TRACEE (10): fired a38aefcb-9259-4a71-b78e-75d5af8b233e
2021/10/23 22:32:52 New event (10): a38aefcb-9259-4a71-b78e-75d5af8b233e
```

### Python (.so SDT notes)

TODO

### Dtrace

TODO

### Dtrace with semaphore

TODO
