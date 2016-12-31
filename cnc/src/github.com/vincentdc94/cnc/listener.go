package cnc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
)

func BeginListen() {
	listener, err := net.Listen("tcp", ":1337")

	if err != nil {

	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			//error
		}

		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {

	writer := bufio.NewWriter(connection)
	tpw := textproto.NewWriter(writer)

	reader := bufio.NewReader(connection)
	tpr := textproto.NewReader(reader)

	torrents := GetYifyTorrents(10)

	for _, torrent := range torrents {

		torrentJson, err := json.Marshal(torrent)

		if err == nil {
			tpw.PrintfLine(string(torrentJson))
		}

	}

	for {

		message, err := tpr.ReadLine()

		if err == nil {
			fmt.Println(message)
		}

	}

}
