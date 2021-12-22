package arrays

/*
Given an unsorted integer array A of size N.

Find the length of the longest set of consecutive elements from the array A.
Example Input

Input 1:

A = [100, 4, 200, 1, 3, 2]
Input 2:

A = [2, 1]


Example Output

Output 1:

 4
Output 2:

 2

*/

func longestConsecutive(nums []int) int {

	ans := 1
	max := -100000
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]-1]; !ok {

			ans = 0
			p, ok := m[nums[i]]
			for ok {
				ans++
				p++
				_, ok = m[p]
			}
			if ans > max {
				max = ans
			}
		}
	}
	return max

}
