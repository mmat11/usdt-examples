#include <sys/sdt.h>
#include <stdio.h>
#include <unistd.h>

int main() {
    for (int i = 0;; i++) {
        printf("tracee: run %d\n", i);
        DTRACE_PROBE1(Capp, Cprobe, i);
        usleep(500000); // 0.5s
    }
    return 0;
}
