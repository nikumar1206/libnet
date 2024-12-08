package main

import (
	"flag"
	"strings"
)

type Protocol int

const (
	ProtocolTCP Protocol = iota
	ProtocolUDP
)

type Config struct {
	protocol       Protocol
	port           int
	bufferSize     int
	initialMessage string
}

func initFlags() *Config {
	protocol := flag.String("protocol", "TCP", "Protocol for messaging. Must be either TCP or UDP.")
	port := flag.Int("port", 6379, "Port to bind.")
	bufferSize := flag.Int("buffer", 1024, "Size of message buffer.")
	message := flag.String("message", "", "Message to send. Will be immediately flushed on connection if provided.")

	flag.Parse()

	var connProto Protocol
	normalized_protocol := strings.ToLower(*protocol)

	if normalized_protocol == "tcp" {
		connProto = ProtocolTCP
	} else {
		connProto = ProtocolUDP
	}

	return &Config{
		protocol:       connProto,
		port:           *port,
		bufferSize:     *bufferSize,
		initialMessage: *message,
	}
}
