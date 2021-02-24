package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var (
	regSpace      = regexp.MustCompile(`[\t\n\r]`)
	regNotLetters = regexp.MustCompile(`[^\p{L} /gu]`)
)

func Top10(text string) []string {
	text = formatText(text)

	if isEmpty := isTextEmpty(text); isEmpty {
		return []string{}
	}

	words := strings.Fields(text)

	wordOcc := make(map[string]int)
	for _, r := range words {
		wordOcc[r]++
	}

	sort.Slice(words, func(i, j int) bool {
		if wordOcc[words[i]] == wordOcc[words[j]] {
			return words[i] < words[j]
		}

		return wordOcc[words[i]] > wordOcc[words[j]]
	})

	top10 := takeFirst10(words)

	return top10
}

func formatText(text string) string {
	text = regSpace.ReplaceAllString(text, " ")

	text = regNotLetters.ReplaceAllString(text, "")

	text = strings.ToLower(text)

	return text
}

func takeFirst10(words []string) []string {
	top10 := []string{words[0]}

	for _, word := range words {
		if len(top10) == 10 {
			break
		}

		if word == top10[len(top10)-1] {
			continue
		}

		top10 = append(top10, word)
	}

	return top10
}

func isTextEmpty(text string) bool {
	isEmpty, _ := regexp.MatchString(`^\s*$`, text)

	return isEmpty
}
