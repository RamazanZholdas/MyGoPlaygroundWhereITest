package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Genre struct {
	GenreIds  []string `json:"GenreIds"`
	GenreName []string `json:"GenreName"`
}

type IdGenre struct {
	GenreIds []string `json:"IdGenre"`
}

type Country struct {
	Name []string `json:"name"`
}

func main() {
	byteSlice1, err := os.ReadFile("./idGenre.json")
	if err != nil {
		log.Fatal(err)
	}

	byteSlice2, err := os.ReadFile("./NameGenre.json")
	if err != nil {
		log.Fatal(err)
	}

	ids4 := IdGenre{}
	countries := Country{}

	err = json.Unmarshal(byteSlice1, &ids4)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteSlice2, &countries)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(ids4.GenreIds), len(countries.Name))

	selectOption := Genre{}
	selectOption.GenreIds = ids4.GenreIds
	selectOption.GenreName = countries.Name

	file, err := os.Create("Genre.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice, err := json.MarshalIndent(selectOption, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(file.Name(), byteSlice, 0644)
}

/*
func getOptionValues() []string {
	body := `
	<option value="1750">аниме</option>
	<option value="22">биография</option>
	<option value="3">боевик</option>
	<option value="13">вестерн</option>
	<option value="19">военный</option>
	<option value="17">детектив</option>
	<option value="456">детский</option>
	<option value="20">для взрослых</option>
	<option value="12">документальный</option>
	<option value="8">драма</option>
	<option value="27">игра</option>
	<option value="23">история</option>
	<option value="6">комедия</option>
	<option value="1747">концерт</option>
	<option value="15">короткометражка</option>
	<option value="16">криминал</option>
	<option value="7">мелодрама</option>
	<option value="21">музыка</option>
	<option value="14">мультфильм</option>
	<option value="9">мюзикл</option>
	<option value="28">новости</option>
	<option value="10">приключения</option>
	<option value="25">реальное ТВ</option>
	<option value="11">семейный</option>
	<option value="24">спорт</option>
	<option value="26">ток-шоу</option>
	<option value="4">триллер</option>
	<option value="1">ужасы</option>
	<option value="2">фантастика</option>
	<option value="18">фильм-нуар</option>
	<option value="5">фэнтези</option>
	<option value="1751">церемония</option>`

	slice := []string{}
	reader := strings.NewReader(body)
	tokenizer := html.NewTokenizer(reader)
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return slice
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return nil
		}
		_, hasAttr := tokenizer.TagName()
		if hasAttr {
			for {
				_, attrValue, moreAttr := tokenizer.TagAttr()
				// if string(attrKey) == "" {
				//     break
				// }
				slice = append(slice, string(attrValue))
				if !moreAttr {
					break
				}
			}
		}
	}
}*/

/*

type SelectOption struct {
	Id   []string `json:"id"`
	Name []string `json:"name"`
}

type IdForCountry struct {
	Ids []string `json:"id"`
}

type Country struct {
	Name []string `json:"name"`
}

func main() {
	byteSlice1, err := os.ReadFile("./id.json")
	if err != nil {
		log.Fatal(err)
	}

	byteSlice2, err := os.ReadFile("./countries.json")
	if err != nil {
		log.Fatal(err)
	}

	ids4 := IdForCountry{}
	countries := Country{}

	err = json.Unmarshal(byteSlice1, &ids4)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteSlice2, &countries)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(ids4.Ids), len(countries.Name))

	selectOption := SelectOption{}
	selectOption.Id = ids4.Ids
	selectOption.Name = countries.Name

	file, err := os.Create("Country.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice, err := json.MarshalIndent(selectOption, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(file.Name(), byteSlice, 0644)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
*/
