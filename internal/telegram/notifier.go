package telegram

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Text struct {
	date string
	data map[string]string
}

func (t *Text) ToString() string {
	result := fmt.Sprintf("%s \n", t.date)
	for key, value := range t.data {
		result += fmt.Sprintf("%s: %s \n", key, value)
	}
	return result
}

// TODO: рефакторинг модуля

func SendMessage(s string, d map[string]string) {
	cfg, err := Init()
	if err != nil {
		log.Fatal()
	}
	t := Text{s, d}
	b := []byte(fmt.Sprintf(`{"chat_id": %d, "text": "%s"}`, cfg.ChatId, t.ToString()))
	tx := bytes.NewReader(b)
	_, err = http.Post(fmt.Sprintf("%s%s/sendMessage", cfg.Url, cfg.Token), "application/json", tx)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendMessage", "error": err}).Error(err)
	}
}

func SendDocument(link string) {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendDocument", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	cfg, err := Init()
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendDocument", "error": err}).Error(err)
	}

	log.Println(link)
	b := []byte(fmt.Sprintf(`{"chat_id": %d, "document": "/E:/mmrp-scraper-telegram-bot/tmp.html"}`, cfg.ChatId))
	tx := bytes.NewReader(b)
	r, err := http.Post(fmt.Sprintf("%s%s/sendDocument", cfg.Url, cfg.Token), "multipart/form-data", tx)
	bodyBytes, err := io.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString, r.Status)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendDocument", "error": err}).Error(err)
	}
}

func SendDocumentRod() {
	fmt.Println("sending")
	cfg, err := Init()
	url := fmt.Sprintf("%s%s/sendDocument", cfg.Url, cfg.Token)
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("chat_id", strconv.Itoa(cfg.ChatId))
	file, errFile2 := os.Open("temp.pdf")
	if errFile2 != nil {
		fmt.Println(errFile2, "zzz")
		return
	}

	defer file.Close()
	part2, errFile2 := writer.CreateFormFile("document", filepath.Base("temp.pdf"))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2, "ggg")
		return
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err, "ttt")
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
