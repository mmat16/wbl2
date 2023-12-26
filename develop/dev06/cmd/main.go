package main

import (
	"cut"
)

func main() {
	err := cut.Cut()
	if err != nil {
		panic(err)
	}
}
