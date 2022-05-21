package main

import (
	"fmt"

	"softwareinstall-app/pkg"
)

func main() {
	s := pkg.NewRuleSet()
	s.AddDep("a", "b")
	s.AddDep("a", "c")
	s.AddConflict("b", "d")
	s.AddConflict("b", "e")
	if !s.IsCoherent() {
		fmt.Println("s.IsCoherent failed")
	}

	selected := pkg.New(s)
	selected.Toggle("d")
	selected.Toggle("e")
	selected.Toggle("a")
	fmt.Println(selected.StringSlice())
}
