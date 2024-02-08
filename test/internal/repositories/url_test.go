package repositories

import (
	"testing"

	"github.com/edgarlatorre/bookmark/internal/repositories"
)

func TestReadWithValidFile(t *testing.T) {
	urls, err := repositories.Read("../../fixtures/files/urls.json")

	if len(urls) != 2 && err == nil {
		t.Fatalf(`Read("../../fixtures/files/urls.json") = 2,  %d, error`, len(urls))
	}
}

func TestReadWhenFileDoesNotExist(t *testing.T) {
	urls, err := repositories.Read("invalid.json")

	if err == nil || urls != nil {
		t.Fatalf(`Read("invalid.json") = %s, error`, err)
	}
}
