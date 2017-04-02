package corpus

import (
	"regexp"
	"sort"
	"strings"
)

// Analyze generates a sorted word frequency slice from a string
func Analyze(text string) WordFrequencies {
	analysis := structure(wordMap(text))
	sort.Sort(sort.Reverse(analysis))
	return analysis
}

func structure(wordMap map[string]int) WordFrequencies {
	var frequencies WordFrequencies
	for word, frequency := range wordMap {
		frequencies = append(frequencies, WordFrequency{
			Word:  word,
			Count: frequency,
		})
	}
	return frequencies
}

func wordMap(text string) map[string]int {
	wordMap := make(map[string]int)

	for _, word := range wordList(text) {
		wordMap[word]++
	}

	return wordMap
}

func wordList(text string) []string {
	return strings.Split(sanitize(text), " ")
}

func sanitize(text string) string {
	result := replace(text, "[\\W]", " ")
	result = replace(result, "[\\s]+", " ")
	return strings.Trim(result, " ")
}

func replace(haystack, needle, replacement string) string {
	r, _ := regexp.Compile(needle)
	return r.ReplaceAllString(haystack, replacement)
}
