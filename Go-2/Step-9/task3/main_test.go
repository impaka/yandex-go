package main

import "testing"

func TestReverseString(t *testing.T) {
	res := ReverseString("Hello world!")
	if res != "!dlrow olleH" {
		t.Errorf("...")
	}
}
