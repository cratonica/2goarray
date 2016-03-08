package main

import "testing"

func Test_Sample(t *testing.T) {
	if DATA == nil {
		t.Error("DATA nil")
		return
	}
	dataStr := string(DATA)
	if dataStr != "hello world\n" {
		t.Error("DATA not equal")
	}
}
