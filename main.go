package main

import (
	"log"
	"net/http"

	"github.com/cildhdi/go-board/conns"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	conns.InitConons()
	if err := http.ListenAndServe(":8001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)

		if err != nil {
			conn.Close()
			return
		}
		conns.AddConn(conn)

		go func() {
			for {
				msg, _, err := wsutil.ReadClientData(conn)
				if err != nil {
					conn.Close()
					conns.RemoveConn(conn)
					return
				}
				conns.Broadcast(msg)
			}
		}()
	})); err != nil {
		log.Fatal(err.Error())
	}
}
