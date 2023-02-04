package exec

import (
	"log"
	"net"
)

type KafkaNodeOrchestrator struct {
}

func (o *KafkaNodeOrchestrator) Submit(conn net.Conn, command Command) {
	response := command.Execute(o)
	_, err := conn.Write(response.Serialize())
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("<", response)
}
