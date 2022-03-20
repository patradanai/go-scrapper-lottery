package webscrapper

import (
	"encoding/json"
	"fmt"
	"log"
	"lottery-web-scrapping/internal/models"
	"lottery-web-scrapping/pkg/utils"
	"time"

	"github.com/gocolly/colly"
)

func GetLotteryByDate(link string) (*models.DrawingLottery, error) {
	url := link

	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	drawingLot := &models.DrawingLottery{}

	c.OnHTML("main[class=article-main]", func(h *colly.HTMLElement) {

		// Get Date
		date := utils.GetDate(h.ChildText("section.hilight-lottery > hgroup > h3"))

		// Shift Pad
		dayString := fmt.Sprintf("%02d", utils.ConvToInteger(date["Day"]))
		monthlyString := fmt.Sprintf("%02d", utils.ConvMonthlyToNum(date["Monthly"]))
		yearString := fmt.Sprintf("%04d", utils.ConvToInteger(date["Year"]))

		drawingLot.Day = dayString
		drawingLot.Month = monthlyString
		drawingLot.Year = yearString
		drawingLot.FullDate = fmt.Sprintf("%v%v%v", dayString, monthlyString, yearString)

		// First Prize
		h.ForEach("div[class=prize]", func(i int, p *colly.HTMLElement) {
			switch p.ChildText("h4") {
			case "รางวัลที่ 1":
				firstPrize := &models.LotteryType{}
				firstPrize.Name = p.ChildText("h4")
				firstPrize.Prize = utils.Filter(p.ChildText("em"))
				firstPrize.Lottery = append(firstPrize.Lottery, p.ChildText("strong"))
				drawingLot.FirstPrize = *firstPrize
			case "เลขหน้า 3 ตัว":
				frontThirdPrize := &models.LotteryType{}
				frontThirdPrize.Name = p.ChildText("h4")
				frontThirdPrize.Prize = utils.Filter(p.ChildText("em"))
				frontThirdPrize.Lottery = append(frontThirdPrize.Lottery, p.ChildText("strong"))
				drawingLot.FrontThird = *frontThirdPrize
			case "เลขท้าย 3 ตัว":
				endThirdPrize := &models.LotteryType{}
				endThirdPrize.Name = p.ChildText("h4")
				endThirdPrize.Prize = utils.Filter(p.ChildText("em"))
				endThirdPrize.Lottery = append(endThirdPrize.Lottery, p.ChildText("strong"))
				drawingLot.EndThird = *endThirdPrize
			case "เลขท้าย 2 ตัว":
				endSecondPrize := &models.LotteryType{}
				endSecondPrize.Name = p.ChildText("h4")
				endSecondPrize.Prize = utils.Filter(p.ChildText("em"))
				endSecondPrize.Lottery = append(endSecondPrize.Lottery, p.ChildText("strong"))
				drawingLot.EndSecond = *endSecondPrize
			}
		})

		// Similar First Prize
		similarFirstPrize := &models.LotteryType{}
		similarFirstPrize.Name = h.ChildText("section.another-first-lottery.lottery-similar-first-prize h4")
		similarFirstPrize.Prize = utils.Filter(h.ChildText("section.another-first-lottery.lottery-similar-first-prize em"))
		h.ForEach("section.another-first-lottery.lottery-similar-first-prize > strong", func(_ int, e *colly.HTMLElement) {
			similarFirstPrize.Lottery = append(similarFirstPrize.Lottery, e.Text)
		})

		drawingLot.NearFirstPrize = *similarFirstPrize

		// Second Prize
		secondPrize := &models.LotteryType{}
		secondPrize.Name = h.ChildText("section.another-lottery.lottery-second-prize h4")
		secondPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-second-prize em"))
		h.ForEach("section.another-lottery.lottery-second-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			secondPrize.Lottery = append(secondPrize.Lottery, e.Text)
		})

		drawingLot.SecondPrize = *secondPrize

		// Third Prize
		thirdPrize := &models.LotteryType{}
		thirdPrize.Name = h.ChildText("section.another-lottery.lottery-third-prize h4")
		thirdPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-third-prize em"))
		h.ForEach("section.another-lottery.lottery-third-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			thirdPrize.Lottery = append(thirdPrize.Lottery, e.Text)
		})

		drawingLot.ThridPrize = *thirdPrize

		// Fourth Prize
		fourthPrize := &models.LotteryType{}
		fourthPrize.Name = h.ChildText("section.another-lottery.lottery-fourth-prize h4")
		fourthPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-fourth-prize em"))
		h.ForEach("section.another-lottery.lottery-fourth-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			fourthPrize.Lottery = append(fourthPrize.Lottery, e.Text)
		})

		drawingLot.FourthPrize = *fourthPrize

		// Fifth Prize
		fifthPrize := &models.LotteryType{}
		fifthPrize.Name = h.ChildText("section.another-lottery.lottery-fifth-prize h4")
		fifthPrize.Prize = utils.Filter(h.ChildText("section.another-lottery.lottery-fifth-prize em"))
		h.ForEach("section.another-lottery.lottery-fifth-prize > div > strong", func(_ int, e *colly.HTMLElement) {
			fifthPrize.Lottery = append(fifthPrize.Lottery, e.Text)
		})

		drawingLot.FifthPrize = *fifthPrize
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("Got this error:", e)
	})

	if err := c.Visit(url); err != nil {
		return nil, err
	}

	return drawingLot, nil
}

func FindAllDrawingDate() ([]models.LotteryLink, error) {
	url := "https://lottery.kapook.com/history"
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	paramLink := make([]models.LotteryLink, 0)

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML("main[class=article-main]", func(e *colly.HTMLElement) {
		e.ForEach("section[class=history-check]", func(_ int, h *colly.HTMLElement) {
			h.ForEach("ul > li", func(_ int, l *colly.HTMLElement) {
				// log.Printf("Date : %v, %v \n", l.ChildText("strong"), l.ChildAttr("a", "href"))
				param := models.LotteryLink{}
				param.Date = l.ChildText("a > strong")
				param.Link = "https://lottery.kapook.com" + l.ChildAttr("a", "href")
				paramLink = append(paramLink, param)
			})
		})
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("Got this error:", e)
	})

	c.Visit(url)

	_, err := json.MarshalIndent(paramLink, "", " ")
	if err != nil {
		panic(err)
	}

	return paramLink, nil
}
