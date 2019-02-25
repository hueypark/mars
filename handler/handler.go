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
		loginResult := message.MakeLoginResult(body)
		log.Println(loginResult.Id())
	default:
		return fmt.Errorf("unhandled message id: %d", id)
	}

	return nil
}
