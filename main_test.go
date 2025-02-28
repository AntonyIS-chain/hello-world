package main

import "testing"

func TestGreet(t *testing.T) {
    result := Greet()
    expected := "Hello, Go!"

    if result != expected {
        t.Errorf("Expected %q but got %q", expected, result)
    }
}
