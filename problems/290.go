package problems

func WordPattern(pattern string, s string) bool {
	pMap := make(map[rune]int)
	sMap := make(map[string]int)
	sSplit := splitString(s)
	if len(pattern) != len(sSplit) {
		return false
	}
	for i, val := range pattern {
		if pMap[val] != sMap[sSplit[i]] {
			return false
		}
		pMap[val] = i + 1
		sMap[sSplit[i]] = i + 1
	}
	return true
}

func splitString(s string) []string {
	result := []string{}
	current := ""
	for _, val := range s {
		if val == ' ' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(val)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
