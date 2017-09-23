//go:generate ../2gobytes --input sample.txt --output hello.go

package main

import "testing"

func Test_Sample(t *testing.T) {
	testData(t, Data)
	testData(t, Sample)
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
