package toiletbowl

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	pubsub *PubSub
)

type PubSub struct {
	upgrader    websocket.Upgrader
	connections []*websocket.Conn
}

// Add will take in a context and upgrade that connection to store in an array
func (p *PubSub) Add(c echo.Context) error {
	ws, err := p.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	p.connections = append(p.connections, ws)
	return nil
}

func (p *PubSub) Publish(message string) {
	errs := []error{}
	for _, ws := range p.connections {
		err := ws.WriteMessage(websocket.TextMessage, []byte(message))
		errs = append(errs, err)
	}
	log.Println(errs)
}

func Instance() *PubSub {
	if pubsub == nil {
		pubsub = &PubSub{}
	}

	return pubsub
}
