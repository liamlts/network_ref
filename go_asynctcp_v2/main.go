package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type ConnectionStats struct {
	Conns []net.Conn
}

func (c ConnectionStats) listConns() []string {
	var connList []string
	for _, conn := range c.Conns {
		connList = append(connList, conn.RemoteAddr().String())
	}
	return connList
}

func connWaiter(listener net.Listener, c chan net.Conn) {
	conn, err := listener.Accept()
	if err != nil {
		return
	}
	c <- conn
}

func Timeout(conn net.Conn, t time.Duration) error {
	dur := time.NewTimer(t)
	<-dur.C
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var conns ConnectionStats
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan net.Conn)

	for {
		go connWaiter(listener, c)
		conns.Conns = append(conns.Conns, <-c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conns.listConns())

		for _, v := range conns.Conns {
			Timeout(v, 7*time.Second)
		}
	}
}
