package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 10
	switch len(os.Args) {
	case 2:
		k, err := strconv.Atoi(os.Args[1])
		if err != nil || k < 1 {
			fmt.Printf("%v is an invalid argument; it must be a integer > 1\n", os.Args[1])
			os.Exit(1)
		}
		n = k
	case 1: // use default of 10
	default:
		fmt.Printf("too many arguments\n")
		os.Exit(1)
	}
	ch := make(chan int, 1)
	go source(ch)
	for i := 0; i < n; i++ {
		prime := <-ch
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(prime)
		out := make(chan int, 1)
		go filter(prime, ch, out)
		ch = out
	}
	fmt.Println()
}

func source(out chan<- int) {
	out <- 2
	for i := 3; ; i += 2 {
		out <- i
	}
}

func filter(prime int, in <-chan int, out chan<- int) {
	for {
		n := <-in
		if n%prime != 0 {
			out <- n
		}
	}
}
