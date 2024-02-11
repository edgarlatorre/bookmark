package models

import (
	"strings"
	"testing"

	"github.com/edgarlatorre/bookmark/internal/models"
)

func TestFormModelView(t *testing.T) {
	m := models.NewFormModel()
	content := m.View()

	if !strings.Contains(content, "Url") {
		t.Fatalf(`View() does not contain Url,  %s, error`, content)
	}
}
