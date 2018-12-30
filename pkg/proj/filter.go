package proj

import (
	"github.com/sahilm/fuzzy"
)

func FuzzyMatch(needle string, haystack Projects) Projects {
	if len(needle) <= 0 {
		return haystack
	}
	matchedProjects := NewProjects()
	matches := fuzzy.FindFrom(needle, haystack)
	for _, match := range matches {
		matchedProjects.Add(haystack.Get(match.Index))
	}
	return matchedProjects
}
