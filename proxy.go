/*
 Simple Transparent Proxy in Go

*/

package main

import "fmt"
import "io"
import "net"
import "os"

var upstream_host string
var upstream_port string;


func handleUpstream(conn net.Conn, upstream net.Conn) {
	// Start listening to upstream
	io.Copy(upstream, conn)
}


func handleConnection(conn net.Conn) {
	upstream, err := net.Dial("tcp", upstream_host + ":" + upstream_port)
	if err != nil {
		fmt.Println("Couldn't connect to upstream")
	}
	go handleUpstream(conn, upstream)
	io.Copy(conn, upstream)
}


func main() {
    var args = os.Args
    if len(args) >= 3 {
        upstream_host = args[1]
        upstream_port = args[2]
        fmt.Println("host:port = %s:%s", upstream_host, upstream_port)
    } else {
        upstream_host = "google.com"
        upstream_port = "80"
    }

	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening")
	if err != nil {
		fmt.Println("OH SHIT, couldn't listen!")
	}
	// I will listen to you forever:
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("OH NOES, a conection error!")
			continue
		}
		fmt.Println("Accepted connection")
		go handleConnection(conn)
	}
}
