# go-script-benchmark

Benchmarking embedded scripting languages inside Go, currently [Otto](https://github.com/robertkrimen/otto) and [go-lua](https://github.com/Shopify/go-lua).

## Results

|**Benchmark**|**Result**|
|--------|---------:|
|BenchmarkGoLuaInit-4|1037 ns/op|
|BenchmarkGoLuaModulo-4 |16283 ns/op|
|BenchmarkGoLuaRecursion-4|33616 ns/op|
|BenchmarkOttoInit-4|161343 ns/op|
|BenchmarkOttoModulo-4|1868061122 ns/op|
|BenchmarkOttoRecursion-4|356973 ns/op|
|BenchmarkOttoCallOverhead-4|10415 ns/op|
|BenchmarkOttoRealistic-4|15177 ns/op|

Run the tests with `go test -bench=.`
