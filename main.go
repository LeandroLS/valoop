package valoop

func IsSameValue(a interface{}, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

func IntSliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
