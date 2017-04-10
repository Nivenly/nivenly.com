package main

import (
	"sort"
	"fmt"
	"unicode/utf8"
)

func main() {
	types := getUnsortedTypes()
	sort.Sort(types)
	fmt.Println(types)
}

type Type struct {
	Value string
}

type Types []Type

func (s Types) Len() int {
	return len(s)
}
func (s Types) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Types) Less(i, j int) bool {
	iRune, _ := utf8.DecodeRuneInString(s[i].Value)
	jRune, _ := utf8.DecodeRuneInString(s[j].Value)
	return int32(iRune) < int32(jRune)
}

func getUnsortedTypes() Types {
	return []Type{
		{
			Value: "the",
		},
		{
			Value: "golf",
		},
		{
			Value: "trip",
		},
		{
			Value: "sounds",
		},
		{
			Value: "like",
		},
		{
			Value: "fun",
		},
	}
}
