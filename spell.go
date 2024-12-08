package spell

import (
	"strings"
)

var (
	oneToTen = []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ay",
		"bee",
		"see",
		"dee",
		"ee",
		"ef",
	}
	elevenToFleventeen = []string{
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
		"abteen",
		"bibteen",
		"cibteen",
		"dibbleteen",
		"ebbleteen",
		"fleventeen",
	}
	multipleOfSixteen = []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
		"atta",
		"bitta",
		"citta",
		"dickety",
		"eckity",
		"fleventy",
	}
	endingName = []string{
		"",
		"millby",
		"billby",
		"trillby",
		"quadrillby",
		"quintilby",
	}
)

// Number returns the spelling of the number.
func Number(x int) string {
	if x < 0 {
		return "negative " + Number(-x)
	}
	return LargeNumber(uint64(x))
}

// LargeNumber returns the spelling of the number.
func LargeNumber(x uint64) string {
	if x == 0 {
		return "zero"
	}
	var stack []string
	suffixIdx := 0
	for x > 0 {
		lowerWord := spellWord(x & 0xFFFF)
		if lowerWord == "" {
			suffixIdx++
			x >>= 16
			continue
		}

		suffix := endingName[suffixIdx]
		if suffix != "" {
			stack = append(stack, suffix)
		}
		stack = append(stack, lowerWord)

		suffixIdx++
		x >>= 16
	}
	reverse(stack)
	return strings.Join(stack, " ")
}

func spellWord(x uint64) string {
	lowerByte := spellByte(x & 0xFF)
	if x < 0x100 {
		return lowerByte
	}
	upperByte := spellByte(x >> 8)
	if x&0xFF == 0 {
		return upperByte + " bitey"
	}
	return strings.Join([]string{upperByte, "bitey", lowerByte}, " ")
}

func spellByte(x uint64) string {
	if x < 0x10 {
		return oneToTen[x]
	}
	if x < 0x20 {
		return elevenToFleventeen[x-0x10]
	}

	upperNibble := multipleOfSixteen[(x>>4)&0xF]
	if x&0xF == 0 {
		return upperNibble
	}
	lowerNibble := oneToTen[x&0xF]
	return upperNibble + " " + lowerNibble
}

func reverse(stack []string) {
	for i := 0; i < len(stack)/2; i++ {
		j := len(stack) - 1 - i
		stack[i], stack[j] = stack[j], stack[i]
	}
}
