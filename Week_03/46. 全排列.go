package Week_03

func permute(nums []int) [][]int {
    result := [][]int{}
    visited := map[int]bool{}
	var dfs func(path []int)
    dfs = func(path []int) {
		if len(path) == len(nums) {
			result = append(result, path)
			return
		}
		for _, value := range nums {
			if setStatus, exists := visited[value]; exists && setStatus {
				continue
			}
            visited[value] = true
            dfs(append(path, value))
            visited[value] = false
        }
    }
    dfs([]int{})
	return result
}
