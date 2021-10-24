CLANG ?= clang-12
CFLAGS := -O2 -Wall -Werror

all: clean c bpf

c:
	$(CLANG) $(CFLAGS) ./c-simple/tracee/tracee.c -o ./c-simple/tracee/tracee.o

bpf:
	BPF_CFLAGS="-D__x86_64__ $(CFLAGS)" go generate ./...

clean:
	find . -name "*.o" -type f -delete
	find . -name "*_bpfel.go" -type f -delete
	find . -name "*_bpfeb.go" -type f -delete
