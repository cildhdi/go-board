package main

import (
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	initConons()
	http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)

		if err != nil {
			conn.Close()
			return
		}
		addConn(conn)

		go func() {
			for {
				msg, _, err := wsutil.ReadClientData(conn)
				if err != nil {
					conn.Close()
					removeConn(conn)
					return
				}
				broadcast(msg)
			}
		}()
	}))
}
