package main

import (
	"runtime"
	"fmt"
)
//并发 goroutine

//goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()//表示让cpu把时间片让给别人,下次某个时刻恢复执行改goroutine
		fmt.Println(s)
	}
}
/**
默认情况下，调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，
需要在我们的程序中显式调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。
GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。
如果n < 1，不会改变当前设置。以后Go的新版本中调度得到改进后，这将被移除。这里有
一篇Rob介绍的关于并发和并行的文章：http://concur.rspace.googlecode.com/hg/talk/concur.html#landing-slide”

 */

func main() {
	go say("world")//开启一个新的goroutines执行
	say("hello")//当前的goroutines执行
}