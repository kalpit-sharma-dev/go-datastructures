package arrays

/*
Problem Description

You have given a string A having Uppercase English letters.

You have to find that how many times subsequence "AG" is there in the given string.

NOTE: Return the answer modulo 109 + 7 as the answer can be very large.



Problem Constraints

1 <= length(A) <= 105



Input Format

First and only argument is a string A.



Output Format

Return an integer denoting the answer.



Example Input

Input 1:

 A = "ABCGAG"
Input 2:

 A = "GAB"


Example Output

Output 1:

 3
Output 2:

 0

*/

func solveag(A string) int {

	countg := 0
	ans := 0
	for i := len(A) - 1; i >= 0; i-- {
		c := string(A[i])

		if c == "G" {
			countg++
		}
		if c == "A" {
			ans += countg
		}
	}
	return ans % 1000000007
}
