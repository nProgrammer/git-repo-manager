package main

//TODO: Lista commitów

import (
	"git-manager/views"
)

func main() {
	err := views.CommitsView()
	if err != nil {
		panic(err)
	}
}
