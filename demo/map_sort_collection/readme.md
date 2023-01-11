# 有序大数据（>1万条）时集合的选择
　　有序集合，可以选择有，slice与tree，加上数据多，就只剩下tree了。　　
为什么数据多时不能使用slice呢。因为在数据比较多，当内存紧张时，分配连续的内存会是一个问题。　　
tree主要有两种btree与llrb。一般的说法是llrb能有更好的插入删除性能。　　
下面是一次bench测试的一次结果　　
go test -bench . -cpu=1　　
goos: linux　　
goarch: amd64　　
pkg: github.com/peacess/go/demo/map_sort_collection　　
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkInsert_LLRB             4279257               264.8 ns/op　　
BenchmarkGet_LLRB               22825192                61.35 ns/op　　
BenchmarkDelete_LLRB             4841445               275.2 ns/op　　
BenchmarkInsert_TidwallBtree    13282300               104.7 ns/op　　
BenchmarkGet_TidwallBtree       13392019               104.1 ns/op　　
BenchmarkDelete_TidwallBtree    12050533               117.1 ns/op　　
BenchmarkInsert_Btree           12098299               114.2 ns/op　　
BenchmarkGet_Btree              12187546               117.3 ns/op　　
BenchmarkDelete_Btree            4264225               288.7 ns/op　　
BenchmarkInsert_GenericBtree     1000000              1300 ns/op　　
BenchmarkGet_GenericBtree        1739172               757.4 ns/op　　
BenchmarkDelete_GenericBtree     1000000              1172 ns/op　　
PASS　　
ok      github.com/peacess/go/demo/map_sort_collection  36.435s　　

从测试结果来看btree的性能比较好（这个可能与我的测试有关，也可能是库的实现相关），依次是是github.com/tidwall/btree，github.com/google/btree，从稳定上考虑最后选择了google的btree　　

总结：　　
数据结构与算法，从时间复杂度上看，一般考虑３个方面：　　
最差情况，平均情况，最好情况。这些内容作一些研究都能找到答案，但项目最终选择什么，第一应该考虑的是需求。　　
真正的需求是什么，这个问题有时很难具体确定（比如，大方向是有性能需要求，但是要具体下来，需要比较多的尝试）　　
我一般的做法是，
首先由大脑运行一下，如果面太多或没有明显的结论时，需要工具来帮助。写test或Ｄemo（大多数时，我会写test，因为快速且可以保留下来，以后在开发时使用）来帮助，把真正的需求找出来　　

注：项目最后经过调整后，从原来的延迟１００秒，降到延迟２０秒（我查看延迟的方法会有最大６秒的误差），而需求要３秒内，加上误差那要９秒。　　
等以后再找时间解决吧

后续：　查看延迟的方法最大有６秒的误差，为什么不找办法解决。而是需求不在那里，它不能排到当前的优先级中。

