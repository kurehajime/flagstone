package main

import (
	"flag"
	"fmt"

	"github.com/kurehajime/flagstone"
)

var (
	start *int
	end   *int
	fizz  *string
	buzz  *string
)

func main() {
	start = flag.Int("start", 1, "The starting index for Fizz buzz")
	end = flag.Int("end", 100, "The ending index for Fizz buzz")

	fizz = flag.String("fizz", "Fizz", "Keyword to replace Fizz")
	buzz = flag.String("buzz", "Buzz", "Keyword to replace Buzz")

	flag.Parse()

	flagstone.SetSort([]string{"start", "end", "fizz", "buzz"})
	flagstone.SetSilent(true)
	flagstone.SetUseNonFlagArgs(false)
	flagstone.SetPort(8008)
	if _, ok := flagstone.Launch("fizzbuzz", `Fizz buzz is a group word game for children to teach them about division.Players take turns to count incrementally, replacing any number divisible by three with the word "fizz", and any number divisible by five with the word "buzz". ---Wikipedia:Fizz buzz`); ok {
		fizzbuzz()
	}
}

func fizzbuzz() {
	for i := *start; i <= *end; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(*fizz + *buzz)
		} else if i%3 == 0 {
			fmt.Println(*fizz)
		} else if i%5 == 0 {
			fmt.Println(*buzz)
		} else {
			fmt.Println(i)
		}
	}
}
