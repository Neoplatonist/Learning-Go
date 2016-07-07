// https://medium.com/@olahol/writing-real-time-web-apps-in-go-chat-4aa058644f73#.ef1bfsmww
// Ola Holmström

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "public")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":3000")
	log.Println("Initialized On: localhost:3000")
}
