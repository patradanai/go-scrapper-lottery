package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type LotteryLink struct {
	Link string
	Date string
}

func FindAllDate() ([]LotteryLink, error) {
	url := "https://lottery.kapook.com/history"
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	paramLink := make([]LotteryLink, 0)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("main[class=article-main]", func(e *colly.HTMLElement) {
		e.ForEach("section[class=history-check]", func(_ int, h *colly.HTMLElement) {
			h.ForEach("ul > li", func(_ int, l *colly.HTMLElement) {
				fmt.Printf("Date : %v, %v \n", l.ChildText("strong"), l.ChildAttr("a", "href"))
				param := LotteryLink{}
				param.Date = l.ChildText("a > strong")
				param.Link = "https://lottery.kapook.com" + l.ChildAttr("a", "href")
				paramLink = append(paramLink, param)
			})
		})
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit(url)

	jsonData, err := json.MarshalIndent(paramLink, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	return paramLink, nil
}
