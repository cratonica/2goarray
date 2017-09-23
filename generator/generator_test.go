package generator

import (
	"testing"
)

func Test_GenerateCode(t *testing.T) {
	code := NewGenerator()
	code.GenerateInfo = false
	err := code.AddFile("../fixture/sample.txt")
	if err != nil {
		t.Error(err)
	}
	if string(code.Data[0].Name) != "Sample" {
		t.Error("VarName not equal")
	}

	sourcecode := code.GenerateCode()
	expected := "package main\n\n"
	expected += "// Sample store the data of '../fixture/sample.txt' as a byte array.\n"
	expected += "var Sample []byte = []byte{\n"
	expected += "\t0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x0a, \n}\n\n"
	if string(sourcecode) != expected {
		t.Error("generated code not equal", string(sourcecode), expected)
	}
}

func Test_FilepathToStructName(t *testing.T) {
	var structNameTests = []struct {
		in  string
		out string
	}{
		{"foo", "Foo"},
		{"foo.txt", "Foo"},
		{"path/to/foo.txt", "Foo"},
		{"path/to/foo bar.txt", "Foo_bar"},
		{"path/to/foo-bar.txt", "Foo_bar"},
		{"path/to/foo~bar.txt", "Foo_bar"},
		{"path/to/foo.bar.txt", "Foo_bar"},
		{"path/to/foo,bar.txt", "Foo_bar"},
		{"path/to/foo:bar.txt", "Foo_bar"},
		{"path/to/foo;bar.txt", "Foo_bar"},
	}
	for _, tt := range structNameTests {
		result := FilepathToStructName(tt.in)
		if result != tt.out {
			t.Errorf("FilepathToStructName '%s' not equal to '%s'\n", result, tt.out)
		}
	}
}
