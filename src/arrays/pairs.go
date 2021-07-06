package arrays

import "fmt"

func Pair() {
	fmt.Println("PAIR")
	s := 11
	arr := []int{7, 2, 10, 9, 3, -6, 1}
	fmt.Println(arr)
	// Brute Force Complexity O(n^2)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == s {
				fmt.Println(arr[j], " and ", arr[i])
			}
		}

	}

	//Optimize

}
