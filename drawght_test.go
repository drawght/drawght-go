package drawght_test

import "github.com/drawght/drawght-parser-go"

import (
	"fmt"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	dataset := map[string]interface{} {
		"title": "This is a title",
		"author": "Nathan Never",
	}

	templates := map[string]string {
		"<h1>{title}</h1>": fmt.Sprintf("<h1>%s</h1>", dataset["title"]),
		"# {title} written by {author}": fmt.Sprintf("# %s written by %s", dataset["title"], dataset["author"]),
	}

	for template, expected := range(templates) {
		result := drawght.ParseTemplate(template, dataset)

		if (result != expected) {
			t.Errorf("Template '%s' not parsed", template)
		}
	}
}

func TestParse(t *testing.T) {
	return
}
