package internal

import (
	"github.com/bat-labs/krake/pkg/cmd"
	"log"
	"net"
)

type KafkaNodeOrchestrator struct {
}

func (o KafkaNodeOrchestrator) submit(conn net.Conn, command cmd.Command) {
	response := command.Execute()
	_, err := conn.Write(response.Serialize())
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("<", response)
}
