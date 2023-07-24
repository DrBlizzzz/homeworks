package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	duration := flag.Duration("timeout", time.Second*10, "a time")
	flag.Parse()
	address := strings.Join(flag.Args(), ":")
	tcpClient := NewTelnetClient(
		address,
		*duration,
		os.Stdin,
		os.Stdout,
	)
	if err := tcpClient.Connect(); err != nil {
		fmt.Println("Not connected")
		return
	}
	tcpClient.Send()
	tcpClient.Receive()
	tcpClient.Close()

}
