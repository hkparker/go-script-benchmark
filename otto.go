package script

import (
	"github.com/robertkrimen/otto"
)

func testOttoInit() {
	otto.New()
}

func testOttoModulo() {
	vm := otto.New()
	vm.Run(javascript_modulo_test)
}

func testOttoRecursion() {
	vm := otto.New()
	vm.Run(javascript_recursion_test)
}

//func testOttoX() {
//	vm := otto.New()
//	vm.Run(javascript_X_test)
//}
