package main

import (
	"flag"

	"github.com/hosod/fridge_server/app/internal/server"
	"github.com/hosod/fridge_server/app/internal/database"

)

func main() {
	var dev = flag.Bool("dev", false, "please specify -dev flag")

	flag.Parse()

	database.Init(*dev)
	defer database.Close()
	server.Init()
}
