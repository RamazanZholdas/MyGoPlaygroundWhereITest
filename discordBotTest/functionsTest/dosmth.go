package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	var res string
	c := colly.NewCollector()
	c.OnHTML("p.styles_root__3jjhA", func(h *colly.HTMLElement) {
		res = h.Text
	})
	c.Visit("https://www.kinopoisk.ru/film/260162/sr/1/")
	fmt.Println(res)
}

/*
str := "Аардавак"
	strLow := ""
	for _, v := range str {
		if v < 1072 {
			v += 32
		}
		strLow += strconv.QuoteRune(v)
	}
	fmt.Printf("%T\n", str)
	fmt.Println(strings.Trim(strLow, "'"))
*/
