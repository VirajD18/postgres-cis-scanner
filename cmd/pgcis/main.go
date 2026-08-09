package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/VirajD18/postgres-cis-scanner/internal/scanner"
	"github.com/VirajD18/postgres-cis-scanner/internal/servers"
)

func main() {

	configPath := flag.String(
		"config",
		"/etc/pgcis/servers.json",
		"Path to servers configuration file",
	)

	flag.Parse()

	serverList, err := servers.Load(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===================================")
	fmt.Println(" PostgreSQL CIS Scanner")
	fmt.Println("===================================")

	fmt.Printf("\nServers Found : %d\n", len(serverList))

	scanner.Run(serverList, 5)

	fmt.Println()
	fmt.Println("All Scans Completed")
	fmt.Println("Reports generated successfully.")
}
