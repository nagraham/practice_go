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