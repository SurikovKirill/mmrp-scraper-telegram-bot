package scrapers

type Scraper interface {
	Scrape(cfg Config)
}
