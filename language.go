package validator

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

const (
	langFilesPath = "languages/name.json"
)

type Language struct {
	Equals        string `json:"equals"`
	Regexp        string `json:"regexp"`
	Required      string `json:"required"`
	MinString     string `json:"min_string"`
	MinNumber     string `json:"min_number"`
	MaxString     string `json:"max_string"`
	MaxNumber     string `json:"max_number"`
	BetweenString string `json:"between_string"`
	BetweenNumber string `json:"between_number"`
}

func NewLanguage(lang string) *Language {
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Join(path.Dir(filename), strReplace(langFilesPath, "name", lang))
	data, _ := ioutil.ReadFile(filePath)

	var language Language
	json.Unmarshal(data, &language)

	return &language
}
