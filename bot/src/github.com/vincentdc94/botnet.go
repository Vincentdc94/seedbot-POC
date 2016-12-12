package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"time"
)

func main() {
	connectionHandler()
}

func connectionHandler() {
	conn, err := net.Dial("tcp", "127.0.0.1:1337")

	defer connectionHandler()
	defer time.Sleep(time.Second * 3)
	defer conn.Close()

	if err != nil {
		//handle error
	}

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	for {

		line, err := tp.ReadLine()

		if err != nil {
			break
		}

		fmt.Println(line)

	}

}
