package main

import "testing"

type addTest struct {
	arg1, arg2, exepected int
}

var addTests=[]addTest{
	addTest{2,3,5},
	addTest{4,5,9},
	addTest{6,7,13},
	addTest{2,5,7},

}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4,6)
	}
}