package main

import (
	"math"
	"sort"
	"strconv"
)

func AddStringToSet(str string, mapAsSet map[string]struct{}) {
	if _, inSet := mapAsSet[str]; !inSet {
		mapAsSet[str] = struct{}{}
	}
}

func GetSortedKeys(m map[string]struct{}) []string {
	sortedKeys := make([]string, len(m))
	i := 0
	for k, _ := range m {
		sortedKeys[i] = k
		i++
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

// totally stole this from https://gist.github.com/DavidVaini/10308388#comment-1391788
func RoundToPrecision(f float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Floor((f*shift)+.5) / shift
}

func PrettyFormatFloat(f float64) string {
	rounded := RoundToPrecision(f, 3)
	return strconv.FormatFloat(rounded, 'f', -1, 64)
}
