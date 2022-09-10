package util

func ContainsString(s []string, v string) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func ReversedStringSlice(s []string) []string {
	for i := len(s)/2 - 1; i >= 0; i-- {
		j := len(s) - 1 - i
		s[i], s[j] = s[j], s[i]
	}
	return s
}
