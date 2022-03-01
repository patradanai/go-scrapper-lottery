package utils

import (
	"encoding/json"
	"fmt"
	"lottery-web-scrapping/models"
	"lottery-web-scrapping/utils"
	"time"

	"github.com/gocolly/colly"
)

type LotteryLink struct {
	Link string
	Date string
}

type IWebScrapping interface {
	FindAllDate() ([]LotteryLink, error)
	FindByDate(link string)
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
				// fmt.Printf("Date : %v, %v \n", l.ChildText("strong"), l.ChildAttr("a", "href"))
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

	_, err := json.MarshalIndent(paramLink, "", " ")
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(jsonData))

	return paramLink, nil
}

func FindByDate(link string) {
	url := link

	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	drawingLot := &models.DrawingLottery{}

	c.OnHTML("main[class=article-main]", func(h *colly.HTMLElement) {

		// First Prize
		h.ForEach("div[class=prize]", func(i int, p *colly.HTMLElement) {
			switch p.ChildText("h4") {
			case "รางวัลที่ 1":
				firstPrize := &models.LotteryType{}
				firstPrize.Name = p.ChildText("h4")
				firstPrize.Prize = utils.Filter(p.ChildText("em"))
				firstPrize.Lottery = append(firstPrize.Lottery, utils.ConvToInteger(p.ChildText("strong")))
				drawingLot.FirstPrize = *firstPrize
			case "เลขหน้า 3 ตัว":
				frontThridPrize := &models.LotteryType{}
				frontThridPrize.Name = p.ChildText("h4")
				frontThridPrize.Prize = utils.Filter(p.ChildText("em"))
				frontThridPrize.Lottery = append(frontThridPrize.Lottery, utils.ConvToInteger(p.ChildText("strong")))
				drawingLot.FrontThird = *frontThridPrize
			case "เลขท้าย 3 ตัว":
				endThridPrize := &models.LotteryType{}
				endThridPrize.Name = p.ChildText("h4")
				endThridPrize.Prize = utils.Filter(p.ChildText("em"))
				endThridPrize.Lottery = append(endThridPrize.Lottery, utils.ConvToInteger(p.ChildText("strong")))
				drawingLot.EndThird = *endThridPrize
			case "เลขท้าย 2 ตัว":
				endSecondPrize := &models.LotteryType{}
				endSecondPrize.Name = p.ChildText("h4")
				endSecondPrize.Prize = utils.Filter(p.ChildText("em"))
				endSecondPrize.Lottery = append(endSecondPrize.Lottery, utils.ConvToInteger(p.ChildText("strong")))
				drawingLot.EndSecond = *endSecondPrize
			}
		})

		// Semilar First Prize
		semilarFirstPrize := &models.LotteryType{}
		semilarFirstPrize.Name = h.ChildText("section.another-first-lottery.lottery-similar-first-prize h4")
		semilarFirstPrize.Prize = utils.Filter(h.ChildText("section.another-first-lottery.lottery-similar-first-prize em"))
		h.ForEach("section.another-first-lottery.lottery-similar-first-prize > strong", func(_ int, e *colly.HTMLElement) {
			semilarFirstPrize.Lottery = append(semilarFirstPrize.Lottery, utils.ConvToInteger(e.Text))
		})

		drawingLot.NearFirstPrize = *semilarFirstPrize

		// Second Prize
		secondPrize := &models.LotteryType{}
		secondPrize.Name = h.ChildText("section.another-lottery.lottery-second-prize h4")
		secondPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-second-prize em"))
		h.ForEach("section.another-lottery.lottery-second-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			secondPrize.Lottery = append(secondPrize.Lottery, utils.ConvToInteger(e.Text))
		})

		drawingLot.SecondPrize = *secondPrize

		// Third Prize
		thridPrize := &models.LotteryType{}
		thridPrize.Name = h.ChildText("section.another-lottery.lottery-third-prize h4")
		thridPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-third-prize em"))
		h.ForEach("section.another-lottery.lottery-third-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			thridPrize.Lottery = append(thridPrize.Lottery, utils.ConvToInteger(e.Text))
		})

		drawingLot.ThridPrize = *thridPrize

		// Fourth Prize
		fourthPrize := &models.LotteryType{}
		fourthPrize.Name = h.ChildText("section.another-lottery.lottery-fourth-prize h4")
		fourthPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-fourth-prize em"))
		h.ForEach("section.another-lottery.lottery-fourth-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			fourthPrize.Lottery = append(fourthPrize.Lottery, utils.ConvToInteger(e.Text))
		})

		drawingLot.FourthPrize = *fourthPrize

		// Fifth Prize
		fifthPrize := &models.LotteryType{}
		fifthPrize.Name = h.ChildText("section.another-lottery.lottery-fifth-prize h4")
		fifthPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-fifth-prize em"))
		h.ForEach("section.another-lottery.lottery-fifth-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			fifthPrize.Lottery = append(fifthPrize.Lottery, utils.ConvToInteger(e.Text))
		})

		drawingLot.FifthPrize = *fifthPrize
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit(url)

	dataJson, err := json.MarshalIndent(drawingLot, "", " ")
	if err != nil {
		panic(err)
	}

	println(string(dataJson))
}
