package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Puerkitobio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}


//Scrape indeed by a term
func Scrape(term) {
	var baseURL string = "https://www.indeed.com/jobs?q=" + term + "limit=50&start="
	var jobs []extractedJob
	totalPages := getPages()
	c := make(chan []extractedJob)

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)

	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs, c)
	fmt.Println("Done, extraced", len(jobs))
}

// getPages gets the number of pages webpages at the bottom of indeed.com and returns a number so we know how many pages to scrape the jobs from
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

// getPage pulls the jobcards from extractJob and puts them through each page.
func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	//pageURL is the main indeed.com and required info incrimented by the number of pages.
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)

	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

// extractJob pulls the data by page class, using goquery for scraping.
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	// classes from the page, white space cleaned up, listed below and stored in variables
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	//stored into struct and sent to the channel so getPage can store it into a slice
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

//takes input and stores it into a csv file called jobs
func writeJobs(jobs []extractedJob, writtenJobs) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		go jobSlice := []string{"https://indeed.com/viewjob?k=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
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

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
