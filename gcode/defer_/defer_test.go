package defer_

import (
	"fmt"
	"testing"
)

// defer f(c),  其中参数c是在defer所在语句处理运行，也就是说如果c是一个函数的返回值，那在defer语句时，就已运行
// defer 中的return 语句也整个函数的return没有关系
// defer 中改变返回值的方法，是修改返回值对应的变量。如果变量是值类型，defer中修改无效，因在return执行后返回值的copy
// 如果返回值是一个函数调用，defer的运行也在函数后面

func TestParameter(t *testing.T) {
	i := 10
	p := func() int {
		return i
	}

	df := func(v int) {
		fmt.Println(v)
	}

	defer df(p()) //这里会打印出10, 而不是11, 因为参数会在defer处理传入
	i = 11
}

func TestDeferReturn(t *testing.T) {
	f := func() int {
		r := 10
		defer func() {
			r = 11
		}()
		return r
	}
	r := f()
	fmt.Println(r) //输出10,而不是11
}

func TestDeferReturnFun(t *testing.T) {
	f := func() int {
		r := 10
		defer func() {
			r = 11
		}()
		ff := func() int {
			return r
		}
		return ff()
	}
	r := f()
	fmt.Println(r) //输出10, 因为return后的函数先运行出返回值
}

func TestDeferReturnName(t *testing.T) {
	f := func() (re int) {
		re = 10
		defer func() {
			re = 11
		}()
		return re
	}
	r := f()
	fmt.Println(r) //输出11, 因为return后的函数先运行出返回值
}

func TestDeferReturnFunName(t *testing.T) {
	f := func() (re int) {
		re = 10
		defer func() {
			re = 11
		}()
		ff := func() int {
			return re
		}
		return ff()
	}
	r := f()
	fmt.Println(r) //输出11, 因为return后的函数先运行出返回值
}

func TestDeferReturnFunName2(t *testing.T) {
	f := func() (re int) {
		re = 10
		defer func() {
			re = 11
		}()
		return 9 //这一句不起作用，所以如果返回值有变量名，不要在return语句上带数据，免得误会
	}
	r := f()
	fmt.Println(r) //输出11, 因为此时 return的返回值没有作用
}
