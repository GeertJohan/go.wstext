## go.wstext
go.wstext is a small wrapper package for [gorilla/websocket](https://github.com/gorilla/websocket). It adds .WriteText(msg) and .ReadText() methods to websocket.Conn. This is useful if you send a lot of "plain-text" commands back and forth.. Otherwise: don't bother and use the WriteMessage() and ReadMessage().

Example usage:
```
import (
	"github.com/GeertJohan/go.wstext"
	"github.com/gorilla/websocket"
	"net/http"
)

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		// handle error
	}

	textconn := wstext.Conn{conn}

	text, err := textconn.ReadText()
	if err != nil {
		if err == wstext.ErrWrongType {
			// a non-text message was received
		}
		// handle error
	}

	err = textconn.WriteText("Hello, world")
	if err != nil {
		// handle error
	}

	// All original methods on `websocket.Conn` are still available on `wstext.Conn`.
	textconn.WriteJson(someData)
}
```

### License
This project is licensed under a Simplified BSD license. Please read the [LICENSE file](LICENSE).
