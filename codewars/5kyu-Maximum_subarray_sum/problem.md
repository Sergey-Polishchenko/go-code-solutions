# 5kyu. Maximum subarray sum

**Difficulty:** 5 kyu  
**URL:** https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go

## Description

The maximum sum subarray problem consists in finding the maximum sum of a contiguous subsequence in an array or list of integers:

Easy case is when the list is made up of only positive numbers and the maximum sum is the sum of the whole array.
If the list is made up of only negative numbers, return 0 instead.

Empty list is considered to have zero greatest sum. Note that the empty list or array is also a valid sublist/subarray.

## Examples

```text
Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6
// [4, -1, 2, 1]

Input: []
Output: 0
```

## Benchmark tests

For benchmark tests related to this task, please refer to the [bench.md](iterations/bench.md) file.
