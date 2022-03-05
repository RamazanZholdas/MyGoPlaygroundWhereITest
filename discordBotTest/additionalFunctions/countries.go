package additionalfunctions

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ladno() {
	c := colly.NewCollector()
	countries := make(map[string]string)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("We are visiting this site:", r.Request.URL)
	})

	c.OnHTML("select.text.el_5 __countrySB__", func(h *colly.HTMLElement) {
		countries[h.ChildAttr("option", "value")] = h.ChildText("option")
	})

	c.Visit("https://www.kinopoisk.ru/s/")

	fmt.Println(countries)
}
