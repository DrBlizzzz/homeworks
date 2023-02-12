package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

func Top10(text string) []string {
	wordsMap := make(map[string]int)
	re := regexp.MustCompile(`\S+`)
	matches := re.FindAllString(text, -1)
	for _, match := range matches {
		if _, valExist := wordsMap[match]; !valExist {
			wordsMap[match] = 0
		}
		wordsMap[match]++
	}
	wordsMapKeys := make([]string, 0)
	for word := range wordsMap {
		wordsMapKeys = append(wordsMapKeys, word)
	}
	sort.Slice(wordsMapKeys, func(i, j int) bool {
		switch {
		case wordsMap[wordsMapKeys[i]] > wordsMap[wordsMapKeys[j]]:
			return true
		case wordsMap[wordsMapKeys[i]] == wordsMap[wordsMapKeys[j]]:
			return wordsMapKeys[i] < wordsMapKeys[j]
		default:
			return false
		}
	})
	if len(wordsMap) <= 10 {
		return wordsMapKeys[:len(wordsMap)]
	}
	return wordsMapKeys[:10]
}
