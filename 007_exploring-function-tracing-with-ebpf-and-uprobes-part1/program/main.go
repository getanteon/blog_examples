package main

import (
	"fmt"
	"math/rand"
	"time"
)

//go:noinline
func Greet(name string) {
	fmt.Println("Hello, " + name)
}

func main() {
	names := []string{"Mauro", "Lucas", "Kerem"}
	tick := time.Tick(1 * time.Second)

	for range tick {
		Greet(names[rand.Intn(len(names))])
	}
}
