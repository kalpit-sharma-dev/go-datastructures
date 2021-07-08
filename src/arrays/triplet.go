package arrays

import "fmt"

func Triplet() {

	fmt.Println("Triplet")
	//For this solution Array Must be Sorted
	targetSum := 15
	arr := []int{-4, -1, 1, 2, 3, 5, 7, 9, 10, 20}
	for i := 0; i < len(arr); i++ {
		j := i + 1
		k := len(arr) - 1

		for j < k {
			currSum := arr[i] + arr[j] + arr[k]
			if currSum == targetSum {
				fmt.Println(arr[i], arr[j], arr[k])
				break
			} else if targetSum > currSum {
				j++
			} else {
				k--
			}

		}
	}

}
