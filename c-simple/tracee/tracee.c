#include <sys/sdt.h>
#include <stdio.h>
#include <unistd.h>

int main() {
    for (int i = 0;; i++) {
        printf("tracee: firing %d\n", i);
        DTRACE_PROBE1(Capp, Cprobe, i);
        sleep(1);
    }
    return 0;
}
