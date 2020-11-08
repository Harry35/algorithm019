package Week_01

func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, val := range nums {
        rest := target - val
        if j, ok := hashMap[rest]; ok {
            return []int{j, i}
        }
        hashMap[val] = i
    }
    return nil
}