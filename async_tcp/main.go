package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func handleConn(conn net.Conn) {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Print(err)
		return
	}
	buf := make([]byte, 1024)
	go func() {
		defer conn.Close()
		defer wg.Done()
		for {
			conn.Write([]byte(">"))
			n, err := conn.Read(buf)
			if err != nil {
				log.Println(err)
				return
			}
			if n > 0 {
				err = conn.SetDeadline(time.Now().Add(30 * time.Second))
				if err != nil {
					log.Print(err)

				}
			}
			str := string(buf)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("[%d] [%s]", n, buf[:n])
			if strings.Contains(str, "exit") {
				conn.Write([]byte("Goodbye\n"))
				conn.Close()
				return
			}
			if strings.Contains(str, "Hi") {
				conn.Write([]byte("Hello client, im the server.\n"))
			}
		}
	}()
	wg.Wait()

}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("Listening on [%s]\n", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		fmt.Printf("[Accepted conn] [%s]\n", conn.RemoteAddr().String())
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}
