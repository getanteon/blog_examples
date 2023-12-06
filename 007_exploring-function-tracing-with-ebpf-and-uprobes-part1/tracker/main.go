package main

import (
	"debug/elf"
	"log"
	"unsafe"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/ringbuf"
	"github.com/cilium/ebpf/rlimit"
)

const SymbolName = "main.Greet"

type GreetEvent struct {
	Msg [6]byte
}

func main() {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	path := "../program/p"
	// open in elf format in order to get the symbols
	ef, err := elf.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer ef.Close()

	ex, err := link.OpenExecutable(path)
	if err != nil {
		log.Fatal(err)
	}

	coll, err := ebpf.LoadCollection("tracker.o")
	if err != nil {
		log.Fatal(err)
	}

	greetEvents, err := ringbuf.NewReader(coll.Maps["greet_params"])
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		countMap := map[string]uint64{}
		for {
			event, err := greetEvents.Read()
			if err != nil {
				log.Fatal(err)
			}
			greetEvent := (*GreetEvent)(unsafe.Pointer(&event.RawSample[0]))
			countMap[string(greetEvent.Msg[:])]++
			log.Printf("COUNT: %v", countMap)
		}
	}()

	_, err = ex.Uprobe(SymbolName, coll.Programs["go_test_greet"], &link.UprobeOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for {
	}
}
