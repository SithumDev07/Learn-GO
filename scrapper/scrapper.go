package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	location string
	title string
	salary string
	description string
}

// * Scrape indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	
	var jobs []extractedJob
	c := make(chan []extractedJob)

	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i :=0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done", len(jobs))
}

func getPage (page int, url string, mainChannel chan<- []extractedJob) {

	var jobs []extractedJob

	c:= make(chan extractedJob)

	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("requesting", pageURL)

	res, err := http.Get(pageURL)
	checkError(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
 		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainChannel <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".jobTitle>span").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	salary := CleanString(card.Find(".salary-snippet").Text())
	description := CleanString(card.Find(".job-snippet").Text())
	c <- extractedJob{
		id: id,
		title: title,
		location: location,
		salary: salary,
		description: description,
	}
}

func getPages(url string) int {

	pages := 0

	res, err := http.Get(url)
	checkError(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	
	return pages
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkError(err)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"ID", "title", "location", "salary", "description",
	}

	writeError := writer.Write(headers)
	checkError(writeError)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.description}
		jobWriteError := writer.Write(jobSlice)
		checkError(jobWriteError)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}	
} 

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status code: ", res.StatusCode)
	}
}

func CleanString (str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}