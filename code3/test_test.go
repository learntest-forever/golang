package code3

import "testing"

type Config struct {
	Title	string	`json:"title"`
	Name 	string	`json:"name"`
	Qtype	[]string	`json:"qtype"`
}

type TestCase struct {
	 Case	[]Config	`json:"Case"`
}

func TestIsalindrome(t *testing.T){
	if true {
		t.Error("11111111111111")
	}
}