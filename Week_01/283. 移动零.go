package Week_01

func moveZeroes(nums []int)  {
    i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
		i++
	}
}