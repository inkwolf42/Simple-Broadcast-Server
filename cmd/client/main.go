package client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"inkwolf.com/broadcast/internal/consts"
	"inkwolf.com/broadcast/internal/utils"
)

const LIMIT = 50

type Arguments struct {
	Url  string
	Name string
}

var connected bool = false
var done chan bool = make(chan bool)

func getArgs() Arguments {
	if len(os.Args) < 4 {
		println("Few Arguments")
		os.Exit(1)
	}
	if len(os.Args[3]) < 3 || len(os.Args[3]) > 25 {
		println("The Name Should be Between 3 and 25 charachter")
		os.Exit(1)
	}
	return Arguments{
		Url:  os.Args[2],
		Name: os.Args[3],
	}
}

func SendName(conn net.Conn, args *Arguments) {
	_, err := conn.Write([]byte("name=" + args.Name + "\n"))
	utils.ExitIfErr(err)
	message, err := bufio.NewReader(conn).ReadString('\n')
	utils.ExitIfErr(err)
	message = message[:len(message)-1]
	if message != consts.ACCEPT_MSG {
		fmt.Println("Message : " + message)
		utils.ExitIfErr(errors.New(message))
	}
	connected = true
}

func Run() {

	args := getArgs()

	conn, err := net.Dial("tcp", args.Url)
	if err != nil {
		println("Faild To Connect : ", err)
		return
	}

	defer conn.Close()

	messages = Stack[string]{
		values: []string{},
		limit:  LIMIT,
	}

	SendName(conn, &args)
	messages.Push(fmt.Sprintf("Connected to : %s Succsessfuly", args.Url))
	ReDraw()

	go ListenFormessages(conn)
	go HandelInputs(conn)

	<-done

}

func HandelInputs(conn net.Conn) {

	reader := bufio.NewReader(os.Stdin)
	for connected {
		message, err := reader.ReadString('\n')
		utils.ExitIfErrCM(err, "Reading Error : ")

		if strings.EqualFold(message, "/quit\n") {
			connected = false
			ClearTerminal()
			fmt.Println("LogOut.")
			break
		}

		_, err = conn.Write([]byte(message))
		utils.ExitIfErrCM(err, "Connection Error (Sending) : ")
	}
	done <- true

}

func ListenFormessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for connected {
		message, err := reader.ReadString('\n')
		if connected {
			utils.ExitIfErrCM(err, "Connection Error (Reciving) : ")
		}
		messages.Push(message)
		ReDraw()
	}

}
