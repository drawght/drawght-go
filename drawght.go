package drawght

import (
	"fmt"
	"regexp"
	"reflect"
	"strings"
	"strconv"
)

const (
	PREFIX, ATTRIBUTE, QUERY, ITEM, SUFFIX = "{", ".", ":", "#", "}"
	KEY, LIST, INDEX = "(.*?)", "(.*?)", "(.*?)"
)


var (
	eol = regexp.MustCompile("\r?\n")
	keyPattern = regexp.MustCompile(PREFIX + KEY + SUFFIX)
	listPattern = regexp.MustCompile(PREFIX + LIST + QUERY + KEY + SUFFIX)
	itemPattern = regexp.MustCompile(LIST + ITEM + INDEX)
)

func Parse(template string, data map[string]interface{}) (result string) {
	return ParseKeys(ParseQueries(template, data), data)
}

func ParseQueries(template string, data map[string]interface{}) (result string) {
	lines := eol.Split(template, -1)
	resultLines := make([]string, len(lines))

	for l := 0; l < len(lines); l++ {
		line := lines[l]
		queries := listPattern.FindAllString(lines[l], -1)
		resultLines[l] = line
		if len(queries) > 0 {
			keyCleaner := regexp.MustCompile(fmt.Sprintf("[%s%s]", PREFIX, SUFFIX))
			for q := 0; q < len(queries); q++ {
				query := queries[q]
				if listPattern.MatchString(query) {
					key := keyCleaner.ReplaceAllString(query, "")
					filter := strings.Split(key, QUERY)
					list := filter[0]
					key = filter[1]
					queryCleaner := regexp.MustCompile(fmt.Sprintf("%s%s", list, QUERY))
					if value := getValueFromKey(list, data); value != nil {
						switch value.(type) {
							case []interface{}:
								list := value.([]interface{})
								partial := queryCleaner.ReplaceAllString(line, "")
								parsedLines := make([]string, len(list))
								for l := 0; l < len(list); l++ {
									parsedLines[l] = fmt.Sprintf("%v", ParseTemplate(partial, list[l].(map[string]interface{})))
								}
								resultLines[l] = strings.Join(parsedLines[:], "\n")
							default:
								resultLines[l] = fmt.Sprintf("%v", value)
						}
					}
				}
			}
		}
	}

	return strings.Join(resultLines[:], "\n")
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
	cleaner := regexp.MustCompile(fmt.Sprintf("[%s%s]", PREFIX, SUFFIX))

	for _, templateKey := range templateKeys {
		parser := regexp.MustCompile(templateKey)
		key := cleaner.ReplaceAllString(templateKey, "")
		value := getValueFromKey(key, data)

		if value == nil { value = templateKey }

		if reflect.TypeOf(value).Kind() == reflect.Slice {
			list := reflect.ValueOf(value)
			lines := make([]string, list.Len())
			for i := 0; i < list.Len(); i++ {
				lines[i] = parser.ReplaceAllString(template, fmt.Sprintf("%v", list.Index(i)));
			}
			result = strings.Join(lines[:], "\n")
		} else {
			result = parser.ReplaceAllString(result, value.(string))
		}
	}

	return result;
}

func getValueFromKey(nestedKey string, data map[string]interface{}) (value interface{}) {
	dataset := data
	if keys := strings.Split(nestedKey, ATTRIBUTE); len(keys) >= 1 {
		for i := 0; i < len(keys); i++ {
			key, index := keys[i], -1

			if itemPattern.MatchString(key) {
				item := strings.Split(key, ITEM)
				key = item[0]
				index, _ = strconv.Atoi(item[1])
				list := dataset[key].([]interface{})
				if index <= len(list) {
					value = list[index - 1]
				} else {
					value = nil
				}
			} else {
				value = dataset[key]
			}

			switch value.(type) {
				case map[string]interface{}:
					dataset = value.(map[string]interface{})
				case []interface{}:
					value = value.([]interface{})
				default:
					value = value
			}
		}
		return value
	}

	value = data[nestedKey]

	return value
}
