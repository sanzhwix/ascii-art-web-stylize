package main

import (
	"fmt"
	"os"

	printing "ascii-art/art"
	validators "ascii-art/validation"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("usage: go run ./cmd <string>")
		return
	}

	input := args[0]
	if !validators.BannerValidity() {
		fmt.Println("Banner file was changed")
		return
	}
	if !validators.AsciiCharValidation(input) {
		return
	}

	printing.Processing(input)
}
