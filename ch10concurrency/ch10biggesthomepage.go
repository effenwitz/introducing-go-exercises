package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	fmt.Println("Homepage with biggest size")

	urls := []string{
		"https://www.apple.com",
		"https://www.amazon.de",
		"https://www.flipkart.com",
		"https://www.microsoft.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			bs, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)

	}

	var biggest HomePageSize
	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("Biggest home page: ", biggest.URL)
}
