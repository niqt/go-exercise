package main

import (
	"fmt"
	"sort"
)

func main() {
	cities := []string{
		"New York", "London", "Tokyo", "Paris", "Sydney", "Berlin", "Moscow", "Rome", "New York", "New York",
		"Tokyo", "Paris", "Sydney", "Tokyo",
	}
	fmt.Println(fiveOccurence(cities))
}

func fiveOccurence(slice []string) []string {
	occurrenceMap := make(map[string]int, 5)

	for _, elem := range slice {
		occurrenceMap[elem]++
	}
	keys := make([]string, 0, len(occurrenceMap))

	for key := range occurrenceMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if occurrenceMap[keys[i]] > occurrenceMap[keys[j]] {
			return true
		}
		if occurrenceMap[keys[i]] == occurrenceMap[keys[j]] && keys[i] > keys[j] {
			return true
		}
		return false
	})

	best := []string{}
	for i, k := range keys {
		if i < 5 {
			best = append(best, k)
		} else {
			break
		}
	}
	return best
}
