package main

import (
	"flag"
	"flagstone"
	"fmt"
)

var who *string

func main() {
	who = flag.String("who", "world", "say hello to ...")
	flag.Parse()

	flagstone.Launch("helloworld", "flagstone sample")

	fmt.Println("hello " + *who + "!")
}
