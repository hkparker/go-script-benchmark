package script

import (
	"testing"
)

func BenchmarkOttoInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testOttoInit()
	}
}

func BenchmarkOttoMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testOttoModulo()
	}
}

func BenchmarkOttoStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testOttoRecursion()
	}
}

//func BenchmarkOttoX(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		testOttoX()
//	}
//}
