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
	s1 := scrapers.MMRPScraper{}
	s2 := scrapers.MAPMScraper{}
	cfg, err := scrapers.Init()
	if err != nil {
		log.Fatal(err)
	}
	scheduler := gocron.NewScheduler(time.UTC)
	_, err = scheduler.Every(15).Minutes().Do(func() {
		s1.Scrape(*cfg)
		s2.Scrape(*cfg)
	})
	if err != nil {
		log.Fatal(err)
	}
	scheduler.StartBlocking()
	return nil
}
