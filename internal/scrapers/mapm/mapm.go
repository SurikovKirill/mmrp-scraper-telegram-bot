package mapm

import (
	"log"
	"mmrp-scraper/internal/telegram"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/spf13/viper"
	"github.com/ysmood/gson"
)

// Scraper Structure for MAPM
type Scraper struct {
	login    string
	password string
}

// Init initialize login and password for MAPM from environment
func (s *Scraper) Init() error {
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
func (s *Scraper) ScrapeWithRod(t *telegram.Config) {
	// Подключение к движку
	log.Println("Connecting to chromium ...")
	path, h := launcher.LookPath()
	if !h {
		log.Fatal("Can't find path to chromium", h, path)
	}
	u := launcher.New().Bin(path).MustLaunch()
	br := rod.New().ControlURL(u).MustConnect()
	defer func() {
		if err := br.Close(); err != nil {
			log.Println(err)
		}
	}()
	// Авторизация на сайте
	log.Println("Working with MAPM ...")
	if err := rod.Try(func() { br.MustPage("https://mapm.ru/Account/Login?returnUrl=%2F") }); err != nil {
		log.Println("Problems with connection to mapm", err)
		return
	}
	log.Println("Connected")
	br.MustPage("https://mapm.ru/Account/Login?returnUrl=%2F").MustWaitStable().MustScreenshot("a.png")
	lp := br.MustPage("https://mapm.ru/Account/Login?returnUrl=%2F")
	time.Sleep(time.Millisecond * 5000)
	lp.MustElement("#UserName").MustInput("NaLogMo")
	lp.MustElement("#Password").MustInput("dfm2jslp")
	lp.MustScreenshot("c.png")
	lp.MustElement("#loginForm > form > div:nth-child(7) > div > input").MustClick()
	time.Sleep(time.Millisecond * 20000)
	br.MustPage("https://mapm.ru/").MustWaitStable().MustScreenshot("b.png")
	log.Println("dsdadasdas")
	// Переход по ссылке на таблицу с данными, формирование запроса
	tp := br.MustPage("https://mapm.ru/Vts")
	tp.MustScreenshot("e.png")
	time.Sleep(time.Millisecond * 5000)
	tp.MustElement("#ddlVtsPort").MustSelect("Мурманск")
	tp.MustWaitStable().MustScreenshot("abc.png")
	tp.MustElement("#wrapper > div:nth-child(4) > div > div:nth-child(3) > div > div > button").MustClick()
	time.Sleep(time.Millisecond * 20000)
	// Дожидаемся полной загрузки страницы и переносим данные в html файл
	tp.MustWaitStable().MustScreenshot("abcd.png")
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
	if err := utils.OutputFile("temp.pdf", pdf); err != nil {
		log.Println(err)
	}
	time.Sleep(time.Millisecond * 3000)
	// Отправляем данные
	telegram.SendDocumentRod(t)
}
