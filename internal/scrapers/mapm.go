package scrapers

import (
	"crypto/md5"
	"fmt"
	"log"
	"mmrp-scraper/internal/telegram"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

// MAPMScraper Structure for MAPM
type MAPMScraper struct {
	lastArrivalCheckSum string
}

// Scrape Scraping MAPM
// TEMPORARILY DEPRECATED
// TODO: отрефакторить модуль
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

// TODO: вынести в переменные среды логопас
const (
	login    = "NaLogMo"
	password = "dfm2jslp"
)

// ScrapeWithRod Scraping MAPM using rod-driver
func (s *MAPMScraper) ScrapeWithRod() {
	// Подключение к движку
	path, h := launcher.LookPath()
	if !h {
		log.Println(h, path)
	}
	u := launcher.New().Bin(path).MustLaunch()
	br := rod.New().ControlURL(u).MustConnect()
	// Авторизация на сайте
	lp := br.MustPage("http://mapm.ru/Account/Login?returnUrl=%2F")
	time.Sleep(time.Millisecond * 5000)
	lp.MustElement("#UserName").MustInput(login)
	lp.MustElement("#Password").MustInput(password)
	lp.MustElement("#loginForm > form > div:nth-child(7) > div > input").MustClick()
	br.MustPage("http://mapm.ru/")
	// Переход по ссылке на таблицу с данными, формирование запроса
	tp := br.MustPage("http://mapm.ru/Vts")
	time.Sleep(time.Millisecond * 5000)
	tp.MustElement("#ddlVtsPort").MustSelect("Мурманск")
	tp.MustElement("#wrapper > div:nth-child(4) > div > div:nth-child(3) > div > div > button").MustClick()
	time.Sleep(time.Millisecond * 5000)
	// Дожидаемся полной загрузки страницы и переносим данные в html файл
	data := tp.MustElement("#dvShipsResults > div.center-block.table-responsive").MustHTML()
	pdf, _ := tp.PDF(&proto.PagePrintToPDF{
		PaperWidth:  gson.Num(25),
		PaperHeight: gson.Num(11),
		PageRanges:  "1-10",
	})
	_ = utils.OutputFile("temp.pdf", pdf)
	f, err := os.Create("tmp.html")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString(data)
	time.Sleep(time.Millisecond * 3000)
	//Отправляем данные
	telegram.SendDocumentRod()
}
