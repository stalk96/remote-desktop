package server

import (
	"log"
)

func Start() {
	log.Println("Server started")
	wsServerStart()
	//initDb()
}
