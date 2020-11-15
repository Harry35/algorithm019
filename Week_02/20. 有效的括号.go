package Week_02

func isValid(s string) bool {
    l := len(s)
    if l == 0 || l % 2 == 1 {
        return false
    }
    pairs := map[byte]byte{
        ')':'(',
        '}':'{',
        ']':'[',
    }

    stack := []byte{}
    for i := 0; i < l; i++ {
        if pairs[s[i]] > 0 {
            if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
                return false
            } 
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }

    return len(stack) == 0

}