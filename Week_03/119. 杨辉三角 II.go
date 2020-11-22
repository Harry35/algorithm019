package Week_03

func getRow(rowIndex int) []int {
    rows := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		rows[i] = 1
		for j := i; j > 1; j-- {
			rows[j-1] = rows[j-1] + rows[j-2]
		}
	}
	return rows
}
