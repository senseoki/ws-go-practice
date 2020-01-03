package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(w, "ws://"+r.Host+"/echo")
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var objmap map[string]interface{}
		_ = json.Unmarshal(message, &objmap)
		event := objmap["event"].(string)
		sendData := map[string]interface{}{
			"event": "res",
			"data":  nil,
		}
		switch event {
		case "open":
			log.Printf("Received: %s\n", event)
		case "req":
			sendData["data"] = objmap["data"]
			log.Printf("Received: %s\n", event)
		}
		refineSendData, err := json.Marshal(sendData)
		err = c.WriteMessage(mt, refineSendData)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
