package main

import (
	//"encoding/json"
	"code/code3"
	"os"
	"fmt"
	"encoding/json"
	"reflect"
)

var (
	test = code3.Config{}
)

type Movie struct {
	Titile	string
	Year	int		`json:"year"`
	Color	bool	`json:"color"`
	Actors	[]string
}

func main() {
	//Testconfig()
	var movies = []Movie{
		{Titile:"filmtitle1",Year:	1942,Color: false,Actors:[]string{"humphrey","ingrid"}},
		{Titile:"filmtitle2",Year:	1943,Color: false,Actors:[]string{"stev","jacq"}},
		{Titile:"filmtitle3",Year:	1944,Color: false,Actors:[]string{"stev11","jacq33"}},
	}
	data ,err := json.Marshal(movies)
	fmt.Printf("the data type is : %v\n",reflect.TypeOf(data))
	if err != nil {
		fmt.Printf("marshal failed~\n")
	}
	fmt.Printf("the data is : %s\n",data)
	var titles []struct{ Title string }
	err3 := json.Unmarshal(data, &titles)
	if err != nil {
		fmt.Printf("failed ,err is : %v $$$$$$$$$$$$$$$\n",err3)
	}
	fmt.Printf("!!!!!!!!!!!!\n")
	fmt.Println(titles)
	ttest()
}

func Testconfig(){
	test1 := new(code3.Config)
	test1.Name = "test1name"
	test2 := code3.Config{}
	test2.Name = "test2name"
	test.Name = "testname"
	path := `D:\soft\goproject\src\code\conf\config1.json`
	Opentest(path)
}

func Opentest(path string){
	file,err := os.Open(path)
	if err != nil {
		fmt.Printf("openfile failed!!!!!!!!\n")
	}
	defer file.Close()
	fs,err := file.Stat()
	if fs.Size()==0  {
		fmt.Printf("file %v is empty\n",file)
	}
	fmt.Printf("the size is : %v\n,%v\n,%v\n",fs.Size(),fs.Name(),fs.Mode())
	buffer := make([]byte,fs.Size())
	n,err := file.Read(buffer)
	fmt.Printf("the n is : %v\n",n)
	fmt.Printf("the buffer is : %s\n",buffer)
	var config1 code3.TestCase
	//config1 := code3.Config{}
	fmt.Printf("!!!!!!!!!!!!!!!!!\n")
	err1 := json.Unmarshal(buffer,&config1)
	if err != nil {
		fmt.Printf("unmarshal failed : %v\n",err1)
	}
	fmt.Printf("start @!@!@!@!@!@!@\n")
	//fmt.Printf("the config is : %v\n",config1)
	//fmt.Printf("the &config is : %v\n",&config1)
	fmt.Printf("the config name is : %v, Qtype is : %s\n",config1.Case[0].Name)
	for i,v := range config1.Case{
		fmt.Printf("the i is : %s\n",i)
		st1 := v.Name
		st2 := v.Title
		//st3 := v.Qtype[0]
		fmt.Printf("#######################\n")
		fmt.Printf("st1: %s, st2 : %s\n",st1,st2)
	}
	fmt.Printf("@@@@@@@@@@@@@@@\n")
}

func ttest () bool{
	s := "abcbaa"
	for i,v :=range s{
		fmt.Printf("the i is :%v\n",i)
		fmt.Printf("the v is : %c\n",v)
		if s[i] != s[len(s)-1-i] {
			fmt.Printf("false")
			return false
		}
	}
	fmt.Printf("true")
	return true
}