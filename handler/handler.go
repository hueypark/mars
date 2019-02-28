package handler

import (
	"fmt"
	"log"

	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// Handle handles message.
func Handle(iClient interface{}) error {
	client := iClient.(*net.Client)

	id, body, err := message.ReadMessage(client.Conn())
	if err != nil {
		return err
	}

	switch id {
	case fbs.LoginResultID:
		loginResult := message.NewLoginResult(body)
		log.Println(loginResult.Id())
	case fbs.NodeID:
		node := message.NewNode(body)
		log.Println(node.ID())
		vector := fbs.Vector{}
		node.Position(&vector)
		log.Println(vector.X())
		log.Println(vector.Y())
	default:
		return fmt.Errorf("unhandled message id: %d", id)
	}

	return nil
}
