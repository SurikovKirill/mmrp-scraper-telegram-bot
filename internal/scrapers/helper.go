package scrapers

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetDocument ...
func GetDocument(url string) *goquery.Document {
	log.Println("Getting document")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func SaveCheckSum(chs string, filename string) {
	log.Println("Saving checksum")
	err := ioutil.WriteFile(filename, []byte(chs), 777)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadCheckSum(filename string) string {
	log.Println("Reading checksum")
	if _, err := os.Stat(filename); err == nil {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		return string(data)
	} else {
		return ""
	}
}
