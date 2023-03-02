package lang_bench

import (
	"reflect"
	"testing"
)

var gdata = []int64{1, 90, 76, 6688}

//go:noinline
func call(data []int64) int64 {
	re := int64(0)
	for _, d := range data {
		re += d
	}
	return re
}

func BenchmarkNoFun(b *testing.B) {
	b.ResetTimer()
	sum := int64(0)
	for i := 0; i < b.N; i++ {
		re := int64(0)
		for _, d := range gdata {
			re += d
		}
		sum = re
	}
	_ = sum
}

func BenchmarkClosure(b *testing.B) {
	b.ResetTimer()
	sum := int64(0)
	for i := 0; i < b.N; i++ {
		sum = func(data []int64) int64 {
			re := int64(0)
			for _, d := range data {
				re += d
			}
			return re
		}(gdata)
	}
	_ = sum
}

func BenchmarkClosureNoParameter(b *testing.B) {
	b.ResetTimer()
	sum := int64(0)
	for i := 0; i < b.N; i++ {
		sum = func() int64 {
			re := int64(0)
			for _, d := range gdata {
				re += d
			}
			return re
		}()
	}
	_ = sum
}

func BenchmarkCall(b *testing.B) {
	b.ResetTimer()
	sum := int64(0)
	for i := 0; i < b.N; i++ {
		sum = call(gdata)
	}
	_ = sum
}

func BenchmarkFuncPoint(b *testing.B) {
	b.ResetTimer()
	fp := call
	sum := int64(0)
	for i := 0; i < b.N; i++ {
		sum = fp(gdata)
	}
	_ = sum
}

func BenchmarkInterface(b *testing.B) {
	b.ResetTimer()
	var caller Caller = &CallerImp{}
	sum := int64(0)
	for i := 0; i < b.N; i++ {
		sum = caller.Call(gdata)
	}
	_ = sum
}

func BenchmarkReflect(b *testing.B) {
	b.ResetTimer()
	var caller CallerImp = CallerImp{}
	f := reflect.ValueOf(&caller).MethodByName("Call")
	param := []reflect.Value{reflect.ValueOf(gdata)}
	sum := int64(0)
	for i := 0; i < b.N; i++ {
		tt := f.Call(param)
		sum = tt[0].Int()
	}
	_ = sum
}

type Caller interface {
	Call(data []int64) int64
}

type CallerImp struct{}

//go:noinline
func (c *CallerImp) Call(data []int64) int64 {
	re := int64(0)
	for _, d := range data {
		re += d
	}
	return re
}
