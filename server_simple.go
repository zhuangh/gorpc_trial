package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *bool) error {
	fmt.Println(string(line))
	res := strings.Split(string(line), " ")
	res_data := 0

	a := 0
	b := 0

	fmt.Println("size ", len(res))
	if len(res) >= 2 {
		a, _ = strconv.Atoi(res[0])
		//fmt.Println(res[0])
		//fmt.Println(strconv.Atoi(res[0]))
		//fmt.Println(strconv.Atoi(res[1]))
		b, _ = strconv.Atoi(res[1])
	}

	res_data = a + b

	/*
		for i := range res {
			fmt.Println(res[i])
		}
	*/

	fmt.Println(res_data)

	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello server")
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)

	fmt.Println("hello server register")
	rpc.Register(listener)

	fmt.Println("hello server accept")
	rpc.Accept(inbound)

	fmt.Println("hello server accept after")
}
