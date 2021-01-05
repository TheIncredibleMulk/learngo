package main

import (
	"net/url"
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Status code > 400\nRequest Failed")

func main() {
	var results = map[string]string{}
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
		result := "Ok"
		err := hitURL(url)
		if err != nil {
			result = "failed"
		}
		results[url] = result return errRequestFailed
	}
	for url, result := range results{
		fmt.Println("checking:", url)
	}
	return nil
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	if resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
