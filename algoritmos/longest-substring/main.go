package main

import (
	"fmt"
	"strconv"
	"strings"
)

func notInSequence(stringList []string, letter string) (bool, int) {
	for index, l := range stringList {
		if l == letter {
			return false, index

		}
	}
	return true, -1
}

func lengthOfLongestSubstring(s string) int {
	unique_letter := make(map[string][]string)
	substring := 1
	for _, l := range s {
		t, index := notInSequence(unique_letter[strconv.Itoa(substring)], string(l))

		if t {
			unique_letter[strconv.Itoa(substring)] = append(unique_letter[strconv.Itoa(substring)], string(l))
		} else {

			actual := unique_letter[strconv.Itoa(substring)]
			substring += 1
			if string(l) != actual[len(actual)-1] {
				unique_letter[strconv.Itoa(substring)] = append(unique_letter[strconv.Itoa(substring)], actual[index+1:]...)
			}
			unique_letter[strconv.Itoa(substring)] = append(unique_letter[strconv.Itoa(substring)], string(l))

		}

	}
	var biggest int

	for _, i := range unique_letter {

		if len(i) > biggest {
			biggest = len(i)
		}

	}

	return biggest
}

func main() {
	const base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ "
	minhaString := strings.Repeat(base, 25) + "abcdefghijklmnopqrstuvwxyzABCD"
	fmt.Println(lengthOfLongestSubstring(minhaString))
}
