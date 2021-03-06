package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "4339"
	CONN_TYPE = "tcp"
)

type Peer struct {
	Name      string
	TcpSocket net.Conn
	PeerList  []Peer
	isLeader  bool
}

type Room struct {
	PeerList []*Peer // current peers in room
	RoomNo   int
}

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
		fmt.Println(conn.RemoteAddr().String())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	_ = reqLen
	fmt.Println(string(buf))
	var peerList []Peer
	newPeer := Peer{"test1", conn, peerList, false}
	peerList = append(peerList, newPeer)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	encoder := gob.NewEncoder(conn)
	encoder.Encode(newPeer)
	//conn.Write(peerList)

	// Close the connection when you're done with it.
	conn.Close()
}

