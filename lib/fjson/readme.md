[see](https://github.com/json-iterator/go-benchmark/tree/master/src/github.com/json-iterator/go-benchmark)  

benchtest go json, result as follow :  

goos: linux  
goarch: amd64  
pkg: peacess/go/demo/fjson  
cpu: AMD Ryzen 9 7900X 12-Core Processor              
BenchmarkDecodeStdStructMedium-24                1000000             16323 ns/op             360 B/op          8 allocs/op  
BenchmarkEncodeStdStructMedium-24                1000000              1652 ns/op             985 B/op          6 allocs/op  
BenchmarkDecodeJsoniterStructMedium-24           1000000             13334 ns/op            3374 B/op         79 allocs/op  
BenchmarkEncodeJsoniterStructMedium-24           1000000               805.1 ns/op           984 B/op          6 allocs/op  
BenchmarkDecodeEasyJsonMedium-24                 1000000              5192 ns/op              96 B/op          3 allocs/op  
BenchmarkEncodeEasyJsonMedium-24                 1000000               465.9 ns/op           600 B/op          4 allocs/op  
BenchmarkDecodeSonicMedium-24                    1000000             11999 ns/op            2880 B/op          5 allocs/op  
BenchmarkEncodeSonicMedium-24                    1000000              1013 ns/op            1107 B/op          9 allocs/op  
PASS  
ok      peacess/go/demo/fjson   50.796s  
