package scrapers

import (
	"crypto/md5"
	"fmt"
	"mmrp-scraper/internal/telegram"
	"os"

	"github.com/go-rod/rod"
	log "github.com/sirupsen/logrus"
)

// MAPMScraper Structure for MAPM
type MAPMScraper struct {
	lastArrivalCheckSum string
}

// Scrape Scraping MAPM
// TEMPORARILY DEPRECATED
func (s *MAPMScraper) Scrape(cfg *Config) {
	// Подготовка лог файла
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	lf := log.Fields{"package": "scrapers", "struct": "MAPMScraper", "function": "Scrape"}
	if err != nil {
		log.WithFields(lf).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()

	// Получение документа
	doc := GetDocument(fmt.Sprintf("%s/Port/Murmansk", cfg.MapmUrl))
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(doc.Find(".text-success").Text())))
	if s.lastArrivalCheckSum == "" {
		s.lastArrivalCheckSum = GetInfoFromFile("lastInfoMapm")
	}
	if s.lastArrivalCheckSum != checksum {
		s.lastArrivalCheckSum = checksum
		ChangeCheckSumFile("lastInfoMapm", checksum)
		link, ex := doc.Find(".text-success").Attr("href")
		if !ex {
			log.WithFields(lf).Error("Link for MAPM doesn't exist")
		}
		log.Println("Send file")
		telegram.SendDocument(fmt.Sprintf("%s%s", cfg.MapmUrl, link))
	}
}

const (
	login    = "NaLogMo"
	password = "dfm2jslp"
)

// ScrapeWithRod Scraping MAPM using rod-driver
func (s *MAPMScraper) ScrapeWithRod() {
	br := rod.New().MustConnect()

	lp := br.MustPage("http://mapm.ru/Account/Login?returnUrl=%2F")
	lp.MustElement("#UserName").MustInput(login)
	lp.MustElement("#Password").MustInput(password)
	lp.MustElement("#loginForm > form > div:nth-child(7) > div > input").MustClick()

	tp := br.MustPage("http://mapm.ru/Vts")
	tp.MustElement("#ddlVtsPort").MustSelect("Мурманск")
	tp.MustElement("#wrapper > div:nth-child(4) > div > div:nth-child(3) > div > div > button").MustClick()

}
