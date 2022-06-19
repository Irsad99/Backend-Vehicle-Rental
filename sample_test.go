package main

import "testing"

func TestHelloname(t *testing.T) {
	result := Helloname("ebi")

	if result != "Hello ebi"{
		t.Fatal("Code Error")
	}
}