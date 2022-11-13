package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/adylanrff/imrenagi-programming-challenges/write-your-own-http-server/httpparser"
	"github.com/adylanrff/imrenagi-programming-challenges/write-your-own-http-server/httpresponse"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "port for the server")
}

func main() {
	flag.Parse()

	srv, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}
	defer srv.Close()

	fmt.Printf("serving at %s\n", srv.Addr().String())

	for {
		conn, err := srv.Accept()
		if err != nil {
			fmt.Printf("connection error, err=%+v", err)
		}

		// for simplicity, just block lol
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var (
		resp *httpresponse.HTTPResponse
	)

	httpPayload, err := httpparser.ParseHTTP(conn)
	if err != nil {
		return
	}

	if httpPayload.Method != "GET" {
		resp = httpresponse.NewHTTPResponse(400, nil, "Only GET is supported")
		fmt.Fprint(conn, resp.ToHTTPFormat())
		return
	}

	resp = httpresponse.NewHTTPResponse(200, nil, fmt.Sprintf("Accessed at %s", time.Now().Format(time.RFC1123)))
	fmt.Fprint(conn, resp.ToHTTPFormat())
}
