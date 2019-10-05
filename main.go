package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST = "10.1.99.196"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	fmt.Println(time.Now())
	// Make a buffer to hold incoming data.
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 1000000)
	//var buf []byte
	// Read the incoming connection into the buffer.
	/*length, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}*/
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			fmt.Println("Error, bro:", err)
			break
		}
		//fmt.Println("got", n, "bytes.")
		buf = append(buf, tmp[:n]...)
		//fmt.Printf("I read %d bytes\n", n)
	}
	fmt.Println(time.Now())
	//fmt.Println("Message:" + string(buf))
	//fmt.Println(buf)
	f, err := os.Create("test.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.Write(buf)
	if err != nil {
		fmt.Println(err)
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Closed successfully")
	// Send a response back to person contacting us.
	_, err = conn.Write([]byte("Message received."))
	if err != nil {
		fmt.Println(err)
	}
	// Close the connection when you're done with it.
	err = conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Now())
}