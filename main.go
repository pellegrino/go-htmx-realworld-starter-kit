package main

import "log"

func main() {

	if err := runServer(); err != nil {
		log.Fatalf("Cannot start the server %e", err)
	}
}
