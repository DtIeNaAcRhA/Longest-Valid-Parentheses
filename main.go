func longestValidParentheses(s string) int {
	max := 0
	stek := []int{-1}
	count := 0
  
	for i, ch := range s {
		if ch == '(' {
			stek = append(stek, i)
			continue
		} else {
			stek = stek[:len(stek)-1]
			if len(stek) == 0 {
				stek = append(stek, i)
			} else {
				count = i - stek[len(stek)-1]
				if count > max {
					max = count
				}
			}
		}
	}
	return max
}
