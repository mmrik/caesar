package main

import "testing"

func TestRot13HelloWorld(t *testing.T) {
	result := ceasarCipher("Hello world! ;)", 13, english)
	expected := "Uryyb jbeyq! ;)"

	if result != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, result)
	}
}

func TestShift1Swedish(t *testing.T) {
	result := ceasarCipher("abcåäö", 1, swedish)
	expected := "bcdäöa"

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
