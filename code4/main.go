package main

import (
	"fmt"
	"time"
	"runtime"
)


func fortest(i int){
	a := 0
	for a < i {
		//fmt.Printf("%v\n",a)
		fmt.Printf("%v\n",time.Now().Second())
		time.Sleep(time.Second*3)
		a++
	}
	//time.Sleep(time.Second*5)
	Q <- 0
}

func fortest2(i int){
	a := 0
	for a < i {
		//fmt.Printf("%v\n",a)
		fmt.Printf("2222 %v\n",time.Now().Second())
		time.Sleep(time.Second*2)
		a++
	}
	//time.Sleep(time.Second*5)
	Q <- 0
}

var (
	Q = make(chan int)
)

func main(){
	x := runtime.NumCPU()
	fmt.Println(x)
	i := 4
	a := 0
	for a < 2 {
		//fortest(i)
		go fortest(i)
		go fortest2(i)
		fmt.Println("end2")
		a++
		<-Q
	}

}