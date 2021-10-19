package commandUtilities

import (
	"strings"
)

func CommandSplit(s string) []string { // split a string along spaces, but not spaces in quotes, while preserving unicode runes.
	trimmed := strings.TrimSpace(s)
	runeConv := []rune(trimmed)
	if len(runeConv) > 0 {
		inQuotes := false
		lastidx := 0
		var result []string
		for idx, curRune := range runeConv {
			if curRune == '"' {
				inQuotes = !inQuotes
			}
			if curRune == ' ' && !inQuotes {
				result = append(result, string(runeConv[lastidx:idx]))
				lastidx = idx + 1
			}
		}
		result = append(result, string(runeConv[lastidx:]))
		return result
	} else {
		return nil
	}
}
