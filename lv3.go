package main

import "fmt"

var (
	ch1   = make(chan int)
	ch2 = make(chan int)
	ch3  = make(chan bool, 3)
	T      bool
)

func main() {

	go func() {
		for i :=0; i <= 1000; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	for i := 0; i < 3; i++ {
		go func() {
			for {
				T = true
				num, ok := <-ch1
				if !ok {
					break
				}
				for i := 2; i < num; i++ {
					if num%i == 0 {
						T = false
						break
					}
				}
				if T {
					ch2 <- num
				}
			}
			fmt.Println("有一个ch2 协程因为取不到数据，退出")
			ch3<- true
		}()
	}

	go func() {
		for i := 0; i < 3; i++ {
			<-ch3
		}
		close(ch2)
	}()
	for {
		shu, ok := <-ch2
		if !ok {
			break
		}
		fmt.Println(shu)
	}

}