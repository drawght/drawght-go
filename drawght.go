package drawght

import (
	"fmt"
	"regexp"
)

const (
	PREFIX, ATTRIBUTE, QUERY, ITEM, SUFFIX = "{", ".", ":", "#", "}"
	KEY, LIST, INDEX = "(.*?)", "(?P<list>.*?)", "(?P<index>.*?)"
)


var (
	eol = regexp.MustCompile("/\r?\n/")
	keyPattern = regexp.MustCompile(PREFIX + KEY + SUFFIX)
)

func Parse(template string, data map[string]interface{}) (result string) {
	return ParseKeys(ParseQueries(template, data), data)
}

func ParseQueries(template string, data map[string]interface{}) (result string) {
	return ""
}

func ParseKeys(template string, data map[string]interface{}) (result string) {
	return ParseTemplate(template, data)
}

func ParseTemplate(template string, data map[string]interface{}) (result string) {
  result = template
	templateKeys := keyPattern.FindAllString(template, -1)
	for _, templateKey := range templateKeys {
		var (
			cleaner = regexp.MustCompile(fmt.Sprintf("[%s%s]", PREFIX, SUFFIX))
			parser = regexp.MustCompile(templateKey)
		)

    var (
			key = cleaner.ReplaceAllString(templateKey, "")
			value = getValueFromKey(key, data)
		)

		if value == nil { value = templateKey }

		result = parser.ReplaceAllString(result, value.(string))
  }

  return result;
}

func getValueFromKey(key string, data map[string]interface{}) (value interface{}) {
	value = data[key]
	return value
}
