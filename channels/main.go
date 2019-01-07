package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://golang.org",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "link might be down")
		return
	}

	fmt.Println(link, "is up!")
	return
}
