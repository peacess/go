
# 各种函数调用性能差异  
测试输出结果  
go test -bench . -cpu=1  
goos: linux  
goarch: amd64  
pkg: github.com/peacess/go/demo/lang_bench  
cpu: AMD Ryzen 9 3900X 12-Core Processor              
BenchmarkNoFun                  741094232                1.549 ns/op  
BenchmarkClosure                642063650                1.664 ns/op  
BenchmarkClosureNoParameter     798082786                1.523 ns/op  
BenchmarkCall                   460731896                2.605 ns/op  
BenchmarkFuncPoint              457155902                2.630 ns/op  
BenchmarkInterface              503673919                2.359 ns/op  
BenchmarkReflect                 3616756               330.8 ns/op  
PASS  
ok      github.com/peacess/go/demo/lang_bench   11.882s  

说明：  
BenchmarkNoFun（无函数调用）               
BenchmarkClosure（闭包）                  
BenchmarkClosureNoParameter（闭包无参）    
BenchmarkCall（一般函数调用）              
BenchmarkFuncPoint（函数指针）            
BenchmarkInterface（interface）          
BenchmarkReflect（反射）  


闭包与差不多相当于没有函数调用， 函数、函数指针、interface的调用差不多，所以不用担心interface的性能  
而反射需要大约300ns，如果要性能一定不要使用反射  
