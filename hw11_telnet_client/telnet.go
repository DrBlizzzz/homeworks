package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

var globalErr error

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &TelnetClientImplementation{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

type TelnetClientImplementation struct {
	io.Closer
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	ctx     context.Context
	conn    net.Conn
	cancel  context.CancelFunc
	pipe    chan int
}

func NewTelnetClientImplementation(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) *TelnetClientImplementation {
	client := TelnetClientImplementation{}
	client.address = address
	client.timeout = timeout
	client.in = in
	client.out = out
	client.ctx = nil
	client.conn = nil
	client.cancel = nil
	client.pipe = nil
	return &client
}

func (client *TelnetClientImplementation) Connect() error {
	dialer := &net.Dialer{}
	client.pipe = make(chan int)
	client.ctx, client.cancel = context.WithTimeout(context.Background(), client.timeout)
	conn, err := dialer.DialContext(client.ctx, "tcp", client.address)
	client.conn = conn
	if err != nil {
		return err
	}
	return nil
}

func (client *TelnetClientImplementation) Send() error {
	go func() {
		scanner := bufio.NewScanner(client.in)
		for scanner.Scan() {
			select {
			case <-client.ctx.Done():
				return
			default:
				text := fmt.Sprintf("%s\n", scanner.Text())
				client.conn.Write([]byte(text))
				fmt.Println(text)
				client.pipe <- 0
			}
		}
		fmt.Println("out Send")
		globalErr = scanner.Err()
	}()
	return globalErr
}

func (client *TelnetClientImplementation) Receive() error {
	scanner := bufio.NewScanner(client.conn)
	net.Conn.
	for {
		select {
		case scanned := <-scanner.Text():
			go func() {

				select {
				case <-client.ctx.Done():
					return
				case <-client.pipe:
					if scanner.Scan() {
						text := fmt.Sprintf("%s\n", scanned)
						client.out.Write([]byte(text))
						fmt.Println(text)
					} else {
						globalErr = scanner.Err()
					}
				}
				fmt.Println("out Receive")
			}()
		}
	}
	return globalErr
}
