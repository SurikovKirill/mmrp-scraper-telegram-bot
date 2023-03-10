package scrapers

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

// MAPMScraper Structure for MAPM
type MAPMScraper struct {
	login    string
	password string
}

// Init initialize login and password for MAPM from environment
func (s *MAPMScraper) Init() error {
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

const (
	loginPage = "http://mapm.ru/Account/Login?returnUrl=%2F"
	mainPage  = "http://mapm.ru/"
	tablePage = "http://mapm.ru/Vts"
	filename  = "temp.pdf"
)

// ScrapeWithRod Scraping MAPM using rod-driver
func (s *MAPMScraper) ScrapeWithRod(t *telegram.Config) {
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
	log.Println("Working with MAPM ...")
	if err := rod.Try(func() { br.MustPage(loginPage) }); err != nil {
		log.Println("Problems with connection to mapm", err)
		return
	}
	lp := br.MustPage(loginPage)
	time.Sleep(time.Millisecond * 5000)
	lp.MustElement("#UserName").MustInput(s.login)
	lp.MustElement("#Password").MustInput(s.password)
	lp.MustElement("#loginForm > form > div:nth-child(7) > div > input").MustClick()
	br.MustPage(mainPage)
	tp := br.MustPage(tablePage)
	time.Sleep(time.Millisecond * 5000)
	tp.MustElement("#ddlVtsPort").MustSelect("Мурманск")
	tp.MustElement("#wrapper > div:nth-child(4) > div > div:nth-child(3) > div > div > button").MustClick()
	time.Sleep(time.Millisecond * 5000)
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
	if err := utils.OutputFile(filename, pdf); err != nil {
		log.Println(err)
	}
	telegram.SendDocumentRod(t)
}
