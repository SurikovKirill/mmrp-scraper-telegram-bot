package scrapers

import (
	"log"
	"mmrp-scraper/internal/telegram"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/spf13/viper"
	"github.com/ysmood/gson"
)

// MAPMScraper Structure for MAPM
type MAPMScraper struct {
	login    string
	password string
}

// Init initialize login and password for MAPM from environment
func (s *MAPMScraper) Init() error {
	os.Setenv("LOGIN", "NaLogMo")
	os.Setenv("PASSWORD", "dfm2jslp")
	if err := viper.BindEnv("login"); err != nil {
		return err
	}
	if err := viper.BindEnv("password"); err != nil {
		return err
	}
	s.login = viper.GetString("login")
	s.password = viper.GetString("password")
	log.Println("MAPM config done")
	return nil
}

// ScrapeWithRod Scraping MAPM using rod-driver
func (s *MAPMScraper) ScrapeWithRod(t *telegram.Config) {
	// Подключение к движку
	log.Println("Connecting to chromium ...")
	path, h := launcher.LookPath()
	if !h {
		log.Fatal("Can't find path to chromium", h, path)
	}
	u := launcher.New().Bin(path).MustLaunch()
	br := rod.New().ControlURL(u).MustConnect()
	// Авторизация на сайте
	log.Println("Working with MAPM ...")
	lp := br.MustPage("http://mapm.ru/Account/Login?returnUrl=%2F")
	time.Sleep(time.Millisecond * 5000)
	log.Println(s.login, s.password)
	lp.MustElement("#UserName").MustInput(s.login)
	lp.MustElement("#Password").MustInput(s.password)
	lp.MustElement("#loginForm > form > div:nth-child(7) > div > input").MustClick()
	log.Println("Working with MAPM ... 1")
	br.MustPage("http://mapm.ru/")
	// Переход по ссылке на таблицу с данными, формирование запроса
	tp := br.MustPage("http://mapm.ru/Vts")
	log.Println("Working with MAPM ... 2")
	time.Sleep(time.Millisecond * 5000)
	log.Println("Working with MAPM ... 3")
	tp.MustElement("#ddlVtsPort").MustSelect("Мурманск")
	tp.MustElement("#wrapper > div:nth-child(4) > div > div:nth-child(3) > div > div > button").MustClick()
	log.Println("Working with MAPM ... 4")
	time.Sleep(time.Millisecond * 10000)
	// Дожидаемся полной загрузки страницы и переносим данные в html файл
	log.Println("Working with MAPM ... 5")
	tp.MustElement("#dvShipsResults > div.center-block.table-responsive")
	log.Println("Making PDF file ...")
	pdf, err := tp.PDF(&proto.PagePrintToPDF{
		PaperWidth:  gson.Num(25),
		PaperHeight: gson.Num(11),
		PageRanges:  "1-10",
	})
	if err != nil {
		log.Println(err)
	}
	err = utils.OutputFile("temp.pdf", pdf)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(time.Millisecond * 3000)
	// Отправляем данные
	telegram.SendDocumentRod(t)
}
