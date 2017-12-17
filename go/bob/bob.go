// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"bytes"
	"fmt"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	fmt.Println("The last char is: ", string(remark[len(remark)-1]))
	arr := make([]byte, len(remark))
	copy(arr[:], remark)

	if isQuestion(remark) {
		arr = bytes.Trim(arr, "?")

		if isUppercase(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure"

	}

	if isUppercase(remark) {
		return "Whoa, chill out!"
	} else {

	}

	switch remark {

	case "":
		return "Fine. Be that way!"
	default:
		return "Whatever."

	}
}

func isQuestion(remark string) (question bool) {

	lastSymbol := string(remark[len(remark)-1])
	if lastSymbol == "?" {

		return true

	}
	return false

}

func isUppercase(remark string) (allUppercase bool) {

	return true
}
