package scrapers

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// GetDocument ...
func GetDocument(url string) *goquery.Document {
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

func GetInfoFromFile(filename string) string {
	if CheckFileExist(filename) {
		if CheckFileIsEmpty(filename) {
			return ""
		} else {
			return ReadCheckSumFromFile(filename)
		}
	} else {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		return ""
	}
}

func CheckFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func CheckFileIsEmpty(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	if info.Size() == 0 {
		return true
	} else {
		return false
	}
}

func ReadCheckSumFromFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 200)
	n, _ := file.Read(data)
	return string(data[:n])
}

func ChangeCheckSumFile(filename string, checksum string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := os.Truncate(filename, 0); err != nil {
		log.Fatal(err)
	}
	if _, err := file.WriteString(checksum); err != nil {
		log.Fatal(err)
	}

}
