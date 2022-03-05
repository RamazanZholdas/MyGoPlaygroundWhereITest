package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

/*https://www.kinopoisk.ru/index.php?level=7&from=forma&result=adv&m_act%5Bfrom%5D=forma&m_act
%5Bwhat%5D=content&m_act%5Bfind%5D=Lalalend&m_act%5Byear%5D=2017&m_act%5Bcountry%5D=1&m_act%5Bgenre%5D%5B%5D=11
&m_act%5Bactor%5D=Rayan&m_act%5Bcast%5D=Roy&m_act%5Bcontent_find%5D=film&m_act%5Bgenre_and%5D=on
*/

/*
https://www.kinopoisk.ru/index.php?level=7&from=forma&result=adv&m_act%5Bfrom%5D=forma&m_act
%5Bwhat%5D=content&m_act%5Bfind%5D=Texas&m_act%5Byear%5D=2012&m_act%5Bcountry%5D=136&m_act%5Bgenre%5D%5B%5D=11
&m_act%5Bactor%5D=Dylan&m_act%5Bcast%5D=Roy&m_act%5Bcontent_find%5D=serial&m_act%5Bgenre_and%5D=on
*/

func main() {
	c := colly.NewCollector()
	//countries := make(map[string]string)
	//slice := []string{}
	sliceVal := []string{}

	extensions.RandomUserAgent(c)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("We are visiting this site:", r.Request.URL)
	})

	c.OnHTML("select.text>option", func(h *colly.HTMLElement) {
		//countries[h.ChildAttr("option", "value")] = h.ChildText("option")
		sliceVal = append(sliceVal, h.Text)
	})

	c.Visit("http://localhost:8080")

	//fmt.Println(slice)
	fmt.Println(sliceVal)

	file, err := os.Open("countries.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	cc := Country{}
	for i := 1; i < len(sliceVal); i++ {
		cc.Name = append(cc.Name, sliceVal[i])
	}

	byteSlice, err := json.Marshal(cc)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(file.Name(), []byte(byteSlice), 0644)
	//fmt.Println(countries)
}

type Country struct {
	Name []string `json:"name"`
}

/*
func main() {
	suggested := []string{}
	movie, rating, urlToVisit := "Dallas", "", ""
	name, nextUrl, timing, info := "", "", "", ""
	VisitUrl := "https://www.kinopoisk.ru/index.php?level=7&from=forma&result=adv&m_act%5Bfrom%5D=forma&m_act%5Bwhat%5D=content&m_act%5Bfind%5D=" + movie
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("We are visiting this site:", r.Request.URL)
	})

	c.OnHTML(".element.most_wanted", func(h *colly.HTMLElement) {
		name = h.ChildText(".info>p.name")
		timing = h.ChildText("span:nth-child(2)")
		info = h.ChildText("span:nth-child(3)")
		rating = h.ChildText(".right>div")
		nextUrl = h.ChildAttr("a", "href")
		nextUrl = "https://www.kinopoisk.ru" + nextUrl[0:len(nextUrl)-5]
	})

	c.OnHTML(".search_results search_results_last>p.header", func(h *colly.HTMLElement) {
		urlToVisit = "https://www.kinopoisk.ru" + h.ChildAttr("a", "href")
	})

	c.Visit(VisitUrl)

	fmt.Println(name, "\n", timing[4:], "\n", info, "\n", rating)
	fmt.Println("for detailed information u can visit this site:", nextUrl)
	fmt.Println(suggested, urlToVisit)
}
*/
