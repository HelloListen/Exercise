package main

import (
// "testing"
// "time"
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

var a = new(HttpCli)

type Handler func(str string) (*HttpCli, string)

var HandlerMap = map[string]Handler{"wget": a.Get, "wpost": a.Post}

func (self *HttpCli) Get(URL string) (*HttpCli, string) {
	return self, ""
}

func (self *HttpCli) Post(URL string) (*HttpCli, string) {
	return self, ""
}
