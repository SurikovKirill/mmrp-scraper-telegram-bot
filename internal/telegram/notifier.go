package telegram

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	chatId = -1001580808284
	token  = "2060850344:AAHpEc_-JdkYdbP_p0ZoUSMC8-U0mv3_a8c"
	url    = "https://api.telegram.org/bot"
)

type BotSendMessageId struct {
	Result struct {
		Message_id int
	}
}

func SendMessage() {
	fmt.Println("create answer")
	data := []byte(fmt.Sprintf(`{"chat_id": %d, "text": "clap"}`, chatId))
	fmt.Println(string(data))
	tx := bytes.NewReader(data)
	res, err := http.Post(fmt.Sprintf("%s%s/sendMessage", url, token), "application/json", tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
	fmt.Println(res.Body)
	fmt.Println(res.Request)

}
