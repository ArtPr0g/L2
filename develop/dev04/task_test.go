package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	tests := []struct {
		name   string
		array  []string
		answer map[string][]string
	}{
		{name: "with one elem", array: []string{"a"}, answer: map[string][]string{}},
		{name: "test reg", array: []string{"ПрИвЕт", "ТеВиРп"}, answer: map[string][]string{
			"привет": getSorted("привет", "тевирп"),
		}},
		{name: "test sort", array: []string{"abc", "acb", "bca", "bac", "cab", "cba"}, answer: map[string][]string{
			"abc": getSorted("abc", "acb", "bca", "bac", "cab", "cba"),
		}},
		{name: "test several elems", array: []string{"abc", "acb", "bca", "bac", "cab", "cba", "ba", "AB"}, answer: map[string][]string{
			"abc": getSorted("abc", "acb", "bca", "bac", "cab", "cba"),
			"ba":  getSorted("ba", "ab"),
		}},
		{name: "test several equal elements", array: []string{"abc", "acb", "acb"}, answer: map[string][]string{
			"abc": getSorted("abc", "acb"),
		}},
		{name: "empty", array: []string{}, answer: map[string][]string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			set := findAnagramSets(test.array)
			ok := reflect.DeepEqual(set, test.answer)
			if !ok {
				t.Fatalf("excepted: %v, actual: %v", test.answer, set)
			}
		})
	}
}

func getSorted(elems ...string) []string {
	sort.Strings(elems)
	return elems
}
