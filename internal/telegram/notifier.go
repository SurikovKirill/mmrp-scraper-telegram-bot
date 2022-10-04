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
	documentName      = "temp.pdf"
)

func SendMessage(s string, d map[string]string, c *Config) {
	t := Text{s, d}
	tb := bytes.NewReader([]byte(fmt.Sprintf(`{"chat_id": %d, "text": "%s"}`, c.ChatID, t.ToString())))
	res, err := http.Post(fmt.Sprintf("%s%s/sendMessage", telegramBotAPIURL, c.Token), "application/json", tb)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
}

func SendDocumentRod(c *Config) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	defer writer.Close()
	if err := writer.WriteField("chat_id", strconv.Itoa(c.ChatID)); err != nil {
		log.Println(err)
	}
	file, err := os.Open(documentName)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	buff, err := writer.CreateFormFile("document", filepath.Base(documentName))
	if err != nil {
		log.Println(err)
	}
	if _, err := io.Copy(buff, file); err != nil {
		log.Println(err)
	}
	res, err := http.Post(fmt.Sprintf("%s%s/sendDocument", telegramBotAPIURL, c.Token), "multipart/form-data", payload)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
}

// func SendDocumentRod(c *Config) {
// 	url := fmt.Sprintf("%s%s/sendDocument", telegramBotApiUrl, c.Token)
// 	method := "POST"
// 	payload := &bytes.Buffer{}
// 	writer := multipart.NewWriter(payload)
// 	_ = writer.WriteField("chat_id", strconv.Itoa(c.ChatId))
// 	file, errFile2 := os.Open("temp.pdf")
// 	if errFile2 != nil {
// 		log.Println(errFile2)
// 		return
// 	}

// 	defer file.Close()
// 	part2, errFile2 := writer.CreateFormFile("document", filepath.Base("temp.pdf"))
// 	if errFile2 != nil {
// 		log.Println(errFile2)
// 		return
// 	}
// 	_, errFile3 := io.Copy(part2, file)
// 	if errFile3 != nil {
// 		log.Println(errFile3)
// 		return
// 	}
// 	err := writer.Close()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	res, err := client.Do(req)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	log.Println(string(body))
// }
