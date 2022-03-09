package translate

import (
	"encoding/json"
	"errors"
	"os"
)

type SelectOption struct {
	Id   []string `json:"id"`
	Name []string `json:"name"`
}

type Genre struct {
	GenreIds  []string `json:"GenreIds"`
	GenreName []string `json:"GenreName"`
}

type Content struct {
	Content          []string `json:"Content"`
	ContentTranslate []string `json:"ContentTranslate"`
}

func CheckGenre(str string) (string, error) {
	byteSlice, err := os.ReadFile("./Genre.json")
	if err != nil {
		return "", err
	}

	data := Genre{}

	err = json.Unmarshal(byteSlice, &data)
	if err != nil {
		return "", err
	}

	for i := range data.GenreName {
		if data.GenreName[i] == str {
			return data.GenreIds[i], nil
		}
	}

	return "", errors.New("genre does not exist")
}

func CheckCountry(str string) (string, error) {
	byteSlice, err := os.ReadFile("./Country.json")
	if err != nil {
		return "", err
	}

	data := SelectOption{}

	err = json.Unmarshal(byteSlice, &data)
	if err != nil {
		return "", err
	}

	for i := range data.Name {
		if data.Name[i] == str {
			return data.Id[i], nil
		}
	}

	return "", errors.New("country does not exist")
}

func CheckContent(str string) (string, error) {
	byteSlice, err := os.ReadFile("./Content.json")
	if err != nil {
		return "", err
	}

	data := Content{}

	err = json.Unmarshal(byteSlice, &data)
	if err != nil {
		return "", err
	}

	for i := range data.Content {
		if data.Content[i] == str {
			return data.ContentTranslate[i], nil
		}
	}

	return "", errors.New("content-type does not exist")
}
