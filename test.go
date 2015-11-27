package main

import (
	// "testing"
	// "time"
	"fmt"
	// "strings"
)

// func Benchmark(b *testing.B) {
// 	customTimerTag := false
// 	if customTimerTag {
// 		b.StopTimer()
// 	}
// 	b.SetBytes(12345678)
// 	time.Sleep(time.Second)
// 	if customTimerTag {
// 		b.StartTimer()
// 	}
// }

type HttpCli struct {
	url string
	//handle typealias.Handler
}

// var a = new(HttpCli)

// type Handler func(str string) (*HttpCli, string)

// var HandlerMap = map[string]Handler{"wget": a.Get, "wpost": a.Post}

// func (self *HttpCli) Get(URL string) (*HttpCli, string) {
// 	return self, ""
// }

// func (self *HttpCli) Post(URL string) (*HttpCli, string) {
// 	return self, ""
// }

func main() {
	// var a = []string{"hello", "world", "!", "listen"}
	// var b = []string{"wallstreet", "myth", "rrr", "hello"}
	// var count int
	// str1 := strings.Join(a, ",")
	// str2 := strings.Join(b, ",")
	// if str1 != str2 {
	// 	for k, v := range b {
	// 		if a[0] == v {
	// 			count = k
	// 		} else {
	// 			count = len(b) - 1
	// 		}
	// 	}
	// }
	// fmt.Println(b[:count])
	var arr = make([]int, 0)
	// var b = 1
	arr1 := []int{1, 2, 3}
	arr = arr1
	fmt.Println(arr)
	// var c = 2
	// var d = 3
	// arr = append(arr, b)
	// arr = append(arr, c)
	// arr = append(arr, d)
	// fmt.Println(arr)
}
