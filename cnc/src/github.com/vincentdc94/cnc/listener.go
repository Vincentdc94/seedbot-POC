package cnc

import (
	"bufio"
	"encoding/json"
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
	tp := textproto.NewWriter(writer)

	torrents := GetYifyTorrents(10)

	for _, torrent := range torrents {

		torrentJson, err := json.Marshal(torrent)

		if err == nil {
			tp.PrintfLine(string(torrentJson))
		}

	}

}
