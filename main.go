package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

const (
	CONN_TYPE = "tcp"
)

var CONN_HOST = ""
var CONN_PORT = "8080"

func main() {
	// Parse command line args
	parseArgs(os.Args[1:])
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

// Parse arguments
func parseArgs(args []string) {
	if len(args) == 0 || (len(args) >= 1 && args[0] == "-h") {
		printHelp()
		os.Exit(0)
	}
	if len(args) >= 1 {
		CONN_HOST = args[0]
	}
	if len(args) >= 2 {
		CONN_PORT = args[1]
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("go run main.go [flags] addr [port]")
	fmt.Println("Flags:")
	fmt.Println("\t-h: prints this message")
	fmt.Println("Arguments:")
	fmt.Println("\taddr - ip address for server")
	fmt.Println("\tport - port for server (default is 8080")
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	fmt.Println(time.Now())
	// Make a buffer to hold incoming data.
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 1000000)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			fmt.Println("Error, bro:", err)
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	fmt.Println(time.Now())
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
