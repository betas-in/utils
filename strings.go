package utils

import "strings"

type StringFunctions struct{}

func String() StringFunctions {
	return StringFunctions{}
}

func (s StringFunctions) SplitAndTrim(str string, separator string) []string {
	output := []string{}
	splits := strings.Split(str, separator)
	for _, split := range splits {
		output = append(output, strings.TrimSpace(split))
	}
	return output
}
