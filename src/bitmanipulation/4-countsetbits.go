package bitmanipulation

/* Problem Description

Write a function that takes an integer and returns the number of 1 bits it has.

Problem Constraints

1 <= A <= 109

Input Format

First and only argument contains integer A

Output Format

Return an integer as the answer

Example Input

Input1:
11
Example Output
Output1:
3
Example Explanation

Explaination1:
11 is represented as 1011 in binary.

*/

func numSetBits(A int) int {

	num := A
	c := 0
	//var bin []int
	rem := 0
	for num > 0 {
		rem = num % 2
		// bin=append(bin,rem)
		num = num / 2
		if rem == 1 {
			c++
		}
	}
	return c

}
