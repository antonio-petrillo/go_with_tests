package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// func CheckWebsite (wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		results[url] = wc(url)
// 	}

// 	return results
// }

func CheckWebsite (wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	// should work for go > 1.20
	// for _, url := range urls {
	// 	go func() {
	// 		results[url] = wc(url)
	// 	} ()
	// }
	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
			resultChannel <- result{u, wc(u)}
		} (url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
