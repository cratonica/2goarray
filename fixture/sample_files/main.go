//go:generate ../../2gobytes --input ../sample.txt --input ../../README.md --output data.go

package main

import "fmt"

func main() {
	fmt.Println(string(*Index["../sample.txt"]))
	fmt.Println(Index)
}
