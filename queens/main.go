package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	up       []bool // up-facing diagonals
	down     []bool // down-facing diagonals
	rows     []bool // rows
	solution []int
)

func main() {
	n := 8
	switch len(os.Args) {
	case 2:
		k, err := strconv.Atoi(os.Args[1])
		if err != nil || k < 4 {
			fmt.Printf("%v is an invalid argument; it must be a integer >= 4\n", os.Args[1])
			os.Exit(1)
		}
		n = k
	case 1: // use default of 8
	default:
		fmt.Printf("too many arguments\n")
		os.Exit(1)
	}
	up = make([]bool, 2*n)
	down = make([]bool, 2*n)
	rows = make([]bool, n)
	solution = make([]int, n)

	queens := make([]chan bool, n)
	for c := 0; c < n; c++ {
		queens[c] = make(chan bool)
	}

	for c, ch := range queens {
		go queen(c, ch, n)
	}

	for c := 0; c >= 0; {
		queens[c] <- true
		if <-queens[c] {
			if c == n-1 {
				fmt.Println(solution)
			} else {
				c++
			}
		} else {
			c--
		}
	}
}

func queen(c int, ch chan bool, n int) {
	for {
		<-ch
		for r := 0; r < n; r++ {
			if test(r, c, n) {
				occupy(r, c, n)
				solution[c] = r + 1
				ch <- true
				<-ch
				release(r, c, n)
			}
		}
		ch <- false
	}
}

func test(r, c, n int) bool {
	return !rows[r] && !up[r-c+n-1] && !down[r+c+1]
}

func occupy(r, c, n int) {
	rows[r], up[r-c+n-1], down[r+c+1] = true, true, true
}

func release(r, c, n int) {
	rows[r], up[r-c+n-1], down[r+c+1] = false, false, false
}
