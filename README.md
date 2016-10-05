# go-script-benchmark

Benchmarking embedded scripting languages inside Go, currently [Otto](https://github.com/robertkrimen/otto) and [go-lua](https://github.com/Shopify/go-lua).

## Results

|**Benchmark**|**Result**|
|--------|---------:|
|BenchmarkGoLuaInit-4|1138 ns/op|
|BenchmarkGoLuaModulo-4 |16658 ns/op|
|BenchmarkGoLuaRecursion-4|52920 ns/op|
|BenchmarkOttoInit-4|166096 ns/op|
|BenchmarkOttoModulo-4|1970924782 ns/op|
|BenchmarkOttoRecursion-4|579791 ns/op|
|BenchmarkOttoCallOverhead-4|10624 ns/op|
|BenchmarkOttoRealistic-4|15382 ns/op|

Run the tests with `go test -bench=.`
