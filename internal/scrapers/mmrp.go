package scrapers

import (
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"mmrp-scraper/internal/telegram"
	"strings"
	"sync"
)

// MMRPScraper ...
type MMRPScraper struct {
	lastArrivalCheckSum string
}

func (s *MMRPScraper) Scrape(cfg Config) {
	//Initialize checksum from file
	var once sync.Once
	once.Do(func() {
		s.lastArrivalCheckSum = ReadCheckSum("./.checksum_mmrp")
		log.Println("Checksum for MMRP has been read")
	})

	//Initialize main document
	doc := GetDocument(fmt.Sprintf("%s/news/74/", cfg.MmrpUrl))

	// Searching the last report
	a, _ := doc.Find(".row").Find(".t_1").First().Find("a").Attr("href")
	fmt.Println(a)
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(a)))

	// Extracting data if report is new, make new checksum
	if s.lastArrivalCheckSum != checksum {
		log.Println("New report in MMRP")
		s.lastArrivalCheckSum = checksum
		SaveCheckSum(checksum, "./.checksum_mmrp")
		reportDoc := GetDocument(fmt.Sprintf("%s%s", cfg.MmrpUrl, a))

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
