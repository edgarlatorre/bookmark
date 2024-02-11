package repositories

import (
	"os"
	"testing"

	"github.com/edgarlatorre/bookmark/internal/repositories"
)

func TestReadWithValidFile(t *testing.T) {
	r := repositories.UrlRepository{FilePath: "../../fixtures/files/urls.csv"}
	urls, err := r.Read()

	if len(urls) != 2 && err == nil {
		t.Fatalf(`Read("../../fixtures/files/urls.csv") = 2,  %d, error`, len(urls))
	}
}

func TestReadWhenFileDoesNotExist(t *testing.T) {
	r := repositories.UrlRepository{FilePath: "invalid.csv"}
	urls, err := r.Read()

	if err == nil || urls != nil {
		t.Fatalf(`Read("invalid.csv") = %s, error`, err)
	}
}

func TestCreateNewUrl(t *testing.T) {
	file, err := os.CreateTemp("", "urls")

	defer os.Remove(file.Name())

	r := repositories.UrlRepository{FilePath: file.Name()}

	url, err := r.Create("https://test.com", "test")

	if url.Name != "test" && url.Url != "https://test.com" && err == nil {
		t.Fatalf(`Create should return the urls created,  %s - %s, error`, url.Name, url.Url)
	}

	urls, err := r.Read()

	if len(urls) != 1 && err == nil {
		t.Fatalf(`Create should return urls with size 1,  %d, error`, len(urls))
	}
}
