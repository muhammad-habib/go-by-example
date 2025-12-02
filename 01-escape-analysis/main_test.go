package main

import "testing"

var resultUser User
var resultUserPtr *User

func BenchmarkSingleByValue(b *testing.B) {
	var r User
	for i := 0; i < b.N; i++ {
		r = ByValue(i)
	}
	resultUser = r
}

func BenchmarkSingleByPointer(b *testing.B) {
	var r *User
	for i := 0; i < b.N; i++ {
		r = ByPointer(i)
	}
	resultUserPtr = r
}
