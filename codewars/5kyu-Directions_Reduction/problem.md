# 5kyu. Directions Reduction

**Difficulty:** 5 kyu  
**URL:** https://www.codewars.com/kata/550f22f4d758534c1100025a/

## Description

Write a function dirReduc which will take an array of strings and returns an array of strings with the needless directions removed (W<->E or S<->N side by side).

## Examples

```text
Input: ["NORTH", "SOUTH", "EAST", "WEST"]
Output: [] 

Input: ["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"]
Output: ["WEST"] 

Input: ["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"]
Output: ["WEST"] 

Input: ["NORTH", "WEST", "SOUTH", "EAST"]
Output: ["NORTH", "WEST", "SOUTH", "EAST"]

Input: ["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"]
Output: ["NORTH"]
```

## Note

Not all paths can be made simpler. 
The path ["NORTH", "WEST", "SOUTH", "EAST"] is not reducible.
"NORTH" and "WEST", "WEST" and "SOUTH", "SOUTH" and "EAST" are not directly opposite of each other and can't become such.
Hence the result path is itself: ["NORTH", "WEST", "SOUTH", "EAST"].

## Benchmark tests

For benchmark tests related to this task, please refer to the [bench.md](iterations/bench.md) file.
