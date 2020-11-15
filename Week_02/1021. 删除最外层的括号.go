package Week_02

func removeOuterParentheses(S string) string {
    res := make([]byte, 0)
    count := 0
    for i:= 0; i < len(S); i++ {
		if S[i] == '(' {
			if count > 0 {
				res = append(res, S[i]) 
			}
			count++
		} else if S[i] == ')' {
			count--
			if count > 0 {
				res = append(res, S[i]) 
			}
		}
	}
    return string(res)
}
