## go.wstext
go.wstext is a small wrapper package for [gorilla/websocket](https://github.com/gorilla/websocket). It adds .WriteText(msg) and .ReadText() methods to websocket.Conn. This is useful if you send a lot of "plain-text" commands back and forth.. Otherwise: don't bother and use the WriteMessage() and ReadMessage().

Example usage:
```
import (
	"github.com/GeertJohan/go.wstext"
	"github.com/gorilla/websocket"
)

conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
if err != nil {
	// handle error
}

textconn := wstext.Conn{conn}

text, err := textconn.ReadText()

err = textconn.WriteText("Hello, world")
```

### License
This project is licensed under a Simplified BSD license. Please read the [LICENSE file](LICENSE).
