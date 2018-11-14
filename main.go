package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"html/template"

	"github.com/delioda79/ttt/game"
	"github.com/gorilla/websocket"
)

var available uint8 = game.Cross
var addr = flag.String("addr", "localhost:9090", "http service address")

var upgrader = websocket.Upgrader{} // use default options

type play struct {
	ch chan []byte
}

type Move struct {
	X, Y int
}

func (pl play) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("Upgraded")
	defer c.Close()
	var player uint8

	if available > 0 {
		player = available
		available--
		fmt.Println("Player: ", player)
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			pos := &Move{}

			json.Unmarshal(message, pos)
			fmt.Printf("POS: %s %+v\n", string(message), pos)
			move := &game.Move{Player: player, X: pos.X, Y: pos.Y}
			moveMsg, _ := json.Marshal(move)
			pl.ch <- moveMsg

			rsp := <-pl.ch
			err = c.WriteMessage(mt, rsp)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	} else {
		c.WriteMessage(0, []byte("full"))
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/play")
}

func main() {
	flag.Parse()

	ch := make(chan []byte, 100)
	gm := game.NewNoughtCross()
	go gm.Run(ch)

	http.Handle("/play", play{ch: ch})
	http.HandleFunc("/", home)

	jsn, _ := json.Marshal(&Move{X: 1, Y: 2})
	fmt.Println("JSON", string(jsn))
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
