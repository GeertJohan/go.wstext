package wstext

import (
	"errors"
	"github.com/gorilla/websocket"
)

var ErrWrongType = errors.New("websocket: received wrong type")

type Conn struct {
	*websocket.Conn
}

// WriteText writes given string with messageType TextMessage
// Is actually a shorthand for (*Conn).WriteMessage(websocket.TextMessage, []byte(message))
func (c Conn) WriteText(message string) error {
	return c.WriteMessage(websocket.TextMessage, []byte(message))
}

// ReadText reads a new text message from the websocket.
// When retrieved message was not of type Text: wstext.ErrWrongType is returned.
func (c Conn) ReadText() (msg string, err error) {
	messageType, p, err := c.ReadMessage()
	if err != nil {
		return "", err
	}
	if messageType != websocket.TextMessage {
		return "", ErrWrongType
	}
	return string(p), nil
}
