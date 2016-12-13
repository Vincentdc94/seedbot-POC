package main

import (
	"bufio"
	"net"
	"net/textproto"
	"time"

	"encoding/json"

	"github.com/vincentdc94/botnet"
)

var starterPort = 50007

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

		var torrentData botnet.TorrentData

		err = json.Unmarshal([]byte(line), &torrentData)

		if err != nil {
			break
		}
		//magnet:?xt=urn:btih:QWG6DKIF4HKBWN3T6JTXE5UTNTQLNALN&dn=Snowden+(2016)+720p+BrRip+x264+YIFY&tr=udp://tracker.zer0day.to:1337/announce&tr=udp://tracker.coppersurfer.tk:6969/announce&tr=udp://mgtracker.org:6969/announce&tr=udp://tracker.leechers-paradise.org:6969/announce&tr=udp://tracker.sktorrent.net:6969/announce&tr=udp://explodie.org:6969/announce

		botnet.DoTorrent("magnet:?xt=urn:btih:"+torrentData.Hash, starterPort)

		if err != nil {
			break
		}

		starterPort++

	}

}
