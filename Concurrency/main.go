package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()        // defer runs cancel() when main exits to clean up resources
	var wg sync.WaitGroup // wg is used to wait for all goroutines to finish

	ch := make(chan string, len(urls)) //Creates a buffered channel of strings with capacity equal to the number of URLs.
	for _, url := range urls {
		wg.Add(1) // add 1 to the wait group for each URL
		go func(u string) {
			statusCode, _ := checkURL(ctx, u)
			printStatus(u, statusCode)
			wg.Done() // decrement the wait group by 1
			ch <- u
		}(url) // pass the url to the goroutine
	}
	wg.Wait() // wait for all goroutines to finish

	// Check if context was cancelled (timeout occurred)
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Context cancelled: Timeout exceeded")
	} else {
		fmt.Println("✓ All URLs checked")
	}
}

func printStatus(url string, statusCode int) {
	if statusCode == 200 {
		fmt.Println("URL: ", url, "is up and running with status code: ", statusCode)
	} else {
		fmt.Println("URL: ", url, "is down with status code: ", statusCode)
	}
}

func checkURL(ctx context.Context, url string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
