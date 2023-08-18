package mmrp

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"mmrp-scraper/internal/telegram"
	"strings"
)

// Scraper ...
type Scraper struct {
	lastArrivalCheckSum string
}

const (
	lastInfo = "lastInfo"
	url      = "https://mmrp.ru/"
)

// Scrape mmrp reports
func (s *Scraper) Scrape(t *telegram.Config) {
	// Initialize main document
	doc, err := GetDocument(url)
	if err != nil {
		log.Println(err)
		return
	}

	// Get last report
	text := doc.Find(".news-list").First().Find(".news-item").Text()
	list := strings.Split(text, "\n")
	list = deleteEmptyElements(list)
	for i := range list {
		list[i] = strings.TrimSpace(list[i])
	}

	// TODO: добавить k/v хранилище для хранения чек-суммы
	// Check last report is new
	checkSum := createCheckSum(list)
	if s.lastArrivalCheckSum == "" {
		s.lastArrivalCheckSum = GetInfoFromFile(lastInfo)
	}
	if s.lastArrivalCheckSum != checkSum {
		ChangeCheckSumFile(lastInfo, checkSum)
		telegram.SendMessage(list, t)
	}
}

func deleteEmptyElements(el []string) []string {
	var buf []string
	for _, val := range el {
		if val != "" {
			buf = append(buf, val)
		}
	}
	return buf
}

func createCheckSum(report []string) string {
	buf := bytes.Buffer{}
	for i := range report {
		buf.WriteString(report[i])
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(buf.String())))
}
