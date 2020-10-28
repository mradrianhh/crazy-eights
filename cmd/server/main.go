package main

import (
	"fmt"
	"net"

	"github.com/mradrianhh/Networker/pkg/network/models"
)

func main() {
}

func receivedUsername(message models.Message, conn net.Conn) error {
	fmt.Println("Got a username!")
	return nil
}
