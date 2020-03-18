package main

import (
	"log"
	"net"
	"sync"

	"github.com/gobwas/ws/wsutil"
)

// Conns 所有连接
type Conns struct {
	connections map[string]net.Conn
	lock        *sync.RWMutex
}

var conns *Conns

func initConons() {
	conns = &Conns{
		connections: map[string]net.Conn{},
		lock:        &sync.RWMutex{},
	}
}

func addConn(conn net.Conn) {
	conns.lock.Lock()
	defer conns.lock.Unlock()

	conns.connections[conn.RemoteAddr().String()] = conn
}

func removeConn(con net.Conn) {
	conns.lock.Lock()
	defer conns.lock.Unlock()

	delete(conns.connections, con.RemoteAddr().String())
}

func broadcast(data []byte) {
	conns.lock.RLock()
	defer conns.lock.RUnlock()

	for _, con := range conns.connections {
		if err := wsutil.WriteServerText(con, data); err != nil {
			log.Println(err.Error())
		}
	}
}
