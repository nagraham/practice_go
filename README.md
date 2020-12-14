# Practice Go

This directory contains different modules encapsulating little practice projects.

## Leet Code

**Problem 1**: [Two Sum](https://leetcode.com/problems/two-sum/) -> leetcode/twosum.go

**Problem 2**: [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) -> leetcode/sumlinkedlist.go

**Problem 3**: [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) -> leetcode/substr_wo_repeats.go

**Problem 4**: [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) -> leetcode/sorted_array_median.go

**Problem 5**: [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) -> leetcode/longest_palindrome.go

**Problem 6**: [Zig Zag Conversion](https://leetcode.com/problems/zigzag-conversion/) -> leetcode/zigzag.go
- I wrote up my answer in a [post](https://leetcode.com/problems/zigzag-conversion/discuss/906017/GoLang-Modular-Solution-8ms)
- I like my introduction:
- > I like to decompose solutions into modular, othogonal, possibly reusable components. At the very least, it allows me to abstract away difficult pieces of the solution space, and focus on one piece at a time. I optimize for this over pure speed.

**Problem 7**: [Reverse Integer](https://leetcode.com/problems/reverse-integer/) -> leetcode/reverse.go

**Problem 8**: [String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/) -> leetcode/my_atoi.go
- Wrote up an [post](https://leetcode.com/problems/string-to-integer-atoi/discuss/911913/Golang-Clean-modular-solution-0ms).
- > I try to decompose these problems down into independent subproblems, and enclose each of those problems in a small module -- a function or object. Ideally, the module should have a somewhat generic interface, decoupling it from the problem at hand. This makes it easier to understand, easier to test, easier to change, and possibly reusable.

**Problem 9**: [Number Palindrome](https://leetcode.com/problems/palindrome-number/) -> leetcode/int_palindrome.go

**Problem 11**: [Container with Most Water](https://leetcode.com/problems/container-with-most-water/) -> leetcode/water_container.go
- I decomposed this into two functions: the main solution (containing the iteration) and a function to find the area from the array and ints. This helped me to focus on the part that needed to be optimized: the iteration
- The first solution was O(n^2) and complicated (more rules based). It didn't solve the quadratic nature of the problem.
- A few hours later, it occurred to me to ask: "how to do this in one loop?" And then I realized that because any value needed a lot of length or else another "high" value, you could walk the list with two pointers, keeping in place the index with the highest value. Store the max area as you go.
- Strategy for interviews:
    - Obviously there will be a brilliant, order of magnitiude solution based on hte nature of the problem (e.g. "the trick"). But don't go for gold at first.
    - First, *decompose the problem*, break into general pieces, make basic abstractions, write out the "master control" algorithm. This will win big points on its own and show confidence. 
    - Code the modules/functions that are straightforward. Get code on the board, so everyone feels progress has been made.
    - Encapsulate the O(n^2) part of the solution. Even write it, or pseudocode it to prove you can do it. Then iterate and ask, "how to make this an order of magnitude better?"
     