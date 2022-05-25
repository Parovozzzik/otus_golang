package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordsStruct struct {
	word  string
	count int
}

func Top10(text string) []string {
	text = prepareText(text)

	wordsMap := strings.Fields(text)
	wordsCountMap := map[string]int{}
	for _, word := range wordsMap {
		wordsCountMap[word]++
	}

	wordsSlice := make([]wordsStruct, 0, len(wordsCountMap))
	for k, v := range wordsCountMap {
		wordsSlice = append(wordsSlice, wordsStruct{k, v})
	}

	sort.Slice(wordsSlice, func(i, j int) bool {
		if wordsSlice[i].count == wordsSlice[j].count {
			return wordsSlice[i].word < wordsSlice[j].word
		}
		return wordsSlice[i].count > wordsSlice[j].count
	})

	return prepareResult(wordsSlice)
}

func prepareText(text string) string {
	text = strings.ToLower(text)
	replacer := strings.NewReplacer(
		" - ", " ",
		",", "",
		".", "",
	)
	text = replacer.Replace(text)

	return text
}

func prepareResult(wordsSlice []wordsStruct) []string {
	result := make([]string, 0, len(wordsSlice))
	var count int
	count = 10
	if len(wordsSlice) < 10 {
		count = len(wordsSlice)
	}

	for _, v := range wordsSlice[0:count] {
		result = append(result, v.word)
	}

	return result
}
