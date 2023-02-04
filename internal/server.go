package internal

import (
	"fmt"
	"github.com/bat-labs/krake/pkg/cmd"
	"github.com/bat-labs/krake/pkg/parse"
	"net"
)

type Krake struct {
	orchestrator *KafkaNodeOrchestrator
}

func NewKrakeServer() *Krake {
	return &Krake{
		&KafkaNodeOrchestrator{},
	}
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

				value := parse.ParseMessage(buf[:bufLen])

				command, err := cmd.ParseCommand(value)
				if err != nil {
					panic(err)
				}

				k.orchestrator.submit(conn, command)
			}
		}(conn)
	}
}

func (k *Krake) Close() {
}
