package conn

import (
	"github.com/hueypark/mars/handler"
	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

var client *net.Client

func init() {
	client = net.NewClient("127.0.0.1:8080", handler.Handle)
}

// SendLogin sens login.
func SendLogin(id int64) {
	login := message.MakeLoginMessage(id)
	message.Write(client.Conn(), fbs.LoginID, login)
}
