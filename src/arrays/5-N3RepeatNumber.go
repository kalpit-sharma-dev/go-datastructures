package arrays

/*
Problem Description

You're given a read only array of n integers. Find out if any integer occurs more than n/3 times in the array in linear time and constant additional space.
If so, return the integer. If not, return -1.

If there are multiple solutions, return any one.

Example:

Input: [1 2 3 1 1]
Output: 1
1 occurs 3 times which is more than 5/3 times.

*/
func majorityElement3(A []int) int {
	me1 := 0
	f1 := 0
	me2 := 0
	f2 := 0
	for i := 0; i < len(A); i++ {
		if me1 == A[i] {
			f1 += 1
		} else if me2 == A[i] {
			f2 += 1
		} else if f1 == 0 {
			me1 = A[i]
			f1++
		} else if f2 == 0 {
			me2 = A[i]
			f2++
		} else {
			f1--
			f2--
		}
	}
	f1 = 0
	f2 = 0

	for i := 0; i < len(A); i++ {
		if me1 == A[i] {
			f1++
		}
		if me2 == A[i] {
			f2++
		}

	}
	if f1 > (len(A) / 3) {
		return me1
	}
	if f2 > (len(A) / 3) {
		return me2
	}
	return -1
}
