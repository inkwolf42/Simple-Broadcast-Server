package server

import (
	"net"
	"sync"
)

var mu sync.RWMutex = sync.RWMutex{}
var connections []Client = []Client{}

func AddClient(c Client) {
	mu.Lock()
	connections = append(connections, c)
	mu.Unlock()
}

func GetClients() []Client {
	mu.RLock()
	defer mu.RUnlock()
	return connections
}

func DeleteClient(targetId int) {
	mu.Lock()
	defer mu.Unlock()
	for i, connection := range connections {
		if connection.id == targetId {
			connections[i] = connections[len(connections)-1]
			connections = connections[:len(connections)-1]
			break
		}
	}

}

func CreateClient(conn net.Conn, name string) Client {
	defer func() { client_current_id++ }()
	return Client{
		id:   client_current_id,
		Conn: conn,
		Name: name,
	}
}
