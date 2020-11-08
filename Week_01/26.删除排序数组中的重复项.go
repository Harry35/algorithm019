package Week_01

func reserve(nums []int, start, end int)  {
    for start < end {
        tmp := nums[start]
        nums[start] = nums[end]
        nums[end] = tmp
        start++
        end--
    }
}

func rotate(nums []int, k int)  {
    l := len(nums)
    k %= l
    reserve(nums, 0, l - 1)
    reserve(nums, 0, k - 1)
    reserve(nums, k, l - 1)
}
