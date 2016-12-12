package cnc

import (
	"bufio"
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

	tp.PrintfLine("Uplink established")
	tp.PrintfLine(".")
	tp.PrintfLine("..")
	tp.PrintfLine("...")
	tp.PrintfLine("Welcome back, Commander")

	torrents := GetYifyTorrents(10)

	for _, torrent := range torrents {
		//if torrent.Seeds <= 10 {
		tp.PrintfLine(torrent.Hash)
		//}
	}

}
