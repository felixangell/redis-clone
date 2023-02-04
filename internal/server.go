package internal

import (
	"fmt"
	"github.com/bat-labs/krake/pkg/parse"
	"log"
	"net"
)

type Krake struct {
}

func NewKrakeServer() *Krake {
	return &Krake{}
}

func (k *Krake) ListenAndServe(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			for {
				buf := make([]byte, 1024)
				bufLen, err := conn.Read(buf)
				if err != nil {
					fmt.Printf("Error reading: %#v\n", err)
					return
				}

				response := parse.ParseMessage(buf[:bufLen])
				log.Println("<", response)
				_, err = conn.Write(response.Serialize())
				if err != nil {
					log.Println(err.Error())
				}
			}
		}(conn)
	}
}

func (k *Krake) Close() {
}
