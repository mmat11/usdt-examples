//go:build linux
// +build linux

package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
	"github.com/cilium/ebpf/rlimit"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang-12 -cflags $BPF_CFLAGS bpf ./bpf/usdt.c -- -I../headers

func main() {
	stopper := make(chan os.Signal, 1)
	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)

	// Allow the current process to lock memory for eBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	// Load pre-compiled programs and maps into the kernel.
	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	sem := flag.Bool("semaphore", false, "if set, use the semaphored tracee")
	flag.Parse()

	tracee := "c/tracee/tracee.o"
	if *sem {
		tracee = "c/tracee_semaphore/tracee.o"
	}

	// Run the tracee in the background.
	cmd := exec.Command(tracee)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal(err)
		}
	}()

	// Open Executable on the tracee PID.
	e, err := link.OpenExecutable(link.WithPID(cmd.Process.Pid))
	if err != nil {
		log.Fatal(err)
	}

	// Open USDT and attach it to the ebpf program.
	u, err := e.USDT("Capp", "Cprobe", objs.Handler)
	if err != nil {
		log.Fatalf("open USDT: %v", err)
	}
	defer u.Close()

	// Open a ringbuf reader from userspace RINGBUF map.
	rd, err := ringbuf.NewReader(objs.Events)
	if err != nil {
		log.Fatalf("opening ringbuf reader: %v", err)
	}
	defer rd.Close()

	// Close the reader when the process receives a signal, which will exit
	// the read loop.
	go func() {
		<-stopper
		if err := rd.Close(); err != nil {
			log.Fatalf("closing ringbuf reader: %v", err)
		}
	}()

	log.Println("Waiting for events..")

	for {
		record, err := rd.Read()
		if err != nil {
			if errors.Is(err, ringbuf.ErrClosed) {
				log.Println("Received signal, exiting..")
				return
			}
			log.Printf("reading from reader: %v", err)
			continue
		}
		log.Printf("New event: %d\n", binary.LittleEndian.Uint32(record.RawSample))
	}
}
