//json 与channel
package main

import(
   "encoding/json"
   "fmt"

   "reflect"
)

type Student struct{
   Name string  `json:"Student_name"`
   Age int      `json:student_age`
   Age1 int     `json:"student_age1"`
   score int    `json:"stutend_score"`
}

type PipeData struct {
value int
handler func(int) int
next chan int
}


func main(){
    //声明
   var stu Student=Student{
      Name:"stu01",
      Age:10,
      Age1:11,
      score:100,
   }
   data,err:=json.Marshal(stu) //打包，返回值为byte
   if err!=nil{
      fmt.Println("json encode stu faild,err",err)
      return
   }
   fmt.Println(string(data))   //把byte转化成string

   var pdata PipeData
   fmt.Printf("the type of hander is : %v\n",reflect.TypeOf(pdata.handler))

   c := make(chan int, 1024)
   c <- 1
   c <- 2
   c <-5

   //for i := range c {
   //   fmt.Println("Received:", i)
   //}

   var fs = [4]func(){}

   for i :=0; i<3; i++ {
      defer fmt.Println("defer i = ",i)
      defer func() {
         fmt.Println("defer_closure i = ",i)
      }()
      fs[i] = func() {fmt.Println("closure i = ",i)}
   }
   for _,f := range fs {
      f()
   }
}
