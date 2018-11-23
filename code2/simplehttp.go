package main

import (
	"net"
	"os"
	"bytes"
	"fmt"
	"io"
)
func main() {

	if len(os.Args) != 2 {
	fmt.Fprintf(os.Stderr, "Usage0: %s\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage2: %s\n", os.Args[2])
	os.Exit(1)
	}
	service := os.Args[1]
	fmt.Printf("the args0 is : %v\n",os.Args[0])
	fmt.Printf("the args1 is : %v\n",service)
	conn, err := net.Dial("tcp", service)

	//conn, err := net.Dial("tcp", "baidu.com:80")
	fmt.Printf("the conn is : %v\n",conn)
	checkError(err)
	p, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	fmt.Printf("the p is: %v\n",p)
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
	fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	os.Exit(1)
	}
}
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	//var buf1 [512]byte
	//tt, _ := conn.Read(buf1[0:])
	//fmt.Printf("the tt is : %v\n",tt)
	//for {
		n, err := conn.Read(buf[0:])
		fmt.Printf("the n is : %v\n",n)
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				//break
				return nil,nil
			}
		return nil, err
		}
	//}
	return result.Bytes(), nil
}