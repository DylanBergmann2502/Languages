package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	// multiline strings
	const names = `Joe
William
Jack`
	p(names)

	p("Contains:  	", s.Contains("test", "es"))
	p("Count:    	", s.Count("test", "t"))
	p("HasPrefix: 	", s.HasPrefix("test", "te"))
	p("HasSuffix: 	", s.HasSuffix("test", "st"))
	p("Index:     	", s.Index("test", "e"))
	p("Join:      	", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    	", s.Repeat("a", 5))
	p("Replace:   	", s.Replace("foo", "o", "0", -1))
	p("Replace:   	", s.Replace("foo", "o", "0", 1))
	p("ReplaceAll:	", s.ReplaceAll("foo", "o", "0"))
	p("Split:     	", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   	", s.ToLower("TEST"))
	p("ToUpper:   	", s.ToUpper("test"))
	p("TrimSpace:	", s.TrimSpace("    test       "))
}