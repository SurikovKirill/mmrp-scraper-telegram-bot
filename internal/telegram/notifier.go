package telegram

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// TODO: make text of the message pretty
// TODO: add logger

type MMRPBot struct {
}

type Text struct {
	date string
	data map[string]string
}

func (t *Text) ToString() string {
	result := fmt.Sprintf("%s \n", t.date)
	for key, value := range t.data {
		result += fmt.Sprintf("%s: %s \n", key, value)
	}
	fmt.Println(result)
	return result
}

func SendMessage(s string, d map[string]string) {
	cfg, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create answer")
	t := Text{s, d}
	b := []byte(fmt.Sprintf(`{"chat_id": %d, "text": "%s"}`, cfg.ChatId, t.ToString()))
	tx := bytes.NewReader(b)
	_, err = http.Post(fmt.Sprintf("%s%s/sendMessage", cfg.Url, cfg.Token), "application/json", tx)
	if err != nil {
		log.Fatal(err)
	}

}

func SendDocument(link string) {
	cfg, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create answer")
	fmt.Println(link)
	b := []byte(fmt.Sprintf(`{"chat_id": %d, "text": "%s"}`, cfg.ChatId, strings.TrimSpace(link)))
	tx := bytes.NewReader(b)
	_, err = http.Post(fmt.Sprintf("%s%s/sendMessage", cfg.Url, cfg.Token), "application/json", tx)
	if err != nil {
		log.Fatal(err)
	}
}
