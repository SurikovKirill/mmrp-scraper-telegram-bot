package scrapers

import (
	"crypto/md5"
	"fmt"
	"log"
	"mmrp-scraper/internal/telegram"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// MMRPScraper ...
type MMRPScraper struct {
	lastArrivalCheckSum string
}

func (s *MMRPScraper) Scrape(cfg *Config) {
	// Initialize main document
	doc := GetDocument(fmt.Sprintf("%s/news/74/", cfg.MmrpUrl))
	// Searching the last report
	link, ex := doc.Find(".row").Find(".t_1").First().Find("a").Attr("href")
	if !ex {
		log.Fatal("Link for MMRP doesn't exist")
	}
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(link)))
	if s.lastArrivalCheckSum == "" {
		s.lastArrivalCheckSum = GetInfoFromFile("lastInfoMmrp")
	}
	// Extracting data, if report is new, make new checksum
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		ChangeCheckSumFile("lastInfoMmrp", checksum)
		reportDoc := GetDocument(fmt.Sprintf("%s%s", cfg.MmrpUrl, link))
		//Getting date of report
		dateReport := strings.ReplaceAll(strings.ReplaceAll(reportDoc.Find(".date-news").Text(), "\t", ""), "\n", "")
		//Extracting headers
		var headers []string
		reportDoc.Find(".container.content").Find(".col-lg-12").Has("b").Find("b").Each(func(i int, s *goquery.Selection) {
			headers = append(headers, strings.ReplaceAll(strings.TrimSpace(s.Text()), ":", ""))
		})
		//Extracting metadata
		data := reportDoc.Find(".container.content").Find(".col-lg-12").Has("b").Text()
		//Cleaning data
		for i := range headers {
			data = strings.ReplaceAll(data, headers[i], "")
		}
		data = strings.ReplaceAll(data, dateReport, "")
		spl := strings.Split(strings.ReplaceAll(strings.TrimSpace(data), " \n ", ""), ":")
		for i := range spl {
			spl[i] = strings.TrimSpace(spl[i])
			spl[i] = strings.ReplaceAll(spl[i], "\"", "\\\"")
		}
		//Mapping data
		m := make(map[string]string)
		for i, val := range headers {
			m[val] = spl[i+1]
		}
		telegram.SendMessage(dateReport, m)
	}
}
