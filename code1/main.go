//文件读写操作，命令行参数交互
//读一个文件，将里面的int型数字写入到另一个文件
package main

import "fmt"
import (
	"flag"
	"reflect"
	"os"
	"bufio"
	"io"
	"strconv"
)

//命令行输入指定参数，falg.string()返回的是字符串指针类型
//var infile *string = flag.String("i","test.txt","input infile name")
var infile *string = flag.String("i","D:\\Program Files (x86)\\goproject\\src\\test\\src\\code\\code1\\test.txt","input infile name")
var outfile *string = flag.String("o","D:\\Program Files (x86)\\goproject\\src\\test\\src\\code\\code1\\outfile","input outfile name")

func readvalues(filein string) (values []int, err error) {
	dir1, _ := os.Getwd()
    	fmt.Println("dir1:", dir1)
	fmt.Printf("the file path is : %v\n",filein)
	f,err := os.Open(filein)
	if err != nil {
		fmt.Printf("read file false,the err is : %v\n",err)
		return
	}
	defer f.Close()

	br :=bufio.NewReader(f)
	values = make([]int,0)

	for {
		line,isPrefix,err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Print("the line is too long")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			fmt.Print("the err is: ",err1)
		}
		values = append(values,	value)
	}
	return
}

func  writevalues(values []int, outfile string) (err error) {
	file , err := os.Create(outfile)
	if err != nil {
		fmt.Printf("creat file failed,the err is %v \n",err)
		return err
	}
	defer file.Close()

	for  _ , value := range values{
		fmt.Printf("!!!!!!!!!!the value type is : %v\n",reflect.TypeOf(value))
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	//获取flag参数
	flag.Parse()
	if infile != nil {
		//reflect.TypeOf判断参数类型
		fmt.Printf("the type of infile is : %v\n", reflect.TypeOf(infile))
		fmt.Printf("infile= %v", *infile, " ", "outfile= %v\n", *outfile)
	}
	//获取当前目录
	dir, _ := os.Getwd()
    	fmt.Printf("dir is : %v\n", dir)
	fmt.Printf("the infile is : %v\n",*infile)
	values ,err := readvalues(*infile)
	if err != nil {
		fmt.Print("readvalues faild ,the err is : ",err)
	} else {
		fmt.Print("the values is : ",values)
		writevalues(values, *outfile)
	}

}