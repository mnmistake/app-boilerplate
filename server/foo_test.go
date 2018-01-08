package main

import "testing"

func TestBase(t *testing.T) {
	if 1 != 1 {
		t.Fatalf("It doesn't pass")
	}
}
