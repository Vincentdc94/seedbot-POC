package botnet

import (
	"fmt"
	"os"

	"github.com/anacrolix/torrent"
)

/*
magnet:?xt=urn:btih:QWG6DKIF4HKBWN3T6JTXE5UTNTQLNALN&dn=Snowden+(2016)+720p+BrRip+x264+YIFY&tr=udp://tracker.zer0day.to:1337/announce&tr=udp://tracker.coppersurfer.tk:6969/announce&tr=udp://mgtracker.org:6969/announce&tr=udp://tracker.leechers-paradise.org:6969/announce&tr=udp://tracker.sktorrent.net:6969/announce&tr=udp://explodie.org:6969/announce
*/

type TorrentData struct {
	Title string
	Url   string
	Hash  string
	Seeds int
	Peers int
}

type ClientError struct {
	Type   string
	Origin error
}

func (clientError ClientError) Error() string {
	return fmt.Sprintf("Error %s: %s\n", clientError.Type, clientError.Origin)
}

// Client manages the torrent downloading.
type Client struct {
	Client   *torrent.Client
	Torrent  *torrent.Torrent
	Progress int64
	Uploaded int64
	Config   ClientConfig
}

// ClientConfig specifies the behaviour of a client.
type ClientConfig struct {
	TorrentPath    string
	Port           int
	TorrentPort    int
	Seed           bool
	TCP            bool
	MaxConnections int
}

func DoTorrent(magnet string, torrentPort int) (client Client, err error) {

	var tor *torrent.Torrent
	var cli *torrent.Client

	cli, err = torrent.NewClient(&torrent.Config{
		DataDir:    os.TempDir(),
		NoUpload:   true,
		Seed:       true,
		DisableTCP: false,
		ListenAddr: fmt.Sprintf(":%d", torrentPort),
	})

	if err != nil {
		return client, ClientError{Type: "creating torrent client", Origin: err}
	}

	client.Client = cli

	tor, err = cli.AddMagnet(magnet)

	if err != nil {
		return client, ClientError{Type: "Adding magnet", Origin: err}
	}

	client.Torrent = tor

	go func() {
		<-tor.GotInfo()
		tor.DownloadAll()

	}()

	return
}

func (c *Client) Close() {
	c.Torrent.Drop()
	c.Client.Close()
}
