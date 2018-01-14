// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"unicode"
)

// UDPServer implements the UDP server specification found at
// https://github.com/uis-dat520/labs/blob/master/lab2/README.md#udp-server
type UDPServer struct {
	conn *net.UDPConn
	// TODO(student): Add fields if needed
}

type rot13Reader struct {
	r io.Reader
}

// NewUDPServer returns a new UDPServer listening on addr. It should return an
// error if there was any problem resolving or listening on the provided addr.
func NewUDPServer(addr string) (*UDPServer, error) {
	// TODO(student): Implement
	svr := &UDPServer{conn: nil}
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	checkError(err)
	svr.conn, err = net.ListenUDP("udp", udpAddr)
	checkError(err)
	return svr, nil
}

// ServeUDP starts the UDP server's read loop. The server should read from its
// listening socket and handle incoming client requests as according to the
// the specification.
func (u *UDPServer) ServeUDP() {
	// TODO(student): Implement
	buf := make([]byte, 1024)

	for {
		n, add, err := u.conn.ReadFromUDP(buf)
		msg := strings.Split(string(buf[:n]), "|:|")
		res := ""
		if msg[1] == "" {
			res = "Unknown command"
		} else {
			switch msg[0] {
			case "UPPER":
				res = Upper(msg[1])
			case "LOWER":
				res = Lower(msg[1])
			case "CAMEL":
				res = CamelCase(msg[1])
			case "ROT13":
				res = Rot13(msg[1])
			case "SWAP":
				res = Swap(msg[1])
			default:
				res = "Unknown command"

			}
		}

		_, err = u.conn.WriteToUDP([]byte(res), add)
		checkError(err)
	}
}

//coverting to uppercase
func Upper(src string) string {
	src = strings.ToUpper(src)
	return src

}

//converting to lowercase
func Lower(src string) string {
	src = strings.ToLower(src)
	return src

}

//coverting to camelcase
func CamelCase(src string) string {
	src = strings.ToLower(src)
	src = strings.Title(src)
	return src
}

//converting rot13
func rotvalue(r rune) rune {
	if r >= 'a' && r <= 'z' {
		// Rotate lowercase letters 13 places.
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' {
		// Rotate uppercase letters 13 places.
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	// Do nothing.
	return r
}

//converting rot13
func Rot13(s string) string {
	return strings.Map(rotvalue, s)
}

func swapvalue(r rune) rune {
	switch {
	case unicode.IsUpper(r):
		return unicode.ToLower(r)
	case unicode.IsLower(r):
		return unicode.ToUpper(r)
	}
	return r
}

//swap strings
func Swap(s string) string {
	return strings.Map(swapvalue, s)
}

// socketIsClosed is a helper method to check if a listening socket has been
// closed.
func socketIsClosed(err error) bool {
	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}
	return false
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error, Something went wrong")
	}
}
