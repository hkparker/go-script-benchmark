package script

import (
	"github.com/Shopify/go-lua"
)

func testGoLuaInit() {
	lua.NewState()
}

func testGoLuaModulo() {
	l := lua.NewState()
	lua.DoString(l, lua_modulo_test)
}

func testGoLuaRecursion() {
	l := lua.NewState()
	lua.DoString(l, lua_recursion_test)
}

//func testGoLuaX() {
//	l := lua.NewState()
//	lua.DoString(l, lua_X_test)
//}
