package bitmanipulation

// SMALLEST XOR
/*
Problem Constraints
0 <= A <= 109
0 <= B <= 30

Input Format

First argument contains a single integer A. Second argument contains a single integer B.

Output Format
Return a single integer X.
Example Input

Input 1:

 A = 3
 B = 3
Input 2:

 A = 15
 B = 2
Example Output
Output 1:
 7
Output 2:
 12

Example Explanation
Explanation 1:
 3 xor 7 = 4 which is minimum
Explanation 2:
 15 xor 12 = 3 which is minimum   */

func Solve(A int, B int) int {

	x := 0
	for i := 30; i >= 0; i-- {
		mask := 1 << uint(i)
		if A&mask > 0 && B > 0 {
			p := 1 << uint(i)
			x = x | p
			B--
		}
	}
	for i := 0; i <= 30; i++ {
		mask := 1 << uint(i)
		if (x&mask) == 0 && B > 0 {
			p := 1 << uint(i)
			x = x | p
			B--
		}
	}
	return x
}
