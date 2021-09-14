package main

import (
	"fmt"
	"log"
	"net/http"
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
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage (page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	
	res, err := http.Get(pageURL)
	checkError(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		extractJob(card)
	})
}

func extractJob(card *goquery.Selection) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	fmt.Println(id, title, location)
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