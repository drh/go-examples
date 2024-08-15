package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

var (
	NEdiag []bool //  /-diagonals: for all squares (r,c) on these diagonals,
	// the differences r-c are constant. NEdiag[r-c+n-1] is true if the square (r,c) is safe.
	NWdiag []bool //  \-diagonals: for all squares (r,c) on these diagonals,
	// the sums r+c are constant. NWdiag[r+c] is true if the square (r,c) is safe.
	row      []bool //  ranks: row[r] is true if any square on rank r is safe.
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
	NEdiag = slices.Repeat([]bool{true}, 2*n-1)
	NWdiag = slices.Repeat([]bool{true}, 2*n-1)
	row = slices.Repeat([]bool{true}, n)
	solution = make([]int, n)

	queens := make([]chan bool, n)
	for c := 0; c < n; c++ {
		ch := make(chan bool)
		queens[c] = ch
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
			if testSquare(r, c, n) {
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

func testSquare(r, c, n int) bool {
	return row[r] && NEdiag[r-c+n-1] && NWdiag[r+c]
}

func occupy(r, c, n int) {
	row[r], NEdiag[r-c+n-1], NWdiag[r+c] = false, false, false
}

func release(r, c, n int) {
	row[r], NEdiag[r-c+n-1], NWdiag[r+c] = true, true, true
}
