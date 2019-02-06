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
	login, size := message.MakeLoginMessage(id)
	_ = message.WriteHead(client, fbs.LoginID, size)
	_ = message.WriteBody(client, login)
}
