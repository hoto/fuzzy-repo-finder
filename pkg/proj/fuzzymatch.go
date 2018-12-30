package proj

import (
	"github.com/sahilm/fuzzy"
)

func FuzzyMatch(needle string, haystack Projects) Projects {
	if len(needle) == 0 {
		return haystack
	}
	matchingProjects := NewProjects()
	matches := fuzzy.FindFrom(needle, haystack)
	for _, match := range matches {
		matchingProjects.Add(haystack.Get(match.Index))
	}
	return matchingProjects
}
