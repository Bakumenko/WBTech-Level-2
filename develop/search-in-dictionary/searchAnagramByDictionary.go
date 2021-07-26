package main

import (
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SearchAnagramByDictionary(slice *[]string) *map[string][]string {
	result := map[string][]string{}
	mapStringToSliceRunes := map[string][]rune{}
	for _, s := range *slice {
		s = strings.ToLower(s)
		var stringInRunes sortRunes = []rune(s)
		sort.Sort(stringInRunes)
		mapStringToSliceRunesIsNeededToExpanded := true
		for key, value := range mapStringToSliceRunes {
			if string(value) == string(stringInRunes) {
				if _, ok := result[key]; ok {
					result[key] = append(result[key], s)
				} else {
					result[key] = []string{s}
				}
				mapStringToSliceRunesIsNeededToExpanded = false
			}
		}
		if mapStringToSliceRunesIsNeededToExpanded {
			mapStringToSliceRunes[s] = stringInRunes
		}
	}
	return &result
}
