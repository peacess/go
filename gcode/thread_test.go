package gcode

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	cancel := make(<-chan struct{})
	const seconds = 2
	interval := time.Second * seconds
	timer := time.NewTimer(interval) //如果interval = 0，那么立刻运行
	for i := 0; i < 3; i++ {
		select {
		case <-cancel:
			//退出清理
			timer.Stop() //尽快清理timer
			return
		case <-timer.C: //这里也可以使用 “time.After(interval)”简化代码，但是不过After函数中，每次都会 new 一个新的timer出来，所以不建议使用
			start := time.Now()
			//do something
			diff := seconds - int(time.Now().Sub(start).Seconds())
			if diff <= 0 { //任务运行已经超过定时器了，依据业务选择怎么处理，这里只是简单的把定时改为一秒
				diff = 1
			}
			interval = time.Second * time.Duration(diff)
		}
		timer.Reset(interval)
	}
	timer.Stop()
}

func TestTickLong(t *testing.T) {

	we := sync.WaitGroup{}
	we.Add(2)
	timer := time.Tick(1)
	go func() {
		for range timer {
			fmt.Println("1")
			we.Done()
			we.Wait()
			break
		}
	}()

	go func() {
		for range timer {
			fmt.Println("2")
			we.Done()
			we.Wait()
			break
		}
	}()

	we.Wait()
	_ = timer
}

func TestTick(t *testing.T) {

	we := sync.WaitGroup{}
	const count = 5
	we.Add(count)
	now := time.Now()
	timer := time.NewTicker(time.Millisecond * 100)
	for i := 0; i < count; i++ {
		go func(i int) {
			for range timer.C {
				fmt.Println(i)
				we.Done()
				we.Wait()
				break
			}
		}(i)
	}
	we.Wait()
	fmt.Println(time.Now().Sub(now).Seconds())
	timer.Stop()
	_ = timer
}

func TestDoOnce(t *testing.T) {

	{ //只关闭一次channel，且不等待关闭完成。
		stopChannel := make(chan bool)
		closing := int32(0) //atomic没有操作bool的
		closeByAtomic := func() {
			if atomic.CompareAndSwapInt32(&closing, 0, 1) {
				close(stopChannel)
			}
		}

		go closeByAtomic() //多次并发关闭，是安全的
		go closeByAtomic()
		closeByAtomic()
	}

	{ //sync.Once 实现
		stopChannel := make(chan bool)
		once := sync.Once{}
		closeByOnce := func() {
			once.Do(func() {
				close(stopChannel)
			})
		}

		go closeByOnce() //多次并发关闭，是安全的
		go closeByOnce()
		closeByOnce()
	}
	//两次实现的区别在于，使用sync.Once时，closeByOnce()方法返回，那么close函数一定运行完成了。
	//而使用atomic实现时，closeByAtomic()方法返回，close函数不确定是否运行完成，可能运行完成，可能刚开始运行，可能还没有运行
}

func TestForI(t *testing.T) {
	{
		//下面是错误的使用方法
		wg := sync.WaitGroup{}
		wg.Add(3)
		for i := 0; i < 3; i++ {
			go func() {
				fmt.Print(i)
				wg.Done()
			}()
		}
		wg.Wait()
		//“0123”中的三个数的重复组合，而不是三个数字"012"的组合
	}
	{ //下面是正确做法，会输出三个数字"012"的组合
		wg := sync.WaitGroup{}
		wg.Add(3)
		for i := 0; i < 3; i++ {
			go func(d int) {
				fmt.Print(d)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
}
