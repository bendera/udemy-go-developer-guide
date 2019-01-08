package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://golang.org",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "link might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
	return
}
