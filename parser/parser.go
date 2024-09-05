package parser

import "strings"

func ParseCommand(input string) (string, []string) {
	parts := strings.Fields(input)
	if (len(parts) > 0) {
		return strings.ToUpper(parts[0]), parts[1:]
	} 
	return "", []string{}
}