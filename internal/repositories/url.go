package repositories

import (
	"encoding/json"
	"io"
	"os"
)

type Urls struct {
	Urls []Url `json:"urls"`
}

type Url struct {
	Name string `json:"title"`
	Url  string `json:"url"`
}

func Read(filePath string) ([]Url, error) {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var urls Urls
	json.Unmarshal([]byte(byteValue), &urls)

	return urls.Urls, nil
}
