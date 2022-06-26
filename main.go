package valoop

// IsSameValue check if interface a has the same value of interface b
func IsSameValue(a interface{}, b interface{}) bool {
	return a == b
}

// IntSliceContains check if int slice s contains int e
func IntSliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// StrSliceContains check if str slice s contains str e
func StrSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
