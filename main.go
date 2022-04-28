package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	for sig := range signalChan {
		fmt.Println(sig)
		time.Sleep(5 * time.Second)

		if sig == os.Kill {
			os.Exit(2)
		}
	}
}
