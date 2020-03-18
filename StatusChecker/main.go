package main

import (
	"fmt"
	"net/http"
	"time"
)

/**
Concurrency - more than one routine can be in progress on a single core
Parallelism - routines run in parallel on multiple cores.
*/
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com.au",
	}
	c := make(chan string)

	for _, link := range links {
		// go before a function call -> spawns a new routine for each call
		go goToLink(link, c)

	}

	for l := range c {
		go func(l1 string) {
			// do not add this time duration for sleep on the main routine.
			time.Sleep(time.Second * 5)
			goToLink(l1, c)
		}(l) // function call
	}
}

func goToLink(s string, c chan string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println(s, " might be down!")
	} else {
		fmt.Println(s, " is up!")
	}

	c <- s
}
