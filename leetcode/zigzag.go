package leetcode

import "strings"

/*
Run a "Zig-Zag" conversion:

PAYPALISHIRING

3 ->
P   A   H   N
A P L S I I G
Y   I   R

Read by line: PAHNAPLSIIGYIR

4 ->
P     I    N
A   L S  I G
Y A   H R
P     I

Read by line: PINALSIGYAHRPI

string.length up to 1000
Any lower/upper case English letters
Up to 1000 rows
 */

/*
## APPROACH ONE:
Create a grid of runes. Iterate string; for each character, add to the grid
in the zig zag pattern.

Then, iterate the grid to get each character.

How to size the grid? (maybe append new columns?)

## APPROACH TWO
Create an list of lists (possibly a linked list, we don't need random access, but we need dynamic sizing + iteration)
Iterate string.
Go across the top-level array, appending characters to lists. When you reach the end, reverse and go backwards.
Bounce between the end of the array and the front.

### Decomposition

Create an "Oscillator"
- Oscillates between two int64 values
- You can set the values, you can set the rate of oscillation
- There is more you could do with this, but it's basic settings should give us what we need

```
type Oscillator interface {
  Move()
  Get() int64
}
func NewOscillator(start int, end int, rate int) Oscillator
// in our solution, we'll call `NewOscillator(0, numRows, 1)`
```

The list of lists already composes modules.

To create the final result:
sb = create a string builder
for each list
  for each rune in list
    add rune to sb
 */

func zigzag(s string, numRows int) string {
	allLists := make([][]rune, numRows)
	oscillator := NewZigZagOscillator(0, numRows)

	// Oscillate through "allLists," appending runes to each list
	for _, r := range s {
		list := allLists[oscillator.Get()]
		allLists[oscillator.Get()] = append(list, r)
		oscillator.Move()
	}

	// For each list, add all the runes to a string builder
	var sb strings.Builder
	sb.Grow(len(s))
	for _, list := range allLists {
		for _, r := range list {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// Controls movement along an Oscillation.
type Oscillator interface {

	// Increment the Oscillation by 1 unit
	Move()

	// Get the current value
	Get() int
}

// Creates an Oscillator that forms a "ZigZag" (i.e. linear, sawtooth) oscillation pattern
func NewZigZagOscillator(start int, end int) Oscillator {
	// PreCondition: end > start
	return &zigZagger{
		start: start,
		end: end,
		curr: start,
		isReverse: false,
	}
}

type zigZagger struct {
	start int
	end int
	curr int
	isReverse bool
}

func (zz *zigZagger) Move() {
	if zz.isReverse {
		zz.curr = zz.curr - 1
	} else if zz.curr +1 != zz.end {
		zz.curr = zz.curr + 1
	}

	if zz.curr == zz.start {
		zz.isReverse = false
	} else if zz.curr + 1 == zz.end {
		zz.isReverse = true
	}
}

func (zz *zigZagger) Get() int {
	return zz.curr
}

