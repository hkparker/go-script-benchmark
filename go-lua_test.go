package script

import (
	"testing"
)

func BenchmarkGoLuaInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testGoLuaInit()
	}
}

func BenchmarkGoLuaModulo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testGoLuaModulo()
	}
}

func BenchmarkGoLuaRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testGoLuaRecursion()
	}
}

//func BenchmarkGoLuaX(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		testGoLuaX()
//	}
//}
