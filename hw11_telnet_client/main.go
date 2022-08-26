package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		return errors.New("args error")
	}

	address := net.JoinHostPort(args[0], args[1])
	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "...Connected to %s", address)
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := client.Send(); err != nil {
			log.Println("error send", err)
		}
		fmt.Fprintln(os.Stderr, "...EOF")
		cancel()
	}()

	go func() {
		if err := client.Receive(); err != nil {
			log.Println("error receive", err)
		}
		fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
		cancel()
	}()

	<-ctx.Done()

	return nil
}
