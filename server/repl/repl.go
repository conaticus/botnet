package repl

import (
	. "botnet/server/net"
	. "botnet/server/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ConnectionsCommand = "connections" // Prints the current amount of online and offline connections to the server
)

func parseCommand(input string) {
	switch input {
		case ConnectionsCommand:
			fmt.Printf("There are currently %d/%d online connections.\n", OnlineConnectionCount(), len(Connections))
	}
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		select {
			case <-LogChannel: // If there is a log, we want to re-take the input
				continue
			default:
				parseCommand(strings.ToLower(input))
		}
	}
}