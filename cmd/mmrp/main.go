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
	s := scrapers.MMRPScraper{}
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(15).Minutes().Do(func() {
		s.Scrape()
	})
	if err != nil {
		log.Fatal(err)
	}
	scheduler.StartBlocking()
	return nil
}
