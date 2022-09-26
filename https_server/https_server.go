package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	serv := &http.Server{
		Addr:              "127.0.0.1:8443",
		Handler:           http.TimeoutHandler(exampleHandler(), 2*time.Minute, ""),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("client active...")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		path := fmt.Sprintf("https://%s", serv.Addr)

		/*
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				log.Fatal(err)
			}
		*/
		fmt.Println("Doing req")

		resp, err := client.Get(path)
		if err != nil {
			log.Fatal(err)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Printf("%s\n%s", resp.Header, b)
	}()

	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", serv.Addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Serving...")

		serv.ServeTLS(listener, "./cert.pem", "./key.pem")
		if err != nil {
			log.Fatal(err)
		}
	}()
	//time.Sleep(10 * time.Minute)
	wg.Wait()
}

func exampleHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func(r io.ReadCloser) {
				_, _ = io.Copy(io.Discard, r)

			}(r.Body)
			_, _ = w.Write([]byte("Hello, client!"))
		},
	)
}
