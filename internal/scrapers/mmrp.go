package scrapers

import (
	"github.com/gocolly/colly"
	"time"
)

// MMRPScraper ...
type MMRPScraper struct {
}

func (s *MMRPScraper) Start() {
	c := colly.NewCollector(colly.AllowedDomains(), colly.AllowURLRevisit(), colly.Async(true))
	c.Limit(&colly.LimitRule{
		Delay: 1 * time.Second,
	})
	c.Visit("")
	c.OnHTML("", func(e *colly.HTMLElement) {

	})
}
