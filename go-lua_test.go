package script

import (
	"github.com/Shopify/go-lua"
	"testing"
)

func BenchmarkGoLuaInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lua.NewState()
	}
}

func BenchmarkGoLuaModulo(b *testing.B) {
	l := lua.NewState()
	for i := 0; i < b.N; i++ {
		lua.DoString(l, lua_modulo_test)
	}
}

func BenchmarkGoLuaRecursion(b *testing.B) {
	l := lua.NewState()
	for i := 0; i < b.N; i++ {
		lua.DoString(l, lua_recursion_test)
	}
}
