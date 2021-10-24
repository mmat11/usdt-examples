#include <unistd.h>
#include <stdio.h>
#include "provider.h"

int main() {
    for (int i = 0;; i++) {
        printf("tracee (with semaphore): run %d\n", i);
        if (CAPP_CPROBE_ENABLED()) {
            CAPP_CPROBE(i);
        }
        usleep(500000); // 0.5s
    }

    return 0;
}
