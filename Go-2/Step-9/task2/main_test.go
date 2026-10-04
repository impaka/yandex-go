package main

import "testing"

func TestContains(t *testing.T) {
	resTrue := Contains([]int{1,2,3},3)
	if resTrue != true {
		t.Errorf("...")
	}
	resFalse := Contains([]int{1,2,3},4)
	if resFalse != false {
		t.Errorf("...")
	}

}
