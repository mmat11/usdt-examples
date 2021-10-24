# USDT examples using cilium/ebpf

This repo contains examples using a [PoC implementation of USDT probes](https://github.com/cilium/ebpf/compare/master...mmat11:matt/usdt)

To regenerate all objects, run `make all`.

### Python (stapsdt)

Dependencies: [libstapsdt](https://github.com/linux-usdt/libstapsdt), [python-stapsdt](https://github.com/linux-usdt/python-stapsdt)

[python-stapsdt](python-stapsdt/) instruments a Python script with python-stapsdt have a bpf program which forwards the probe args to userspace:


```console
~ go run -exec sudo ./python-stapsdt/

INFO:root:counter=0, uuid=ffba5aac-e917-48aa-b1c4-be330db99069
2021/10/24 22:05:59 Waiting for events..
2021/10/24 22:06:00 New event: 1 ad5b87bb-9a4d-4b7b-99b4-802c6dab5ef2
INFO:root:counter=1, uuid=ad5b87bb-9a4d-4b7b-99b4-802c6dab5ef2
INFO:root:counter=2, uuid=ff8c3113-008c-4349-823a-84459ec8f33d
2021/10/24 22:06:00 New event: 2 ff8c3113-008c-4349-823a-84459ec8f33d
2021/10/24 22:06:01 New event: 3 0a760cdf-1166-4bc5-8777-51e6be8bf826
INFO:root:counter=3, uuid=0a760cdf-1166-4bc5-8777-51e6be8bf826
2021/10/24 22:06:01 New event: 4 48b26f7b-6222-408f-b6d3-fdcbc5fbcd45
```

### Python (builtin SDT notes with semaphore)

[python](python/) uses Python builtin SDT notes (https://docs.python.org/3/howto/instrumentation.html) to trace function calls:

```console
~ go run -exec sudo ./python/

tracee: floordiv(100,5) = 20
2021/10/24 21:18:21 Waiting for events..
tracee: floordiv(100,3) = 33
2021/10/24 21:18:21 New event: /usr/lib64/python3.9/random.py:randint()
2021/10/24 21:18:21 New event: /usr/lib64/python3.9/random.py:randrange()
2021/10/24 21:18:21 New event: /usr/lib64/python3.9/random.py:_randbelow_with_getrandbits()
2021/10/24 21:18:21 New event: /home/matt/github/mmat11/usdt-examples/python/tracee.py:floordiv()
tracee: floordiv(100,9) = 11
2021/10/24 21:18:22 New event: /usr/lib64/python3.9/random.py:randint()
2021/10/24 21:18:22 New event: /usr/lib64/python3.9/random.py:randrange()
2021/10/24 21:18:22 New event: /usr/lib64/python3.9/random.py:_randbelow_with_getrandbits()
2021/10/24 21:18:22 New event: /home/matt/github/mmat11/usdt-examples/python/tracee.py:floordiv()
tracee: floordiv(100,1) = 100
2021/10/24 21:18:22 New event: /usr/lib64/python3.9/random.py:randint()
2021/10/24 21:18:22 New event: /usr/lib64/python3.9/random.py:randrange()
2021/10/24 21:18:22 New event: /usr/lib64/python3.9/random.py:_randbelow_with_getrandbits()
2021/10/24 21:18:22 New event: /home/matt/github/mmat11/usdt-examples/python/tracee.py:floordiv()
tracee: floordiv(100,5) = 20
```

### C

[c](c/) demonstrates how to read arguments exported via SDT notes on a C executable:

```console
~ go run -exec sudo ./c/

tracee: run 0
2021/10/24 21:48:32 Waiting for events..
tracee: run 1
2021/10/24 21:48:32 New event: 1
tracee: run 2
2021/10/24 21:48:32 New event: 2
tracee: run 3
2021/10/24 21:48:33 New event: 3
```

Another example which uses a different tracee can be run by setting the `-semaphore` flag:

```console
~ go run -exec sudo ./c/ -semaphore

tracee (with semaphore): run 0
2021/10/24 21:55:36 Waiting for events..
tracee (with semaphore): run 1
2021/10/24 21:55:37 New event: 1
tracee (with semaphore): run 2
2021/10/24 21:55:37 New event: 2
tracee (with semaphore): run 3
2021/10/24 21:55:38 New event: 3
```
