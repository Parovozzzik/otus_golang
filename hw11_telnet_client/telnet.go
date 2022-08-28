package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) *TelnetClient {
	return &TelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *TelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *TelnetClient) Close() error {
	return c.conn.Close()
}

func (c *TelnetClient) Send() error {
	_, err := io.Copy(c.conn, c.in)
	return err
}

func (c *TelnetClient) Receive() error {
	_, err := io.Copy(c.out, c.conn)
	return err
}
