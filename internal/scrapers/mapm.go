package scrapers

import (
	"crypto/md5"
	"fmt"
	"log"
	"mmrp-scraper/internal/telegram"
	"sync"
)

// MAPMScraper ...
type MAPMScraper struct {
	lastArrivalCheckSum string
}

func (s *MAPMScraper) Scrape(cfg Config) {
	var once sync.Once
	once.Do(func() {
		s.lastArrivalCheckSum = ReadCheckSum("./.checksum_mapm")
		log.Print("Checksum for MAPM has been read")
	})
	doc := GetDocument(fmt.Sprintf("%s/Port/Murmansk", cfg.MapmUrl))
	text := doc.Find(".text-success").Text()
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	if s.lastArrivalCheckSum != checksum {
		log.Println("New report in MAPM")
		s.lastArrivalCheckSum = checksum
		SaveCheckSum(checksum, "./.checksum_mapm")
		link, _ := doc.Find(".text-success").Attr("href")
		telegram.SendDocument(fmt.Sprintf("%s%s", cfg.MapmUrl, link))
	}
}
