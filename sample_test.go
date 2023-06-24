package sample

import "testing"

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld("hoge")
	expected := "Hello world, hoge"
	if actual != expected {
		t.Errorf("actual %v, want %v", actual, expected)
	}
}
