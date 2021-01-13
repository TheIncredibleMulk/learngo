package main

import (
	"log"
	"net/http"

	"github.com/Puerkitobio/goquery"
)

var baseURL string = "https://www.indeed.com/jobs?q=python&start=20"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination")

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Responce Failed with Status:", res.StatusCode)
	}
}
