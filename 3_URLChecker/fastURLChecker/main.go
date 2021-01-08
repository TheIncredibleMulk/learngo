package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Status code > 400\nRequest Failed")

func main() {
	var results = map[string]string{}
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.candelacontrols.com/",
		"https://www.andrewmulkey.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		hitURL(url)
	}

}

func hitURL(url string, c chan result) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- result{url: url, status: "FAILED"}
	} else {

	}
	if resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
