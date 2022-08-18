package scrapers

import (
	"crypto/md5"
	"fmt"
	"log"
	"mmrp-scraper/internal/telegram"
	"os"
	"time"

	"github.com/go-rod/rod"
)

// MAPMScraper Structure for MAPM
type MAPMScraper struct {
	lastArrivalCheckSum string
}

// Scrape Scraping MAPM
// TEMPORARILY DEPRECATED
func (s *MAPMScraper) Scrape(cfg *Config) {
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
			log.Println("Link for MAPM doesn't exist")
			return
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
// TODO: make config with logopass for mapm
// TODO: make helper function for html-info from mapm
func (s *MAPMScraper) ScrapeWithRod() {
	br := rod.New().MustConnect()

	lp := br.MustPage("http://mapm.ru/Account/Login?returnUrl=%2F")
	lp.MustElement("#UserName").MustInput(login)
	lp.MustElement("#Password").MustInput(password)
	lp.MustElement("#loginForm > form > div:nth-child(7) > div > input").MustClick()

	tp := br.MustPage("http://mapm.ru/Vts")
	tp.MustElement("#ddlVtsPort").MustSelect("Мурманск")
	tp.MustElement("#wrapper > div:nth-child(4) > div > div:nth-child(3) > div > div > button").MustClick()
	time.Sleep(time.Millisecond * 1200)
	data := tp.MustElement("#dvShipsResults > div.center-block.table-responsive").MustHTML()

	f, err := os.Create("tmp.html")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString(data)
	time.Sleep(time.Millisecond * 2000)
	telegram.SendDocumentRod()
}
