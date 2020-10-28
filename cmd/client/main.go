package main

import (
	"fmt"

	"github.com/mradrianhh/Networker/pkg/network"
	"github.com/mradrianhh/Networker/pkg/network/models"
)

func main() {
	fmt.Println("Client interface")
	c := network.NewClient("tcp", "0.0.0.0:1200")
	if response, err := c.Request(models.NewRequest(models.USERNAME)); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(response.ResponseCode())
	}
}
