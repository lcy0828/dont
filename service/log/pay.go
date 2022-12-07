package log

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var tickets = 100

func Pay() {
	var logmesg = LogMes("")
	rand.Seed(time.Now().UnixNano())
	wg.Add(4)
	var chan1 = make(chan int, 1)
	var chan2 = make(chan int, 1)
	var chan3 = make(chan int, 1)
	var chan4 = make(chan int, 1)
	go mai("1号窗口", chan1)
	go mai("2号窗口", chan2)
	go mai("3号窗口", chan3)
	go mai("4号窗口", chan4)
	for i := 0; i <= 1000; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		if tickets <= 0 {
			close(chan1)
			close(chan3)
			close(chan2)
			close(chan4)
			break
		}

		switch i % 4 {
		case 0:
			//fmt.Println(i+1, "号顾客正在1号窗口买票")
			s := fmt.Sprintf("%d"+"号顾客正在1号窗口买票", i+1)
			logmesg = LogMes((s))
			logmesg.INFO()
			chan1 <- i
		case 1:
			//fmt.Println(i+1, "号顾客正在2号窗口买票")
			s := fmt.Sprintf("%d"+"号顾客正在2号窗口买票", i+1)
			logmesg = LogMes((s))
			logmesg.INFO()
			chan2 <- i
		case 2:
			//fmt.Println(i+1, "号顾客正在3号窗口买票")
			s := fmt.Sprintf("%d"+"号顾客正在3号窗口买票", i+1)
			logmesg = LogMes((s))
			logmesg.INFO()
			chan3 <- i
		case 3:
			//fmt.Println(i+1, "号顾客正在4号窗口买票")
			s := fmt.Sprintf("%d"+"号顾客正在4号窗口买票", i+1)
			logmesg = LogMes((s))
			logmesg.INFO()
			chan4 <- i

		}

	}
	wg.Wait()
	fmt.Println("结束")
}
func mai(num string, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		i, ok := <-ch
		mutex.Lock()

		if tickets > 0 && ok {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			//fmt.Println(i+1, "号客人在", num, "买到票", tickets)
			s := fmt.Sprintf("%d号客人在%s买到票%d", i+1, num, tickets)
			logmesg := LogMes((s))
			logmesg.INFO()
			tickets--
		} else {
			//fmt.Println("来晚了，票卖完了")
			s := fmt.Sprintf("客人来晚了，票卖完了 窗口: %s", num)
			logmesg := LogMes((s))
			logmesg.WARN()
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}

}

func Paytest() {
	for i := 0; i <= 100; i++ {
		s := fmt.Sprintf("%d"+"号顾客正在买票", i+1)
		logmesg := LogMes((s))
		logmesg.INFO()
	}

}
func maitest(num string, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		i, ok := <-ch
		mutex.Lock()

		if tickets > 0 && ok {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			//fmt.Println(i+1, "号客人在", num, "买到票", tickets)
			s := fmt.Sprintf("%d号客人在%s买到票%d", i+1, num, tickets)
			logmesg := LogMes((s))
			logmesg.INFO()
			tickets--
		} else {
			//fmt.Println("来晚了，票卖完了")
			s := fmt.Sprintf("客人来晚了，票卖完了 窗口: %s", num)
			logmesg := LogMes((s))
			logmesg.WARN()
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}

}
