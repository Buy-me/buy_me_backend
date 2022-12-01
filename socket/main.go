package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func main() {
	router := gin.New()

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		ticker := time.NewTicker(1 * time.Second)

		go func() {
			i := 0
			for {
				<-ticker.C
				i++
				s.Emit("bye", i)
			}
		}()

		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		// s.Emit("reply", "have "+msg)
	})

	// server.OnEvent("/", "bye", func(s socketio.Conn) string {
	// 	// last := s.Context().(string)

	// })

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFile("/", "./demo.htm")

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
