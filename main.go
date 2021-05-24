package main

import (
	"os"
)

func main() {
	_, err := os.GetWd()
	if err != nil {
		panic("Failed to search work directory")
	}
}
