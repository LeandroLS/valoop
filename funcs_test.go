package valoop

import (
	"testing"
)

func TestIsSameValue(t *testing.T) {
	isTrue := IsSameValue("a", "a")
	isFalse := IsSameValue("a", "b")
	if isTrue != true {
		t.Errorf("result %t, expected %t", isTrue, true)
	}

	if isFalse != false {
		t.Errorf("result %t, expected %t", isFalse, false)
	}
}

func TestIntSliceContains(t *testing.T) {
	intslice := []int{2}
	isTrue := IntSliceContains(intslice, 2)
	isFalse := IntSliceContains(intslice, 3)
	if isTrue != true {
		t.Errorf("result %t, expected %t", isTrue, false)
	}
	if isFalse != false {
		t.Errorf("result %t, expected %t", isFalse, false)
	}
}
