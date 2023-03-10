package telegram

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const (
	telegramBotAPIURL = "https://api.telegram.org/bot"
	filename          = "temp.pdf"
)

func SendMessage(s string, d map[string]string, c *Config) {
	log.Println("Send MMRP data to chat")
	t := Text{s, d}
	tb := bytes.NewReader([]byte(fmt.Sprintf(`{"chat_id": %d, "text": "%s"}`, c.ChatID, t.ToString())))
	res, err := http.Post(fmt.Sprintf("%s%s/sendMessage", telegramBotAPIURL, c.Token), "application/json", tb)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
}

func SendDocumentRod(c *Config) {
	log.Println("Send MAPM data to chat")
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("chat_id", strconv.Itoa(c.ChatID))
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	part, err := writer.CreateFormFile("document", filepath.Base(filename))
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := io.Copy(part, file); err != nil {
		log.Println(err)
		return
	}

	if err := writer.Close(); err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s/sendDocument", telegramBotAPIURL, c.Token), payload)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if _, err := client.Do(req); err != nil {
		log.Println(err)
		return
	}
}
