package main

import (
	. "botnet/server/net"
	. "botnet/server/repl"
	. "botnet/server/util"
)

func main() {
	go StartServer()
	go LogHandler()

	StartRepl()
}