package Week_01

func plusOne(digits []int) []int {
        for i := len(digits) - 1; i >= 0; i-- {
             tmp := digits[i] + 1
             digits[i] = tmp % 10
             if digits[i] != 0 {
                     return digits
             }
         }
         digits = append([]int{1}, digits...)
         return digits
     }