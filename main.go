package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DanjumaLabs/SHIELD-DNS/filter"
	"github.com/DanjumaLabs/SHIELD-DNS/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: SHIELD-DNS <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "start":
		startDNS()

		err := filter.LoadBlacklist("blocked.txt")

		if err != nil {
			fmt.Println("Blacklist loading error:", err)
			return
		}

		fmt.Println("Blacklist loaded:", filter.BlockedDomain)

		server.StartServer()

	case "stats":
		ShowStats()

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: SHIELD-DNS add <domain>")
			return
		}

		//AddDomain(os.Args[2])//
		err := filter.AddDomain(os.Args[2], "blocked.txt")

		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("ADDED:", os.Args[2])

	default:
		fmt.Println("unknown command:", command)
	}

	filter.LoadBlacklist("blocked.txt")
}

func startDNS() {
	err := server.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	go server.StartLogger(server.AppCfg.LogFile)
	fmt.Printf("SHIELD-DNS Starting Smoothly on port %s...\n", server.AppCfg.Port)
}
