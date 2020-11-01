package main

import (
	"flag"
	"fmt"
	"strings"
)

type language struct {
	lowerCase []rune
	upperCase []rune
}

var (
	english = language{
		[]rune("abcdefghijklmnopqrstuvwxyz"),
		[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	}
	swedish = language{
		[]rune("abcdefghijklmnopqrstuvwxyzåäö"),
		[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ"),
	}
	german = language{
		[]rune("abcdefghijklmnopqrstuvwxyzäöüß"),
		[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÜẞ"),
	}
)

func index(sequence []rune, searchFor rune) int {
	for ix, s := range sequence {
		if s == searchFor {
			return ix
		}
	}
	return -1
}

func shiftRune(from int, shift int, sequence []rune) rune {
	return sequence[(from+shift)%len(sequence)]
}

func rotateRune(r rune, shift int, lang language) rune {
	lowerCaseRuneIndex := index(lang.lowerCase, r)
	if lowerCaseRuneIndex > -1 {
		return shiftRune(lowerCaseRuneIndex, shift, lang.lowerCase)
	}

	upperCaseRuneIndex := index(lang.upperCase, r)
	if upperCaseRuneIndex > -1 {
		return shiftRune(upperCaseRuneIndex, shift, lang.upperCase)
	}

	return r
}

func ceasarCipher(input string, shift int, lang language) string {
	builder := strings.Builder{}
	for _, r := range input {
		builder.WriteRune(rotateRune(r, shift, lang))
	}
	return builder.String()
}

func main() {
	shift := flag.Int("shift", 13, "The amount of characters to shift")
	debug := flag.Bool("debug", false, "To enable debug output")
	lang := flag.String("language", "en", "The language the decrypted text is in. Accepted values are en, de, sv")
	flag.Parse()

	var selectedLang language
	switch *lang {
	case "en":
		selectedLang = english
	case "de":
		selectedLang = german
	case "sv":
		selectedLang = swedish
	default:
		fmt.Printf("Unknown language flag: %v\n", *lang)
		return
	}

	if *shift < 0 {
		*shift = len(selectedLang.lowerCase) + *shift
	}

	input := strings.Join(flag.Args(), " ")
	output := ceasarCipher(input, *shift, selectedLang)

	if *debug {
		fmt.Printf("Input: %v, Output: %v\n", input, output)
		fmt.Printf("Tail: %v\n", flag.Args())
		fmt.Printf("Shift: %v\n", *shift)
	}
	fmt.Printf("%v\n", output)
}
