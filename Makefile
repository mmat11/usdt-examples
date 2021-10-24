CLANG ?= clang-12
CFLAGS := -O2 -Wall -Werror

all: clean c_simple c_sem bpf

lint:
	black .
	isort .

c_simple:
	$(CLANG) $(CFLAGS) c/tracee/tracee.c -o c/tracee/tracee.o

c_sem:
	dtrace -G -s c/tracee_semaphore/provider.d -o c/tracee_semaphore/provider.so
	dtrace -h -s c/tracee_semaphore/provider.d -o c/tracee_semaphore/provider.h
	$(CLANG) $(CFLAGS) c/tracee_semaphore/tracee.c c/tracee_semaphore/provider.so -o  c/tracee_semaphore/tracee.o

bpf:
	BPF_CFLAGS="-D__x86_64__ $(CFLAGS)" go generate ./...

clean:
	find . -name "*.o" -type f -delete
	find . -name "*.so" -type f -delete
	find . -name "*_bpfel.go" -type f -delete
	find . -name "*_bpfeb.go" -type f -delete
