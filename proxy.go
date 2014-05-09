/*
 Simple Transparent Proxy in Go

*/

package main

import "fmt"
import "io"
import "net"


func handleUpstream(conn net.Conn, upstream net.Conn) {
	// Start listening to upstream
	io.Copy(upstream, conn)
}


func handleConnection(conn net.Conn) {
	upstream, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Couldn't connect to upstream")
	}
	go handleUpstream(conn, upstream)
	io.Copy(conn, upstream)
}


func main() {
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
