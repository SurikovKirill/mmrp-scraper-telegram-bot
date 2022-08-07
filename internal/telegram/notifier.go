package telegram

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

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

func SendMessage(s string, d map[string]string) {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendMessage", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	cfg, err := Init()
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendMessage", "error": err}).Error(err)
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
	b := []byte(fmt.Sprintf(`{"chat_id": %d, "document": "%s"}`, cfg.ChatId, strings.TrimSpace(link)))
	tx := bytes.NewReader(b)
	_, err = http.Post(fmt.Sprintf("%s%s/sendDocument", cfg.Url, cfg.Token), "application/json", tx)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendDocument", "error": err}).Error(err)
	}
}

func SendMessage1(s string, d map[string]string) {
	f, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendMessage", "error": err}).Error("Error in opening log file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	cfg, err := Init()
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendMessage", "error": err}).Error(err)
	}
	p := "<thead><tr><th>Тип</th><th>№</th> <th>Судно, Флаг, Агент</th><th>Контр.срок подтв. заявки</th><th>Уточн. времяНР</th><th>Время/место прихода/отхода</th><th>Время/ местоГКО</th><th>Длина LOA</th><th>Маршрут</th><th>Цель захода</th><th>План. груз</th><th>Факт. груз</th><th>Заявл. время</th><th>Буксиры</th><th>Осадка нос / корма</th><th>Примечание</th></tr></thead>"

	// t := Text{s, d}
	b := []byte(fmt.Sprintf(`{"chat_id": %d, "text": "%s", "parse_mode": "HTML"}`, cfg.ChatId, p))
	tx := bytes.NewReader(b)
	_, err = http.Post(fmt.Sprintf("%s%s/sendMessage", cfg.Url, cfg.Token), "application/json", tx)
	if err != nil {
		log.WithFields(log.Fields{"package": "scrapers", "function": "SendMessage", "error": err}).Error(err)
	}
}
