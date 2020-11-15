package Week_02

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    wordMap := [26]int{}
    for i := 0; i < len(s); i++ {
        wordMap[s[i] - 'a']++
        wordMap[t[i] - 'a']--
    }

    for _, count := range wordMap {
        if count != 0 {
            return false
        }
    }

    return true
}