package server

import "net"

var client_current_id = 1

type Client struct {
	id   int
	Conn net.Conn
	Name string
}
