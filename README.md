# USDT examples using cilium/ebpf

This repo contains examples using a [PoC implementation of USDT probes](https://github.com/cilium/ebpf/compare/master...mmat11:matt/usdt)

To regenerate all objects, run `make all`.

### Python (stapsdt)

Dependencies: [libstapsdt](https://github.com/linux-usdt/libstapsdt), [python-stapsdt](https://github.com/linux-usdt/python-stapsdt)

[python-stapsdt](python-stapsdt/) instruments a Python script with python-stapsdt have a bpf program which forwards the probe args to userspace:


```console
~ go run -exec sudo ./python-stapsdt/

INFO:root:counter=0, uuid=d4e857f9-3c80-4017-9f16-7623cac2e5d4
2021/10/23 23:45:25 Waiting for events..
2021/10/23 23:45:26 New event (1): 8364b259-018f-4770-9b0d-a151ae69bcb8
INFO:root:counter=1, uuid=8364b259-018f-4770-9b0d-a151ae69bcb8
2021/10/23 23:45:26 New event (2): 070e5caa-efcc-41b0-9c55-c2cfdaa5e49b
INFO:root:counter=2, uuid=070e5caa-efcc-41b0-9c55-c2cfdaa5e49b
2021/10/23 23:45:27 New event (3): d8bf5088-7db1-438d-9aa1-4b3b8f6fa5c5
INFO:root:counter=3, uuid=d8bf5088-7db1-438d-9aa1-4b3b8f6fa5c5
2021/10/23 23:45:27 New event (4): 679c29d5-516a-4693-9683-324f17619b24
INFO:root:counter=4, uuid=679c29d5-516a-4693-9683-324f17619b24
2021/10/23 23:45:28 New event (5): b8581c82-29e6-4b87-bec3-de3209821953
INFO:root:counter=5, uuid=b8581c82-29e6-4b87-bec3-de3209821953
INFO:root:counter=6, uuid=da0717e9-1982-418e-8738-2fc3dbbfae7e
2021/10/23 23:45:28 New event (6): da0717e9-1982-418e-8738-2fc3dbbfae7e
2021/10/23 23:45:29 New event (7): 3111c754-760a-4835-b5a4-c6f8e127114c
INFO:root:counter=7, uuid=3111c754-760a-4835-b5a4-c6f8e127114c
```

### C

[c-simple](c-simple/) demonstrates how to read arguments exported via SDT notes on a C executable:

```console
~ go run -exec sudo ./c-simple/

tracee: firing 0
2021/10/24 16:24:29 Waiting for events..
tracee: firing 1
2021/10/24 16:24:30 New event: 1
tracee: firing 2
2021/10/24 16:24:31 New event: 2
tracee: firing 3
2021/10/24 16:24:32 New event: 3
tracee: firing 4
2021/10/24 16:24:33 New event: 4
tracee: firing 5
2021/10/24 16:24:34 New event: 5
```

### C (semaphore)

TODO

### Python (.so SDT notes)

TODO
