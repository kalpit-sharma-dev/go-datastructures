package arrays

/*
Problem Description
Given an array of size n, find the majority element. The majority element is the element that appears more than floor(n/2) times.
You may assume that the array is non-empty and the majority element always exist in the array.

Example :

Input : [2, 1, 2]
Return  : 2 which occurs 2 times which is greater than 3/2.

*/

func majorityElement(A []int) int {

	maj := 0
	freq := 0
	for i := 0; i < len(A); i++ {
		if maj == 0 {
			maj = A[i]
			freq = 1
		} else if maj == A[i] {
			freq += 1
		} else {
			freq -= 1
		}
		if freq == 0 {
			maj = 0
		}
	}
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == maj {
			count++
		}
	}
	if count >= (len(A)/2)+1 {
		return maj
	}
	return 0
}
