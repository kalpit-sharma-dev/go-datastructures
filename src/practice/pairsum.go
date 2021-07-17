package practice

import "fmt"

func PairSorted() {
	arr := []int{7, 2, 10, 9, 3, -6, 1}
	targetSum := 19
	i := 0
	j := len(arr) - 1
	//fmt.Println(j)
	for i < j {
		if arr[i]+arr[j] == targetSum {
			fmt.Println(arr[i], arr[j])
		} else if targetSum > arr[i]+arr[j] {
			i++
		} else {
			j--
		}
	}
}

func Pair() {
	arr := []int{7, 2, 10, 9, 3, -6, 1}
	targetSum := 19
	//i := 0
	m := make(map[int]int)

	for i := 1; i < len(arr); i++ {
		find := targetSum - arr[i]
		_, ok := m[find]
		if ok {
			fmt.Println(find, arr[i])
		} else {
			m[arr[i]] = arr[i]
		}
	}

}
