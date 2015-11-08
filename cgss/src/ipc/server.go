package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"params"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcSever(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
				return
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			if err != nil {
				fmt.Println(err)
			}
			c <- string(b)
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new sesson has been created successfully.")
	return session
}
