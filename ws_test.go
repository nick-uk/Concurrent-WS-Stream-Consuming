package main

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// StartTestWSServer is a fake server to test my client
func StartTestWSServer() {

	port := "12345"
	http.HandleFunc("/ws", wsServerTestHandler)

	server := &http.Server{
		Addr: "127.0.0.1:" + port,
	}
	log.Println("Service listen on :" + port + " under /ws")
	log.Fatal(server.ListenAndServe())
}

func wsServerTestHandler(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer func() {
		log.Println("closing WS")
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "stream is over"))
		c.Close()
	}()

	writes := 0

	for {
		// Read input message with 30secs timeout
		c.SetReadDeadline(time.Now().Add(time.Second * 30))
		mt, _, err := c.ReadMessage()
		if err != nil {
			log.Println("[WsServer] Read:", err)
			break
		}
		//log.Println("[WsServer] Read:", string(message))

		// Send a random string with write timeout 30 secs
		c.SetWriteDeadline(time.Now().Add(time.Second * 30))
		err = c.WriteMessage(mt, []byte(genRandomStr()))
		if err != nil {
			log.Println("[WsServer] Write:", err)
			break
		}
		// Stop streaming after 1k writes
		writes++
		if writes > 20 {
			break
		}
	}
}

func Test_wsClient(t *testing.T) {

	// Start my fake mock server & the results reader. Wait a bit to ensure server is ready
	go StartTestWSServer()
	go readResults()
	time.Sleep(time.Second * 1)

	type args struct {
		host       string
		wsresponse chan string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// This should fail
		{"Connectivity check", args{host: "ws://127.0.0.1:1234", wsresponse: wsresponse}, true},
		// This should fail as well because of the wrong path
		{"Mock server", args{host: "ws://127.0.0.1:12345", wsresponse: wsresponse}, true},
		// This shouldn't fail
		{"Mock server2", args{host: "ws://127.0.0.1:12345/ws", wsresponse: wsresponse}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := wsClient(tt.args.host, tt.args.wsresponse); (err != nil) != tt.wantErr {
				t.Errorf("wsClient() NOT expected error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Log("wsClient() OK expected error:", err)
			}
		})
	}
	wg.Wait()
	log.Println("Stream is over & processing is done. Cheers")
}
