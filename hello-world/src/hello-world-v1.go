package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
        version = "1"
)

func main() {
	/* Clean up on ctrl-c */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

        ver := os.Getenv("version")
        if version == "" {
                version = ver
        }

	for {
		fmt.Printf("Hello-world App - Deployment Version %s\n", version)
		time.Sleep(time.Second * 3)
	}
}
