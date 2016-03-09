package main

import (
	"testing"
)

func Test_GenerateCode(t *testing.T) {
	code := NewGenerator()
	code.GenerateInfo = false
	err := code.SetDataFromFile("./fixture/sample.txt")
	if err != nil {
		t.Error(err)
	}
	if string(code.VarName) != "Sample" {
		t.Error("VarName not equal")
	}

	sourcecode := code.GenerateCode()
	expected := "package main\n\n"
	expected += "var Sample []byte = []byte{\n"
	expected += "\t0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x0a, \n}\n\n"
	if string(sourcecode) != expected {
		t.Error("generated code not equal")
	}
}
