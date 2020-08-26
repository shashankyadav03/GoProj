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
		"http://twitter.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)

	for _, link := range links {
		go check(c, link)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			check(c, link)
		}(l)
	}
	// resp, err := http.Get("http://google.com/")
	// fmt.Println(resp.Status)
	// if err != nil {
	// 	fmt.Println("error")
	// }

}

func check(c chan string, link string) bool {

	resp, err := http.Get(link)
	if resp.Status == "200 OK" {
		fmt.Println(link, " Working ")
		c <- link
		return true
	}
	fmt.Println(err)
	return false
}
