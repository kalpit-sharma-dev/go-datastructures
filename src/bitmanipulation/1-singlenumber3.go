package bitmanipulation

/*
Single Number III

Problem Description

Given an array of numbers A , in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

Note: Output array must be sorted.

Problem Constraints

2 <= |A| <= 100000
1 <= A[i] <= 109

Input Format

First argument is an array of interger of size N.

Output Format

Return an array of two integers that appear only once.

Example Input

Input 1:

A = [1, 2, 3, 1, 2, 4]
Input 2:

A = [1, 2]

Example Output

Output 1:

[3, 4]
Output 2:

[1, 2]

Example Explanation

Explanation 1:

 3 and 4 appear only once.
Explanation 2:

 1 and 2 appear only once.

*/

func solve1(A []int) []int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = sum ^ A[i]
	}

	rsbm := sum & (-sum)
	a1 := 0
	a2 := 0

	for i := 0; i < len(A); i++ {
		if (A[i] & rsbm) > 0 {
			a1 = a1 ^ A[i]
		} else {
			a2 = a2 ^ A[i]
		}
	}
	M := make([]int, 0)
	if a1 < a2 {
		M = append(M, a1)
		M = append(M, a2)
		return M
	}
	M = append(M, a2)
	M = append(M, a1)
	return M
}
