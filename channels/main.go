package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // go is used only in front of function calls
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			//checkLink(l, c) // if we are directly referencing the outer value(without passing it as argument in function literal) so it will take the same copy instead we need to provide a fresh value
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up and working...")
	c <- link
}
