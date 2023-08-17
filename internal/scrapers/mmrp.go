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

const (
	lastInfoMmrp = "lastInfoMmrp"
	urlMmrp      = "https://mmrp.ru"
)

// Scrape ...
func (s *MMRPScraper) Scrape(t *telegram.Config) {
	// Initialize main document
	doc, err := GetDocument(fmt.Sprintf("%s/news/74/", urlMmrp))
	if err != nil {
		log.Println(err)
		return
	}
	// Searching the last report
	link, ex := doc.Find(".row").Find(".t_1").First().Find("a").Attr("href")
	if !ex {
		log.Println("Link for MMRP doesn't exist")
		return
	}
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(link)))
	if s.lastArrivalCheckSum == "" {
		s.lastArrivalCheckSum = GetInfoFromFile(lastInfoMmrp)
	}
	// Extracting data, if report is new, make new checksum
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		ChangeCheckSumFile(lastInfoMmrp, checksum)
		reportDoc, err := GetDocument(fmt.Sprintf("%s%s", urlMmrp, link))
		if err != nil {
			log.Println(err)
		}
		// Getting date of report
		dateReport := strings.ReplaceAll(strings.ReplaceAll(reportDoc.Find(".date-news").Text(), "\t", ""), "\n", "")
		// Extracting headers
		var headers []string
		reportDoc.Find(".container.content").Find(".col-lg-12").Has("b").Find("b").Each(func(i int, s *goquery.Selection) {
			headers = append(headers, strings.ReplaceAll(strings.TrimSpace(s.Text()), ":", ""))
		})
		headers = checkEmptyElements(headers)
		// Extracting metadata
		data := reportDoc.Find(".container.content").Find(".col-lg-12").Has("b").Text()
		// Cleaning data
		for i := range headers {
			data = strings.ReplaceAll(data, headers[i], "")
		}
		data = strings.ReplaceAll(data, dateReport, "")
		spl := strings.Split(strings.ReplaceAll(strings.TrimSpace(data), " \n ", ""), ":")
		for i := range spl {
			spl[i] = strings.TrimSpace(spl[i])
			spl[i] = strings.ReplaceAll(spl[i], "\"", "\\\"")
		}
		// Mapping data
		m := make(map[string]string)
		for i, val := range headers {
			m[val] = spl[i+1]
		}
		telegram.SendMessage(dateReport, m, t)
	}
}

func checkEmptyElements(el []string) []string {
	var buf []string
	for _, val := range el {
		if val != "" {
			buf = append(buf, val)
		}
	}
	return buf
}
