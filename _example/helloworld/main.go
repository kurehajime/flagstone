package main

import (
	"flag"
	"fmt"

	"github.com/kurehajime/flagstone"
)

var who *string

func main() {
	who = flag.String("who", "world", "say hello to ...")
	flag.Parse()

	flagstone.Launch("helloworld", "flagstone sample")

	fmt.Println("hello " + *who + "!")
}
