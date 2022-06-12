package main

import "testing"

func TestRun(t *testing.T) {
	_, err := run([]string{"-v"})
	if err != nil {
		t.Errorf("failed test %#v", err)
	}
}

func TestLs(t *testing.T) {

	_, err := ls("./")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	// _, err = ls("./aaa")
	// if err != nil {
	// 	t.Fatalf("failed test %#v", err)
	// }
}
