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
	listPattern = regexp.MustCompile(PREFIX + LIST + QUERY + KEY + SUFFIX)
)

func Parse(template string, data map[string]interface{}) (result string) {
	return ParseKeys(ParseQueries(template, data), data)
}

func ParseQueries(template string, data map[string]interface{}) (result string) {
	lines := eol.Split(template, -1)
	for i := 0; i < len(lines); i++ {
		lines[i] = ParseKeys(lines[i], data)
	}
	return strings.Join(lines[:], "\n")
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

func getValueFromKey(nestedKey string, data map[string]interface{}) (value interface{}) {
	if keys := strings.Split(nestedKey, ATTRIBUTE); len(keys) > 1 {
		value = data[keys[0]]
		mapPattern := regexp.MustCompile(`map\[string\]*`)
		for i := 1; i < len(keys); i++ {
			valueType := fmt.Sprintf("%T", value)
			if mapPattern.MatchString(valueType) {
				value = getValueFromKey(keys[i], value.(map[string]interface{}))
			} else {
				return nil
			}
		}

		return value
	}

	value = data[nestedKey]

	return value
}
