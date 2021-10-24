CLANG ?= clang-12
CFLAGS := -O2 -Wall -Werror

all: clean c bpf #c_sem

c:
	$(CLANG) $(CFLAGS) ./c-simple/tracee/tracee.c -o ./c-simple/tracee/tracee.o

#c_sem:
#	dtrace -G -s usdt_sem_provider.d -o usdt_sem_provider.so
#	dtrace -h -s usdt_sem_provider.d -o usdt_sem_provider.h
#	$(CLANG) $(CFLAGS) usdt_sem.c usdt_sem_provider.so -o usdt_sem.elf

bpf:
	BPF_CFLAGS="-D__x86_64__ $(CFLAGS)" go generate ./...

clean:
	rm -f -- **/*.o
	rm -f -- **/*_bpfel.go
	rm -f -- **/*_bpfeb.go
