package main

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

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {

	var jobs []extractedJob

	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done", len(jobs))
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

func getPage (page int) []extractedJob{

	var jobs []extractedJob

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("requesting", pageURL)

	res, err := http.Get(pageURL)
	checkError(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
 		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(card *goquery.Selection) extractedJob{
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := cleanString(card.Find(".salary-snippet").Text())
	description := cleanString(card.Find(".job-snippet").Text())
	return extractedJob{
		id: id,
		title: title,
		location: location,
		salary: salary,
		description: description,
	}
}

func getPages() int {

	pages := 0

	res, err := http.Get(baseURL)
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

func cleanString (str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}