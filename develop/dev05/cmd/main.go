package main

import (
	"grep"
)

func main() {
	err := grep.Grep()
	if err != nil {
		panic(err)
	}
}
