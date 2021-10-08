package scrapers

import (
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
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
		metadata := data.Find(".col-lg-12").Find("b").Text()

		fmt.Println(metadata)
		//notify user
	}
}
