package main

import "testing"

func TestAreAnagrams(t *testing.T) {
	resTrue := AreAnagrams("кот", "ток")
	if resTrue != true {
		t.Errorf("...")
	}
	resFalse := AreAnagrams("папа", "мама")
	if resFalse != false {
		t.Errorf("...")
	}
	resLenErr := AreAnagrams("чебурек", "велосипедист")
	if resLenErr != false {
		t.Errorf("...")
	}
}
