package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://ifood.com.br",
		"http://loggi.com.br",
		"http://gympass.com.br",
		"http://amazon.com.br",
		"http://nubank.com.br",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
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

	fmt.Println(link, "is up!")
	c <- link
}
