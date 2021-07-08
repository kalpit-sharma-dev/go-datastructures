package arrays

import "fmt"

func Triplet() {

	fmt.Println("Triplet")
	//For this solution Array Must be Sorted
	s := 15
	arr := []int{1, 2, 3, 5, 7, 9, 10}
	for i := 0; i < len(arr); i++ {
		j := i + 1
		k := len(arr) - 1
		find := s - arr[i]
		for j < k {
			if arr[j]+arr[k] == find {
				fmt.Println(arr[i], arr[j], arr[k])
				break
			} else if find > arr[j]+arr[k] {
				j++
			} else {
				k--
			}

		}
	}

}
