package arrays

/*
Problem Description

Given a array A of non negative integers, arrange them such that they form the largest number.

Note: The result may be very large, so you need to return a string instead of an integer.

Problem Constraints

1 <= len(A) <= 100000
0 <= A[i] <= 2*109

Input Format

First argument is an array of integers.

Output Format

Return a string representing the largest number.

Example Input

Input 1:

 A = [3, 30, 34, 5, 9]
Input 2:

 A = [2, 3, 9, 0]

Example Output

Output 1:

 "9534330"
Output 2:

 "9320"
Example Explanation

Explanation 1:

 A = [3, 30, 34, 5, 9]
 Reorder the numbers to [9, 5, 34, 3, 30] to form the largest number.
Explanation 2:

 Reorder the numbers to [9, 3, 2, 0] to form the largest number 9320.

*/

import (
	"sort"
	"strconv"
)

func largestNumber(A []int) string {

	//str:= strings.Trim(strings.Replace(fmt.Sprint(A), " ", "", -1), "[]")
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
	}
	if sum == 0 {
		return "0"
	}

	sort.Slice(A, func(i, j int) bool {
		if strconv.Itoa(A[i])+strconv.Itoa(A[j]) > strconv.Itoa(A[j])+strconv.Itoa(A[i]) {
			return true
		} else if strconv.Itoa(A[i])+strconv.Itoa(A[j]) < strconv.Itoa(A[j])+strconv.Itoa(A[i]) {
			return false
		} else {
			return true
		}
	})
	s := ""
	for i := 0; i < len(A); i++ {
		s = s + strconv.Itoa(A[i])
	}
	return s
}
