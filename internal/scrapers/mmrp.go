package scrapers

import (
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

// MMRPScraper ...
type MMRPScraper struct {
	//config *Config
	lastArrivalCheckSum string
}

func (s *MMRPScraper) Scrape() {
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

	a, _ := doc.Find(".row").Find(".t_1").First().Find("a").Attr("href")
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(a)))
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		link := fmt.Sprintf("http://mmrp.ru%s", a)
		fmt.Println(link)
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
		metadata1 := data.Find(".container.content").Find(".container.full-width").Find(".col-lg-12").Has(".container.full-width").Find(".date-news").Text()

		arriving := strings.ReplaceAll(metadata1, "\t", "")
		metadata2 := data.Find(".container.content").Find(".col-lg-12").Has("br").Text()
		dataset := strings.ReplaceAll(metadata2, "\t", "")
		fmt.Println(arriving)
		fmt.Println(dataset)
		//notify user
	}
}
