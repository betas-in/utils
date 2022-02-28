package utils

import "strings"

type ArrayFunctions struct {
}

func Array() ArrayFunctions {
	return ArrayFunctions{}
}

func (a ArrayFunctions) Contains(s []string, e string, exactMatch bool) int {
	for k, a := range s {
		if !exactMatch {
			a = strings.ToLower(strings.TrimSpace(a))
			e = strings.ToLower(strings.TrimSpace(e))
		}
		if a == e {
			return k
		}
	}
	return -1
}

func (a ArrayFunctions) Delete(s []string, i int) []string {
	if i == -1 {
		return s
	}
	if i > len(s) {
		return s
	}

	return append(s[:i], s[i+1:]...)
}
