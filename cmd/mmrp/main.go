package main

import (
	"github.com/go-co-op/gocron"
	"log"
	"mmrp-scraper/internal/scrapers"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	sMmrp := scrapers.MMRPScraper{}
	sMapm := scrapers.MAPMScraper{}
	cs, err := scrapers.Init()

	if err != nil {
		log.Fatal("Error in config", err)
	}
	scheduler := gocron.NewScheduler(time.UTC)
	_, err = scheduler.Every(15).Minutes().Do(func() {
		sMmrp.Scrape(cs)
		sMapm.Scrape(cs)
	})
	if err != nil {
		log.Fatal(err)
	}
	scheduler.StartBlocking()
	return nil
}
