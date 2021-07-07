package drawght_test

import "github.com/drawght/drawght-parser-go"

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	dataset := map[string]interface{} {
		"title": "This is a title",
		"author": "Nathan Never",
		"tags": []string {
			"Template",
			"Draft",
		},
	}

	templates := map[string]string {
		"<h1>{title}</h1>": fmt.Sprintf("<h1>%s</h1>", dataset["title"]),
		"# {title} written by {author}": fmt.Sprintf("# %s written by %s", dataset["title"], dataset["author"]),
		"- {tags}": "",
	}

	tagsSlice := make([]string, len(dataset["tags"].([]string)))

	for i, tag := range (dataset["tags"].([]string)) {
		tagsSlice[i] = fmt.Sprintf("- %s", tag)
	}

	templates["- {tags}"] = strings.Join(tagsSlice[:], "\n")

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
