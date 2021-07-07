package drawght

import (
	"fmt"
	"regexp"
	"reflect"
	"strings"
)

const (
	PREFIX, ATTRIBUTE, QUERY, ITEM, SUFFIX = "{", ".", ":", "#", "}"
	KEY, LIST, INDEX = "(.*?)", "(?P<list>.*?)", "(?P<index>.*?)"
)


var (
	eol = regexp.MustCompile("\r?\n")
	keyPattern = regexp.MustCompile(PREFIX + KEY + SUFFIX)
)

func Parse(template string, data map[string]interface{}) (result string) {
	return ParseKeys(ParseQueries(template, data), data)
}

func ParseQueries(template string, data map[string]interface{}) (result string) {
	return ""
}

func ParseKeys(template string, data map[string]interface{}) (result string) {
	lines := eol.Split(template, -1)
	for i := 0; i < len(lines); i++ {
		lines[i] = ParseTemplate(lines[i], data)
	}
	return strings.Join(lines[:], "\n")
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

		if reflect.TypeOf(value).Kind() == reflect.Slice {
			list := reflect.ValueOf(value)
			lines := make([]string, list.Len())
			for i := 0; i < list.Len(); i++ {
				lines[i] = parser.ReplaceAllString(template, list.Index(i).String());
			}
			result = strings.Join(lines[:], "\n")
		} else {
			result = parser.ReplaceAllString(result, value.(string))
		}
	}

	return result;
}

func getValueFromKey(key string, data map[string]interface{}) (value interface{}) {
	value = data[key]
	return value
}
