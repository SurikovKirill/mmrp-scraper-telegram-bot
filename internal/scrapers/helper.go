package scrapers

import (
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

// GetDocument gets goquery-document from URL
func GetDocument(url string) *goquery.Document {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocument", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	res, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocument", "error": err}).Error("Bad request")
	}
	if res.StatusCode != 200 {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocument", "error": err}).Error("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocument", "error": err}).Error(err)
	}
	return doc
}

// GetInfoFromFile gets saved hash-checksum from file
func GetInfoFromFile(filename string) string {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetInfoFromFile", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	if CheckFileExist(filename) {
		if CheckFileIsEmpty(filename) {
			return ""
		} else {
			return ReadCheckSumFromFile(filename)
		}
	} else {
		file, err := os.Create(filename)
		if err != nil {
			log.WithFields(log.Fields{"package": "scrapers", "function": "GetInfoFromFile", "error": err}).Error("Error in opening file: %v", err)
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
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "CheckFileIsEmpty", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	info, err := os.Stat(filename)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "CheckFileIsEmpty", "error": err}).Error(err)
	}
	if info.Size() == 0 {
		return true
	} else {
		return false
	}
}

func ReadCheckSumFromFile(filename string) string {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "ReadCheckSumFromFile", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	file, err := os.Open(filename)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "ReadCheckSumFromFile", "error": err}).Error(err)
	}
	data := make([]byte, 200)
	n, _ := file.Read(data)
	return string(data[:n])
}

// ChangeCheckSumFile change or save hash-checksum in file
func ChangeCheckSumFile(filename string, checksum string) {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "ChangeCheckSumFile", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "ChangeCheckSumFile", "error": err}).Error(err)
	}
	defer file.Close()
	if err := os.Truncate(filename, 0); err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "ChangeCheckSumFile", "error": err}).Error(err)
	}
	if _, err := file.WriteString(checksum); err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "ChangeCheckSumFile", "error": err}).Error(err)
	}
}

func GetDocumentWithCookie(url string) (*goquery.Document, string) {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocumentWithCookie", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	res, err := http.Get(url)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocumentWithCookie", "error": err}).Error("Bad request")
	}
	if res.StatusCode != 200 {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocumentWithCookie", "error": err}).Error("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "GetDocumentWithCookie", "error": err}).Error(err)
	}
	return doc, res.Cookies()[0].Value
}

// func CustomGetWithCookie(url string) (*goquery.Document, string) {

// }
