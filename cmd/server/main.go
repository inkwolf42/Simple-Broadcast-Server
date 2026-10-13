package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"inkwolf.com/broadcast/internal/consts"
)

func broadcastMessageFromServer(msg string) {
	for _, c := range GetClients() {
		var buf []byte
		buf = fmt.Appendf(buf, "%-25s > %s\n", "Server", msg)
		c.Conn.Write(buf)
	}
	fmt.Printf("%-25s > %s\n", "Server", msg)
}
func broadcastMessage(client Client, msg string) {
	for _, c := range GetClients() {
		var buf []byte
		buf = fmt.Appendf(buf, "%-25s > %s", client.Name, msg)
		c.Conn.Write(buf)
	}
	fmt.Printf("%-25s > %s", client.Name, msg)
}
func handleConnection(conn net.Conn) {
	defer conn.Close()

	readingBuff := bufio.NewReader(conn)

	name, err := readingBuff.ReadString('\n')

	if err != nil {
		println("Client Disconnected!")
		return
	}

	if !strings.HasPrefix(name, "name=") {
		println("Invalid Arguments.")
		conn.Write([]byte("Server : Invalid Arguments"))
		return
	}

	name = name[5 : len(name)-1]
	nameLen := len(name)

	if nameLen > 25 || nameLen < 3 {
		println("Name Should be between 3 and 25 charachter")
		conn.Write([]byte("Name Should be between 3 and 25 charachter\n"))
		return
	}

	client := CreateClient(conn, name)
	AddClient(client)
	conn.Write([]byte(consts.ACCEPT_MSG + "\n"))
	broadcastMessageFromServer("The user \"" + name + "\" Has Joined.")
	defer DeleteClient(client.id)

	for {
		message, err := readingBuff.ReadString('\n')

		if err != nil {
			broadcastMessageFromServer("The user \"" + name + "\" Has Left.")
			break
		}
		broadcastMessage(client, message)
	}
}

func Run() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		print("failed to init server listener : ")
		println(err)
		return
	}

	defer listener.Close()
	fmt.Print("\033[2J")
	print("Server is running on localhost:8080.\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			print("Connection Error : ")
			println(err)
			continue
		}
		go handleConnection(conn)
	}
}
