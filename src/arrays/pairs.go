package arrays

import "fmt"

func Pair() {
	fmt.Println("PAIR")
	s := 4
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
	fmt.Println("################# 7, 2, 10, 9, 3, -6, 1 ###############")
	fmt.Println("############### 7, 2, 10, 9, 3, -6, 1 #################")
	//Optimize arr[i]+x=s
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		find := s - arr[i]
		fmt.Println("find = ", find)
		_, ok := m[find]
		fmt.Println("ok = ", ok)
		if ok {
			fmt.Println(arr[i], " and ", find)
		} else {
			m[arr[i]] = arr[i]
			fmt.Println(m)
		}

	}
}
