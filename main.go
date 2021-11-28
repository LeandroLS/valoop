package valoop

func IsSameValue(a interface{}, b interface{}) bool {
	return a == b
}

func IntSliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
