package scrapers

import (
	"crypto/md5"
	"fmt"
	"log"
	"mmrp-scraper/internal/telegram"
)

// MAPMScraper ...
type MAPMScraper struct {
	lastArrivalCheckSum string
}

func (s *MAPMScraper) Scrape() {
	cfg, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	doc := GetDocument(fmt.Sprintf("%s/Port/Murmansk", cfg.MapmUrl))
	fmt.Println("searching new report")
	text := doc.Find(".text-success").Text()
	fmt.Println(text)
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		link, _ := doc.Find(".text-success").Attr("href")
		fmt.Println(link)
		telegram.SendDocument(fmt.Sprintf("%s%s", cfg.MapmUrl, link))
	}
}
