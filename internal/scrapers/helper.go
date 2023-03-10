package scrapers

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// GetDocument gets goquery-document from URL
func GetDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	return doc, nil
}

// GetInfoFromFile gets saved hash-checksum from file
func GetInfoFromFile(filename string) string {
	if checkFileExist(filename) {
		if checkFileIsEmpty(filename) {
			return ""
		}
		return readCheckSumFromFile(filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error in opening file: ", err)
	}
	file.Close()
	return ""
}

func checkFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func checkFileIsEmpty(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
	}
	if info.Size() == 0 {
		return true
	}
	return false
}

func readCheckSumFromFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	data := make([]byte, 200)
	n, _ := file.Read(data)
	return string(data[:n])
}

// ChangeCheckSumFile change or save hash-checksum in file
func ChangeCheckSumFile(filename string, checksum string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if err := os.Truncate(filename, 0); err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString(checksum); err != nil {
		log.Println(err)
	}
}
