package leetcode

import "math"

/*

# Problem
Given n non-negative integers a1, a2, ..., an , where each represents a point
at coordinate (i, ai). n vertical lines are drawn such that the two endpoints
of the line i is at (i, ai) and (i, 0). Find two lines, which, together with
the x-axis forms a container, such that the container contains the most water.

# Brute Force:
- (A) Iterate through possible combinations of areas
  - 2x nested loops
  - Compare all values

- (B) Calculate the area with the 2 values
  - The height of the water container is min(a, b) of the two values
  - The width of the water container is the difference between their two indices in the array

- (C) Compare the area to the cached max area, and replace if greater
  - Dead simple max compare

# Optimize:
- (B) and (C) are simple O(1) functions, so let's set them aside for now.
- (A) is the O(n^2) piece, so that's where we'll get our gain.
  - Can we use the properties of the problem to shake out an iteration?
- I feel any solution is going to require

## Idea One:
- Store an array of Point coordinates (x, y) where the x is the index, and y is the value
- Sort the array by the "y" value and/or the distance of the x value from the midpoint (favor locations at the edge)
- Nested loop O(n^2) loop on these values, with higher probability of quickly finding the max value
  - This is the problem ... we need to find the max, you can only find the max if you compare all values.

(REJECTED)


## Idea Two:

- For any value you are comparing, the container can be no taller
  - for instance, for a value of 1, despite higher numbers, the container height must be one
  - The only other thing to do is compare the distance with the furthest element that is at least one high
  - Once you have compared that furthest element, there's no point comparing others.

- There are a set of rules we can use to terminate the loops. It's O(n^2) / nested, but optimized
  - The outer loop iterates [0 -> len]; the inner loop iterates [len-1 -> curr]
  - if the innerloop value is >= the outerloop, you know you will not find any better value and you can iterate outerloop

SOLUTION:
Runtime: 2396 ms, faster than 5.05% of Go online submissions for Container With Most Water.
Hmm, I solved in just about an hour, but I'm disappointed with the performance


## Idea Three
- For each value, calculate it's highest potential
  - Height * distance to left or right, based on where it is in the array
- Have list sorted by highest potentials, and compare them in O(n^2) fashion
- Cache the max
- Stop iterating list when max > potential (you know the rest of the list wont make the cut)


## Idea Four
The idea here is that one height is only as good as having another height further away. Once you find a tall height,
it's a pretty good candidate for contributing to the max area. So keep it, and look for a complimentary tall height
on the other end. Move towards the middle, switching for higher and higher values as you walk the list.

Two pointers, one at either end. Compare area. Move towards middle. Walk the pointer that has the lowest value.

SOLUTION:
Runtime: 16 ms, faster than 92.77% of Go online submissions for Container With Most Water.

*/

func MaxWaterContainer(heights []int) int {
	max := 0

	i := 0
	j := len(heights)-1
	for i < j {
		area := calculateWaterArea(heights, i, j)
		max = MaxInt(max, area)

		if heights[i] >= heights[j] {
			j--
		} else {
			i++
		}
	}

	return max
}

func MaxWaterContainerOld(heights []int) int {
	max := 0

	for i := 0; i < len(heights); i++ {
		a := heights[i]

		// this is the best potential value the current height could achieve;
		// if less than the current max, there is no point wasting our time
		if a * (len(heights) - i) < max {
			continue
		}

		for j := len(heights)-1; j > i; j-- {
			b := heights[j]
			area := calculateWaterArea(heights, i, j)
			max = MaxInt(max, area)

			// if we have found a value as high or higher than a, we can't find any better
			if b >= a {
				break // the inner loop
			}
		}
	}

	return max
}


// Finds the area of a container based on the height values located at index i and j
// Prerequisites: all values in heights should be >= 0; i should be < j
func calculateWaterArea(heights []int, i int, j int) int {
	if j <= i {
		panic("Invalid values!")
	}

	a := heights[i]
	b := heights[j]

	min := math.Min(float64(a), float64(b))

	return int(min) * (j - i)
}
