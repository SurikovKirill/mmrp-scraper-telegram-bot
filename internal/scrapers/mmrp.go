package scrapers

import (
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"mmrp-scraper/internal/telegram"
	"strings"
)

// MMRPScraper ...
type MMRPScraper struct {
	lastArrivalCheckSum string
}


func GetDocument(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func (s *MMRPScraper) Scrape() {
	//Initialize main document
	doc := GetDocument("http://mmrp.ru/news/74/")

	// Searching the last report
	fmt.Println("searching new report")
	a, _ := doc.Find(".row").Find(".t_1").First().Find("a").Attr("href")
	fmt.Println(a)
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(a)))

	// Extracting data if report is new, make new checksum
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		fmt.Println("New report!!!")
		reportDoc := GetDocument(fmt.Sprintf("http://mmrp.ru%s", a))

		//Getting date of report
		dateReport := strings.ReplaceAll(strings.ReplaceAll(reportDoc.Find(".date-news").Text(), "\t", ""), "\n", "")

		//Extracting headings
		var headers []string
		reportDoc.Find(".container.content").Find(".col-lg-12").Has("b").Find("b").Each(func(i int, s *goquery.Selection) {
			headers = append(headers, strings.TrimSpace(s.Text()))
		})

		//Extracting metadata
		data := reportDoc.Find(".container.content").Find(".col-lg-12").Has("b").Text()

		//Cleaning off duplicate data
		for i := range headers {
			data = strings.ReplaceAll(data, headers[i], "")
		}
		data = strings.ReplaceAll(data, dateReport, "")

		//Cleaning off spaces and split data
		spl := strings.Split(strings.ReplaceAll(strings.TrimSpace(data), " \n ", ""), "\n")
		for i := range spl {
			spl[i] = strings.TrimSpace(spl[i])
			spl[i] = strings.ReplaceAll(spl[i], "\"", "\\\"")
		}

		//Mapping data
		m := make(map[string]string)
		for i, val := range headers {
			m[val] = spl[i]
		}

		telegram.SendMessage(dateReport, m)

	}
}
