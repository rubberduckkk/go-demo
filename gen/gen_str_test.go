package main

import (
	"testing"
)

func TestGenerator_RandString(t *testing.T) {
	result := ASCII.RandString(5)
	t.Logf("%s", result)
	//f, err := os.Create("input.txt")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer f.Close()
	//_, err = f.WriteString(result)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Logf("Generated file: %s", f.Name())
}
