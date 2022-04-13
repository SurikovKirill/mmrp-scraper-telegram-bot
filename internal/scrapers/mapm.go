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

func (s *MAPMScraper) Scrape(cfg *Config) {
	doc := GetDocument(fmt.Sprintf("%s/Port/Murmansk", cfg.MapmUrl))
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(doc.Find(".text-success").Text())))
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		link, ex := doc.Find(".text-success").Attr("href")
		if ex == false {
			log.Fatal("Link for MAPM doesn't exist")
		}
		telegram.SendDocument(fmt.Sprintf("%s%s", cfg.MapmUrl, link))
	}
}
