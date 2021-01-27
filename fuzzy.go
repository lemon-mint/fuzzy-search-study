package main

import "strings"

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
