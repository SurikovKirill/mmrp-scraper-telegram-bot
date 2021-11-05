package scrapers

import (
	"crypto/md5"
	"fmt"
	"mmrp-scraper/internal/telegram"
)

// MAPMScraper ...
type MAPMScraper struct {
	lastArrivalCheckSum string
}

func (s *MAPMScraper) Scrape() {
	doc := GetDocument("http://mapm.ru/Port/Murmansk")
	fmt.Println("searching new report")
	text := doc.Find(".text-success").Text()
	fmt.Println(text)
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		link, _ := doc.Find(".text-success").Attr("href")
		fmt.Println(link)
		telegram.SendDocument(fmt.Sprintf("http://mapm.ru%s", link))
	}
}
