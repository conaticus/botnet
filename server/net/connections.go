// This file handles client connections to the server

package net

import (
	. "botnet/server/util"
	"encoding/json"
	"net"
	"time"
)

var Connections map[net.Addr]*net.Conn // The value is nil if the client is not currently connected, but has connected before.

func init() {
	Connections = make(map[net.Addr]*net.Conn)
}

func KeepAlive(conn net.Conn) {
	ticker := time.NewTicker(time.Duration(Config.KeepAliveInterval) * time.Second)

	go func() {
		for {
			<-ticker.C

			payload := Payload{Type: CodeKeepAlive}
			encodedPayload, _ := json.Marshal(payload)

			err := Write(conn, encodedPayload)
			if err != nil {
				ticker.Stop()
				break
			}
		}
	}()
}

func Write(conn net.Conn, payload []byte) error {
	_, err := conn.Write(payload)	
	if err != nil {
		Connections[conn.RemoteAddr()] = nil
	}

	return err
}

func OnlineConnectionCount() int {
	count := 0

	for _, conn := range Connections {
		if conn != nil {
			count += 1
		}
	}

	return count
}