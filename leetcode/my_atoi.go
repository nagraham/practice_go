package leetcode

import (
	"math"
	"strings"
)

/*
PROBLEM:
Convert string to integer. Discard whitespace until first non-whitespace rune found.
Take optional +/- character. Convert digits into an integer. If non-digit is encountered,
the integer is found, return it. If no integer found, return zero. If the given digit is
less than Math.MinInt32, return the min int; if greater than MaxInt32, return max int.

SOLUTION:
Iterate string (not a loop, maintain a pointer)

Step 1:
Discard white space (' ') runes

Step 2:
First non-white space
  - If +/-, handle (store sign)
  - If digit, convert
  - If non-space, return zero

Step 3:
While each rune is digit (if non-digit or end of string, break)
  - convert

Step 4:
Apply the sign to the result

Step 5:
Handle integer overflow
 */

type Sign string
const (
	Positive Sign = "positive"
	Negative Sign = "negative"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	s, sign := parseSignAndTrim(s)
	result := convertNum(s, sign)
	return keepWithin(result, math.MinInt32, math.MaxInt32)
}

func parseSignAndTrim(s string) (string, Sign){
	if len(s) == 0 {
		return s, Positive
	} else if s[0] == '-' {
		return s[1:], Negative
	} else if s[0] == '+' {
		return s[1:], Positive
	} else {
		return s, Positive
	}
}

func convertNum(s string, sign Sign) int {
	result := 0
	numSize := 0
	for _, r := range s {
		isNumBeyondInt32 := numSize > 10
		if !isNumber(r) || isNumBeyondInt32 {
			break
		}
		if result < 0 || result > 0 {
			result *= 10
			numSize++
		}
		result += int(r - '0')
	}

	if sign == Negative {
		result *= -1
	}

	return result
}


func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

// return the max or min if the given number is beyond those bounds;
// otherwise, return the number unchanged.
func keepWithin(i int, min int, max int) int {
	if i >= max {
		return max
	} else if i <= min {
		return min
	} else {
		return i
	}
}