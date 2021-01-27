package main

import (
	"sort"
	"strings"
)

func match(src, query string) bool {
	if len(query) > len(src) {
		return false
	}
	cur := 0
	for i := range query {
		IsCharFound := false
		for j := cur; j < len(src); j++ {
			if src[j] == query[i] {
				cur = j
				IsCharFound = true
				break
			}
		}
		if !IsCharFound {
			return false
		}
	}

	return true
}

func score(src, query string) int {
	matchScore := 0
	if len(query) > len(src) {
		return matchScore
	}
	cur := 0
	lastword := ""
	wordScore := 0
	for i := range query {
		for j := cur; j < len(src); j++ {
			if src[j] == query[i] {
				if j-cur == 1 {
					lastword += src[j : j+1]
					wordScore += len(lastword) * 2
				} else {
					lastword = src[j : j+1]
					matchScore += wordScore
					wordScore = 0
				}
				cur = j
				matchScore += 10
				break
			} else if strings.ToLower(src[j:j+1]) == strings.ToLower(query[i:i+1]) {
				if j-cur == 1 {
					lastword += src[j : j+1]
					wordScore += len(lastword) * 1
				} else {
					lastword = src[j : j+1]
					matchScore += wordScore
					wordScore = 0
				}
				cur = j
				matchScore += 5
				break
			}
		}
	}

	return matchScore
}

type fuzzySorter struct {
	data []string
	key  string
}

func (r fuzzySorter) Len() int {
	return len(r.data)
}

func (r fuzzySorter) Less(i, j int) bool {
	return score(r.data[i], r.key) < score(r.data[j], r.key)
}

func (r fuzzySorter) Swap(i, j int) {
	r.data[i], r.data[j] = r.data[j], r.data[i]
}

func sortSlice(data []string, key string) []string {
	sort.Sort(fuzzySorter{data: data, key: key})
	return data
}
