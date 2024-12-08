package main

import (
	"fmt"
	"net"
)

type Connection interface {
	Write(string)
	Read() string
	Close() error
	GetConn() net.Conn
}

func NewConnection(c Config) Connection {
	switch c.protocol {
	case ProtocolTCP:
		return NewTCPConnection(&c)
	case ProtocolUDP:
		return NewUDPConnection(&c)
	default:
		return NewTCPConnection(&c)
	}
}

type TCPConnection struct {
	conn     net.Conn
	protocol Protocol
}

func NewTCPConnection(c *Config) *TCPConnection {
	conn, err := net.Dial("tcp", getPort(c.port))
	handleErr(err)

	return &TCPConnection{
		conn:     conn,
		protocol: ProtocolTCP,
	}
}

func (t *TCPConnection) Read() string {
	buf := make([]byte, 1024)

	_, err := t.conn.Read(buf)
	handleErr(err)

	return string(buf)
}

func (t *TCPConnection) Write(text string) {
	_, err := t.conn.Write([]byte(text))
	handleErr(err)
}

func (t *TCPConnection) GetConn() net.Conn {
	return t.conn
}

func (t *TCPConnection) Close() error {
	return t.conn.Close()
}

type UDPConnection struct {
	conn     net.Conn
	protocol Protocol
}

func NewUDPConnection(c *Config) *UDPConnection {
	conn, err := net.Dial("udp", getPort(c.port))
	handleErr(err)

	return &UDPConnection{
		conn:     conn,
		protocol: ProtocolUDP,
	}
}

func (t *UDPConnection) Read() string {
	buf := make([]byte, 1024)

	_, err := t.conn.Read(buf)
	handleErr(err)

	return string(buf)
}

func (t *UDPConnection) Write(text string) {
	_, err := t.conn.Write([]byte(text))
	handleErr(err)
}

func (t *UDPConnection) GetConn() net.Conn {
	return t.conn
}

func (t *UDPConnection) Close() error {
	return t.conn.Close()
}

func getFullAddress(c Connection) string {
	return fmt.Sprintf("%s://%s", c.GetConn().RemoteAddr().Network(), c.GetConn().RemoteAddr().String())
}
