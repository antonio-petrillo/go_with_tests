package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3 ; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, "Go!")
}

// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := 3 ; i > 0; i-- {
// 		fmt.Fprintf(out, "%d\n", i)
// 	}

// 	for i := 3 ; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	fmt.Fprintf(out, "Go!")
// }

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
	cSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, cSleeper)
}
