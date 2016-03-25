package ws

import (
	"fmt"
	"golang.org/x/net/websocket"
	"lim.com"
	"net/http"
)

func Im(w http.ResponseWriter, r *http.Request) {
	lim_com.R.HTML(w, http.StatusOK, "chat/im", map[string]string{
		"Title": "Websocket",
	})
}

func Echo(ws *websocket.Conn) {
	fmt.Println("ws,coming!")
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive!")
			break
		}
		fmt.Println("Received back from client: " + reply)
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
