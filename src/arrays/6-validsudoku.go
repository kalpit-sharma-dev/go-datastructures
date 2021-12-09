package arrays

/**
 * @input A : array of strings
 *
 * @Output Integer
 */
func isValidSudoku(A []string) int {

	m := make(map[string]int)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if A[i][j] != '.' {
				if v, ok := m["row"+string(i)+string(A[i][j])]; ok {
					v += 1
					m["row"+string(i)+string(A[i][j])] = v
				} else {
					m["row"+string(i)+string(A[i][j])] = 1
				}
				if v, ok := m["col"+string(j)+string(A[i][j])]; ok {
					v += 1
					m["col"+string(j)+string(A[i][j])] = v
				} else {
					m["col"+string(j)+string(A[i][j])] = 1
				}
				if v, ok := m["box"+string((i/3)*3+(j/3))+string(A[i][j])]; ok {
					v += 1
					m["box"+string((i/3)*3+(j/3))+string(A[i][j])] = v
				} else {
					m["box"+string((i/3)*3+(j/3))+string(A[i][j])] = 1
				}

				if v, _ := m["row"+string(i)+string(A[i][j])]; v > 1 {
					return 0
				}
				if v, _ := m["col"+string(j)+string(A[i][j])]; v > 1 {
					return 0
				}
				if v, _ := m["box"+string((i/3)*3+(j/3))+string(A[i][j])]; v > 1 {
					return 0
				}

			}
		}
	}
	return 1
}
