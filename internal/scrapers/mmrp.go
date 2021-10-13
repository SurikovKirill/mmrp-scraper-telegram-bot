package scrapers

import (
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"mmrp-scraper/internal/telegram"
	"net/http"
	"strings"
)

// MMRPScraper ...
type MMRPScraper struct {
	lastArrivalCheckSum string
}

func (s *MMRPScraper) Scrape() {
	//Initialize document
	res, err := http.Get("http://mmrp.ru/news/74/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Searching the last report
	a, _ := doc.Find(".row").Find(".t_1").First().Find("a").Attr("href")
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(a)))

	// Extracting data if report is new, make new checksum
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		link := fmt.Sprintf("http://mmrp.ru%s", a)
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		}
		data, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		//Getting date of report
		dateReport := strings.ReplaceAll(strings.ReplaceAll(data.Find(".date-news").Text(), "\t", ""), "\n", "")

		//Extracting headings
		var headings []string
		data.Find(".container.content").Find(".col-lg-12").Has("b").Find("b").Each(func(i int, s *goquery.Selection) {
			headings = append(headings, s.Text())
		})

		//Extracting metadata
		metadata := strings.ReplaceAll(data.Find(".container.content").Find(".col-lg-12").Has("b").Text(), "\t", "")
		for i := range headings {
			metadata = strings.ReplaceAll(metadata, headings[i], "")
		}
		metadata = strings.ReplaceAll(metadata, dateReport, "")
		metadata = strings.TrimSpace(metadata)
		metadata = strings.ReplaceAll(metadata, " \n ", "")
		spl := strings.Split(metadata, "\n")
		for i := range spl {
			spl[i] = strings.TrimSpace(spl[i])
		}
		fmt.Println("Nofify")
		telegram.SendMessage()

		//notify user

	}
}
