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
		//Getting date
		arriving := strings.ReplaceAll(strings.ReplaceAll(data.Find(".date-news").Text(), "\t", ""), "\n", "")

		//Getting headings
		var headings []string
		data.Find(".container.content").Find(".col-lg-12").Has("b").Find("b").Each(func(i int, s *goquery.Selection) {
			headings = append(headings, s.Text())
		})

		metadata := strings.ReplaceAll(data.Find(".container.content").Find(".col-lg-12").Has("b").Text(), "\t", "")
		for i := range headings {
			metadata = strings.ReplaceAll(metadata, headings[i], "")
		}
		metadata = strings.ReplaceAll(metadata, arriving, "")
		metadata = strings.TrimSpace(metadata)
		metadata = strings.ReplaceAll(metadata, " \n ", "")
		spl := strings.Split(metadata, "\n")

		for i := range spl {
			spl[i] = strings.TrimSpace(spl[i])
		}
		fmt.Println(spl)
		//for i, _ := range spl{
		//	spl[i] = strings.ReplaceAll(spl[i], "\n", "")
		//}
		//metadata = strings.ReplaceAll(metadata, "\t", "")
		//metadata = strings.ReplaceAll(metadata, "\n", "")
		//spl := strings.Split(metadata, "  ")
		//fmt.Println(spl)
		//metadata = strings.ReplaceAll(metadata, "\t", "")
		//metadata = strings.ReplaceAll(metadata, "\n", "")
		//metad, _ := data.Find(".container.content").Find(".col-lg-12").Has("b").NotSelection(metadata).Html()

		fmt.Println(arriving)
		fmt.Println(headings)
		fmt.Println(metadata)
		//notify user
	}
}
