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

	// nm command can be used to get the symbols as well
	// symbols, err := ef.Symbols()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// print the symbols
	// for _, s := range symbols {
	// 	log.Printf("%s: %x", s.Name, s.Value)
	// }

	ex, err := link.OpenExecutable(path)
	if err != nil {
		log.Fatal(err)
	}

	coll, err := ebpf.LoadCollection("tracker.o")
	if err != nil {
		log.Fatal(err)
	}

	// _, err = ringbuf.NewReader(coll.Maps["greet_events"])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for x, prog := range coll.Programs {
		log.Printf("PROG %s: %s", x, prog.Type)
	}

	for x, map2 := range coll.Maps {
		log.Printf("MAP %s: %s", x, map2.String())
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
	// defer l.Close()
}
