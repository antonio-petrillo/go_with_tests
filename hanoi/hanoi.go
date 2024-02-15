package main

import (
	"flag"
	"fmt"
)

var nFlag = flag.Int("n", 3, "number of disk to use (suggested max 15, default 3)")

func hanoi(n int, src, aux, dest rune) {
	if n == 0 {
		return
	}
	hanoi(n-1, src, dest, aux)
	fmt.Printf("Move %d from %c to %c\n", n, src, dest)
	hanoi(n-1, aux, src, dest)
}

func main() {
	flag.Parse()
	hanoi(*nFlag, 'A', 'B', 'C')
	return
}
