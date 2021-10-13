package main

import (
	"log"
	"mmrp-scraper/internal/scrapers"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	s := scrapers.MMRPScraper{}
	s.Scrape()
	return nil
}
