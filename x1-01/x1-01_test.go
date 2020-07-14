package main

import "testing"

func BenchmarkBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main1()
	}
}

func BenchmarkStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main3()
	}
}
