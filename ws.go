package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsClient(host string, wsresponse chan string) error {

	// Let's put the errors in a channel for further checking
	wsErrors := make(chan error)

	log.Printf("=> Connecting to %s", host)
	clientConn, _, err := websocket.DefaultDialer.Dial(host, nil)
	if err != nil {
		return errors.New("Connecting: " + err.Error())
	}
	defer clientConn.Close()

	go func() {
		for {
			// Write something to ws server
			if err := clientConn.WriteMessage(websocket.TextMessage, []byte("I got it")); err != nil {
				wsErrors <- err
				return
			}
			// Read remote stream
			wstype, message, err := clientConn.ReadMessage()
			if err != nil {
				log.Println("ReadMessage()", err)
				wsErrors <- err
				return
			}
			log.Println("Send:", string(message), " of type:", wstype, "for processing")
			wsresponse <- string(message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := clientConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return errors.New("Ping Error " + err.Error())
			}
		case err := <-wsErrors:
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				// This is okay, in this case means server has no more stream
				close(wsresponse)
				return nil
			}
			return err
		}
	}
}
