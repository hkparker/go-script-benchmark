package script

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"math/rand"
	"testing"
)

type Frame struct {
	Number   int
	Name     string
	Variable int
}

func BenchmarkOttoInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testOttoInit()
	}
}

func BenchmarkOttoModulo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testOttoModulo()
	}
}

func BenchmarkOttoRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testOttoRecursion()
	}
}

func BenchmarkOttoCallOverhead(b *testing.B) {
	vm := otto.New()
	alerts := make(chan string, 0)
	go func(chan string) {
		for {
			<-alerts
		}
	}(alerts)

	vm.Set("alert", func(call otto.FunctionCall) otto.Value {
		alerts <- call.Argument(0).String()
		return otto.Value{}
	})
	vm.Run(`
		var analyze = function(frame) {
	`)

	for i := 0; i < b.N; i++ {
		frame := Frame{
			Number: rand.Int(),
		}
		frame_bytes, _ := json.Marshal(frame)
		vm.Run(fmt.Sprintf("analyze(%s)", string(frame_bytes)))
	}
}

func BenchmarkOttoRealistic(b *testing.B) {
	vm := otto.New()
	alerts := make(chan string, 0)
	go func(chan string) {
		for {
			<-alerts
		}
	}(alerts)

	vm.Set("alert", func(call otto.FunctionCall) otto.Value {
		alerts <- call.Argument(0).String()
		return otto.Value{}
	})
	vm.Run(`
		var max = 0
		var map = {}

		var analyze = function(frame) {
			if (frame.Number > max)	{
				alert("setting max")
				max = frame.Number
			}
			map[frame.Name] = frame.Variable
		}
	`)

	for i := 0; i < b.N; i++ {
		frame := Frame{
			Name:     "Foo",
			Number:   rand.Int(),
			Variable: rand.Int(),
		}
		frame_bytes, _ := json.Marshal(frame)
		vm.Run(fmt.Sprintf("analyze(%s)", string(frame_bytes)))
	}
}

//func BenchmarkOttoX(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		testOttoX()
//	}
//}
