package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	//client, err := rpc.Dial("tcp", "localhost:42586")
	client, err := rpc.Dial("tcp", "moore.ucsd.edu:42586")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	//in := bufio.NewReader("hello")
	for {

		fmt.Println("hi, please input so that I can send: ")
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply bool
		fmt.Println("hi")
		err = client.Call("Listener.GetLine", line, &reply)
		fmt.Println("GetLine: %d\n", reply)
		if err != nil {
			log.Fatal(err)
		}
	}
}
