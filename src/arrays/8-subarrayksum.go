package arrays

/**
 * @input A : Integer array
 * @input B : Integer
 *
 Problem Description

Given an array of positive integers A and an integer B, find and return first continuous subarray which adds to B.

If the answer does not exist return an array with a single element "-1".

First sub-array means the sub-array for which starting index in minimum.

Problem Constraints

1 <= length of the array <= 100000
1 <= A[i] <= 109
1 <= B <= 109

Input Format

The first argument given is the integer array A.

The second argument given is integer B.

Output Format

Return the first continuous sub-array which adds to B and if the answer does not exist return an array with a single element "-1".

Example Input

Input 1:

 A = [1, 2, 3, 4, 5]
 B = 5
Input 2:

 A = [5, 10, 20, 100, 105]
 B = 110

Example Output

Output 1:

 [2, 3]
Output 2:

 -1

Example Explanation

Explanation 1:

 [2, 3] sums up to 5.
Explanation 2:

 No subarray sums up to required number.
 * @Output Integer array.
*/
func subarrayksum(A []int, B int) []int {
	flag := false
	pos := 0
	end := 0
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	ps := make([]int, len(A))
	ps[0] = A[0]
	for i := 1; i < len(A); i++ {
		ps[i] = A[i] + ps[i-1]
		if ps[i] == B {
			return A[0 : i+1]
		}
	}
	for i := 0; i < len(A); i++ {
		sum += A[i]
		find := sum - B
		end = i
		if _, ok := m[find]; ok {
			for p := 0; p <= i; p++ {
				if ps[p] == find {
					pos = p
					flag = true
				}
			}
			break
		}
		m[sum] = 1
	}
	if flag {
		out := make([]int, end-pos)
		t := 0
		for i := pos + 1; i <= end; i++ {
			out[t] = A[i]
			t++
		}
		return out
	}
	out := make([]int, 1)
	out[0] = -1
	return out
}
