package main

import "fmt"

func main() {

	practice2()
}

func practice2() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("程序发生错误:%v\n", e)
		}
	}()

	//我们回到开始的那一个问题，一个需求,统计1~80000中有哪些素数
	//我们当时想到的方法是将统计素数的任务分配给4个CPU去做（我只有4个并行，用8个并发
	initDataChan := make(chan int, 50)
	calcHandlerChan := make(chan int, 8)
	printDataChan := make(chan int, cap(initDataChan))

	go productData(initDataChan, cap(initDataChan))

	fmt.Println(cap(calcHandlerChan))
	for i := 1; i <= cap(calcHandlerChan); i++ {
		go consumerData(initDataChan, printDataChan, calcHandlerChan, i)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-calcHandlerChan
		}
		close(printDataChan)
	}()

	sl := make([]int, 0)
	for {
		res, ok := <-printDataChan
		if !ok {
			break
		}
		sl = append(sl, res)
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println(sl)
	fmt.Printf("总的素数：%d\n", cap(sl))
	fmt.Println("main线程退出")

}

func consumerData(dataChan <-chan int, printDataChan chan<- int, handlerChan chan<- int, num int) {
	for {
		//got:
		//	v, ok := <-dataChan
		//	if !ok {
		//		break
		//	}
		//	for i := 0; i < v; i++ {
		//		for j := 0; j < v; j++ {
		//			if i*j == v {
		//				printDataChan <- v
		//				goto got
		//			}
		//		}
		//	}
		v, ok := <-dataChan
		if !ok {
			break
		}
		flag := false
		for i := 1; i < v; i++ {
			for j := 1; j < v; j++ {
				if i*j == v {
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
		if v > 1 && flag {
			printDataChan <- v
		}
	}
	handlerChan <- 1
	fmt.Printf("协程 %v 处理完成\n", num)
}

func productData(dataChan chan int, count int) {
	for i := 1; i <= count; i++ {
		dataChan <- i
	}
	fmt.Printf("数据生产完成:%d\n", len(dataChan))
	//关闭intChan
	close(dataChan)
}

func baseTest() {
	var c chan int = make(chan int, 2)
	fmt.Println(<-c)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	c <- 3

	fmt.Println(<-c)
}
