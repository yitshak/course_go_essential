// Get content type of sites
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string) string{
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return ""
	}
	defer resp.Body.Close()
	return resp.Header.Get("content-type")
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

  ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			ch<-returnType(url)
		}(url)
	}

  

  for range urls{
    content_type := <-ch
		fmt.Printf("%v\n", content_type)
	}
  close(ch)
	
}
