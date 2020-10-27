package leetcode

import (
	"log"
	"math"
	"strconv"
)

// Problem space:
// Given a 32-bit signed integer, reverse digits of an integer.

/*
Solution Space:

## SOLUTION ONE
- convert to string (array of bytes) (Itoa)
- reverse the string (reversing the array of bytes)
- convert string back to int, and return
- Return 0 if the integer is -/+ than Min/MaxInt32

## SOLUTION TWO
- Use % and / to create an array of digits
- Reverse the array: `reverse(list []interface{}) []interface{}`
- Use + and *10 to append digits into an int64
- Return 0 if the integer is -/+ than Min/MaxInt32

 */

func reverseInt(x int) int {
	isNegative := x < 0
	if isNegative {
		x = x * -1
	}

	reversedInt, err := strconv.Atoi(reverseString(strconv.Itoa(x)))
	if err != nil {
		log.Panic(err)
	}

	if isNegative {
		reversedInt = reversedInt * -1
	}

	if reversedInt < math.MinInt32 || reversedInt > math.MaxInt32 {
		return 0
	} else {
		return reversedInt
	}
}

func reverseString(s string) string {
	reversed := make([]rune, len(s))
	for i, r := range s {
		reversed[len(s) - i - 1] = r
	}
	return string(reversed)
}
