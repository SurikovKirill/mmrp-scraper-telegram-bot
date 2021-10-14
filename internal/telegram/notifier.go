package telegram

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// TODO: make text of the message pretty
// TODO: add logger

const (
	chatId = -1001580808284
	token  = "2060850344:AAHpEc_-JdkYdbP_p0ZoUSMC8-U0mv3_a8c"
	url    = "https://api.telegram.org/bot"
)

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
	fmt.Println("create answer")
	t := Text{s, d}
	b := []byte(fmt.Sprintf(`{"chat_id": %d, "text": "%s"}`, chatId, t.ToString()))
	tx := bytes.NewReader(b)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", url, token), "application/json", tx)
	if err != nil {
		log.Fatal(err)
	}

}
