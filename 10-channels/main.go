package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	l := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://indeed.com",
	}

	c := make(chan string)

	for _, link := range l {
		go checkLink(link, c)
	}

	for i := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(i)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "appears to be down!")
		c <- l
		return
	}

	fmt.Println(l, "appears to be up!")
	c <- l
}
