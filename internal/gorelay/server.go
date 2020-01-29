package server

import (
	"fmt"
	"log"
	"github.com/googollee/go-socket.io"
)

func NewServer() (*socketio.Server) {
	Server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	Server.OnConnect("/", func(s socketio.Conn) error {
			s.SetContext("")
			fmt.Println("connected:", s.ID())
			return nil
	})

	Server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have " + msg)
	})

	Server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	Server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	// server.OnError("/", func(s socketio.Conn, e error) {
	// 	fmt.Println("meet error: ", e)
	// })

	Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	return Server
}