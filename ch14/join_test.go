package prose

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	actual := JoinWithCommas(list)
	expected := "apple and orange"
	if actual != expected {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, actual, expected)
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	actual := JoinWithCommas(list)
	expected := "apple, orange, and pear"
	if actual != expected {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, actual, expected)
	}
}
