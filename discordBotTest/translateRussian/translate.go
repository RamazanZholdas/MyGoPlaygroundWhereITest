package translate

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	cipher   = []string{}
	cirillic = []rune("йцукенгшщзхъфывапролджэячсмитьбю")
	decode   = "E9F6F3EAE5EDE3F8F9E7F5FAF4FBE2E0EFF0EEEBE4E6FDFFF7F1ECE8F2FCE1FE"
)

type SelectOption struct {
	Id   []string `json:"id"`
	Name []string `json:"name"`
}

func initialize() {
	for i := 0; i < len(decode); i += 2 {
		cipher = append(cipher, decode[i:i+2])
	}
}

func Translating(str string) string {
	initialize()
	res := ""
	for _, v := range str {
		for i, g := range cirillic {
			if v == g {
				res += "%" + cipher[i]
			}
		}
	}
	return res
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
