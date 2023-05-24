package main

import (
	"bufio"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &TelnetClientImplementation{
		address:    address,
		timeout:    timeout,
		in:         in,
		out:        out,
		conn:       nil,
		wg:         &sync.WaitGroup{},
		signalPipe: nil,
	}
}

type TelnetClientImplementation struct {
	io.Closer
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	conn       net.Conn
	wg         *sync.WaitGroup
	signalPipe chan os.Signal
}

func (client *TelnetClientImplementation) Connect() error {
	var err error
	client.signalPipe = make(chan os.Signal, 1)
	signal.Notify(client.signalPipe, syscall.SIGQUIT)
	client.conn, err = net.DialTimeout("tcp", client.address, client.timeout)
	client.wg.Add(2)
	if err != nil {
		return err
	}
	return nil
}

func (client *TelnetClientImplementation) Send() error {
	go func() {
		defer client.wg.Done()
		scanner := bufio.NewScanner(client.in)
		for scanner.Scan() {
			select {
			case <-client.signalPipe:
				return
			default:
				if _, err := client.conn.Write([]byte(scanner.Text() + "\n")); err != nil {
					return
				}
			}
		}
		if err := scanner.Err(); err != nil {
			return
		}
	}()
	return nil
}

func (client *TelnetClientImplementation) Receive() error {
	go func() {
		defer client.wg.Done()
		scanner := bufio.NewScanner(client.conn)
		for scanner.Scan() {
			select {
			case <-client.signalPipe:
				return
			default:
				if _, err := client.out.Write([]byte(scanner.Text() + "\n")); err != nil {
					return
				}
			}
		}
		if err := scanner.Err(); err != nil {
			return
		}
	}()
	client.wg.Wait()
	return nil
}

func (client *TelnetClientImplementation) Close() error {
	if err := client.conn.Close(); err != nil {
		return err
	}
	return nil
}
