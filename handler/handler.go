package handler

import (
	"fmt"
	"log"
	"net"

	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// Handle handles message.
func Handle(conn net.Conn) error {
	id, body, err := message.ReadMessage(conn)
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
