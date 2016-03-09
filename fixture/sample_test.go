//go:generate ../2goarray --input sample.txt --output --array-name Hello hello.go

package main

import "testing"

func Test_Sample(t *testing.T) {
	testData(t, DATA)
	testData(t, Hello)
}

func testData(t *testing.T, d []byte) {
	if d == nil {
		t.Error("data nil")
		return
	}
	dataStr := string(d)
	if dataStr != "hello world\n" {
		t.Error("data not equal")
	}
}
