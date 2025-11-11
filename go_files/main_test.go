package main

import (
	"os"
	"testing"
)

func Hello() string {
	return "Hello, World!"
}

func MyMsg1() string {
	return os.Getenv("VAR_T1")
}

func MyMsg2() string {
	return os.Getenv("VAR_T2")
}

func main() {
	println(Hello())
	println(MyMsg1())
	println(MyMsg2())
}

func TestHello(t *testing.T) {
	want := "Hello, World!"
	got := Hello()
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestMyMsg1(t *testing.T) {
	want := "T1"
	got := MyMsg1()
	if got != want {
		t.Errorf("TestMyMsg1() = %q, want %q", got, want)
	}
}

func TestMyMsg2(t *testing.T) {
	want := "T2"
	got := MyMsg2()
	if got != want {
		t.Errorf("TestMyMsg2() = %q, want %q", got, want)
	}
}
