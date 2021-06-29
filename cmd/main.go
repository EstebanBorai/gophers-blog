package main

import s "github.com/EstebanBorai/gosk/pkg/server"

func main() {
	var httpServer s.Server = s.MakeServer(4200)

	httpServer.UseLogs()
	httpServer.Serve()
}
