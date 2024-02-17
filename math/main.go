package main

import (
	"os"
	"time"
)

func main() {
	t := time.Now()
	SVGWriter(os.Stdout, t)
}
