package repositories

import (
	"encoding/csv"
	"os"
)

type Url struct {
	Name string
	Url  string
}

func (i Url) Title() string       { return i.Name }
func (i Url) Description() string { return i.Url }
func (i Url) FilterValue() string { return i.Name }

type UrlRepository struct {
	FilePath string
}

func (r UrlRepository) Read() ([]Url, error) {
	jsonFile, err := os.Open(r.FilePath)

	defer jsonFile.Close()

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(jsonFile)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	urls := make([]Url, len(records))

	for i, l := range records {
		urls[i] = Url{Name: l[0], Url: l[1]}
	}

	return urls, nil
}

func (r UrlRepository) Create(link string, title string) (Url, error) {
	jsonFile, err := os.OpenFile(r.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return Url{}, err
	}

	url := Url{Name: title, Url: link}

	_, err = jsonFile.WriteString(url.Name + "," + url.Url + "\n")

	if err != nil {
		return Url{}, err
	}

	defer jsonFile.Close()

	return url, nil
}
