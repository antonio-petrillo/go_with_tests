package racer

import (
	"fmt"
	"net/http"
	"time"
)

// not concurrent implementation
// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
// }

func ConfigurableRacer(a, b string, timeout time.Duration ) (winner string, error error)  {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed after %q out waiting for %s and %s", timeout, a, b)
	}
}

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, 10 * time.Second)
}

func ping(url string) chan struct{} {
	ch := make(chan struct {})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
