package handler

import (
	"log"

	"github.com/hueypark/marsettler/game/message/fbs"
)

func handleLoginResult(loginResult *fbs.LoginResult) {
	log.Println(loginResult.Id())
}
