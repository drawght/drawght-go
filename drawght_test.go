package drawght_test

import "github.com/drawght/drawght-parser-go"

import (
	"fmt"
	"strings"
	"testing"
)

var dataset = map[string]interface{} {
	"title": "This is a title",
	"author": "Nathan Never",
	"tags": []string {
		"Template",
		"Draft",
	},
}

var templatesWithKeysOnly = map[string]string {
	"<h1>{title}</h1>": fmt.Sprintf("<h1>%s</h1>", dataset["title"]),
	"# {title} written by {author}": fmt.Sprintf("# %s written by %s", dataset["title"], dataset["author"]),
	"- {tags}": "",
}

func TestParseTemplate(t *testing.T) {
	tagsSlice := make([]string, len(dataset["tags"].([]string)))

	for i, tag := range (dataset["tags"].([]string)) {
		tagsSlice[i] = fmt.Sprintf("- %s", tag)
	}

	templatesWithKeysOnly["- {tags}"] = strings.Join(tagsSlice[:], "\n")

	for template, expected := range(templatesWithKeysOnly) {
		result := drawght.ParseTemplate(template, dataset)

		if (result != expected) {
			t.Errorf("Template '%s' not parsed", template)
		}
	}
}

func TestParseKeys(t *testing.T) {
	templateLines := make([]string, len(templatesWithKeysOnly))
	expectedLines := make([]string, len(templatesWithKeysOnly))

	i := 0
	for template, expected := range(templatesWithKeysOnly) {
		templateLines[i] = template
		expectedLines[i] = expected
		i++
	}

	template := strings.Join(templateLines[:], "\n\n")
	expected := strings.Join(expectedLines[:], "\n\n")

	result := drawght.ParseKeys(template, dataset)

	if result != expected {
		t.Errorf("Template not parsed")
		t.Errorf("Template:\n%v\n", template)
		t.Errorf("Expected:\n%v\n", expected)
		t.Errorf("Result:\n%v\n", result)
	}
}

func TestParse(t *testing.T) {
	t.Skip("Not implemented, yet")
}
