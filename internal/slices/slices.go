package slices

import "math/bits"

// Repeat returns a new slice that repeats the provided slice the given number of times.
// The result has length and capacity len(x) * count.
// The result is never nil.
// Repeat panics if count is negative or if the result of (len(x) * count)
// overflows.
// NB: to appear in 1.23.0 stdlib slices package.
func Repeat[S ~[]E, E any](x S, count int) S {
	if count < 0 {
		panic("cannot be negative")
	}
	const maxInt = ^uint(0) >> 1
	if hi, lo := bits.Mul(uint(len(x)), uint(count)); hi > 0 || lo > maxInt {
		panic("the result of (len(x) * count) overflows")
	}
	newslice := make(S, len(x)*count)
	n := copy(newslice, x)
	for n < len(newslice) {
		n += copy(newslice[n:], newslice[:n])
	}
	return newslice
}
