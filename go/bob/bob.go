// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)



// Hey should have a comment documenting it.
func Hey(remark string) string {

	strings.ContainsAny(remark, "?")

	switch remark {



	case "question?":
		return "Sure"
	case "Yell (uppercase)":
		return "Whoa, chill out!"
	case "yell question? (uppercase with?)":
		return "Calm down, I know what I'm doing!"
	case "":
		return "Fine. Be that way!"
	default:
		return "Whatever."


	}
}
