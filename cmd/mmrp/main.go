package main

import (
	"github.com/go-co-op/gocron"
	"time"

	//"github.com/go-co-op/gocron"
	//"log"
	"log"
	"mmrp-scraper/internal/scrapers"
	//"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	s1 := scrapers.MMRPScraper{}
	s2 := scrapers.MAPMScraper{}
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(15).Minutes().Do(func() {
		s1.Scrape()
		s2.Scrape()
	})
	if err != nil {
		log.Fatal(err)
	}
	scheduler.StartBlocking()
	return nil
}
