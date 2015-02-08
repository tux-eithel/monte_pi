// I'm going to test the three implemtation of test
package main

import "testing"

const iterazioni = 10000000

func BenchmarkSimple(b *testing.B) {
	// Test the simple way
	for n := 0; n < b.N; n++ {
		GetPISimple(iterazioni)
	}
}

func BenchmarkWithNoCore(b *testing.B) {
	// Test using only routines
	for n := 0; n < b.N; n++ {
		GetPIMultiChannel(iterazioni, 4)
	}
}

func BenchmarkWithCore(b *testing.B) {
	// Test using max CPU available
	for n := 0; n < b.N; n++ {
		GetPIMultiCPU(iterazioni)
	}
}
