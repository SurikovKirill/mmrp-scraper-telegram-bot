package main

import "mmrp-scraper/internal/scrapers"

func main() {
	scrap := scrapers.MMRPScraper{}
	scrap.Scrape()
	//bot, err := tgbotapi.NewBotAPI("2060850344:AAHpEc_-JdkYdbP_p0ZoUSMC8-U0mv3_a8c")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//bot.Debug = true
	//updateConfig := tgbotapi.NewUpdate(0)
	//updateConfig.Timeout = 30
	//
	//// Start polling
	//updates, _ := bot.GetUpdatesChan(updateConfig)
	//
	//for update := range updates {
	//	if update.Message == nil {
	//		continue
	//	}
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//	msg.ReplyToMessageID = update.Message.MessageID
	//	if _, err := bot.Send(msg); err != nil {
	//		panic(err)
	//	}
	//}
}

//c := colly.NewCollector()
////arrivals := make([]string, 10)
////var a string
//c.OnHTML(".row", func(e *colly.HTMLElement){
//	e.ForEach(".t_1", func(_ int, el *colly.HTMLElement){
//		a := el.ChildAttr("a", "href")
//		fmt.Println(a)
//	})
//arrivals = append(arrivals, a)
//fmt.Println(arrivals)
//el := e.ChildText(".t_1")
//link := e.ChildAttr("a", "href")
//fmt.Printf(el)
//})
//c.OnRequest(func(r *colly.Request) {
//	fmt.Println("Visiting", r.URL.String())
//})
//c.Visit("http://mmrp.ru/news/74/")

//fmt.Print(arrivals[0])
