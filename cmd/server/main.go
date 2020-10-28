package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/Navigator/pkg/navigation"
	navmodels "github.com/mradrianhh/Navigator/pkg/navigation/models"
	"github.com/mradrianhh/Networker/pkg/network"
	netmodels "github.com/mradrianhh/Networker/pkg/network/models"
	serverpkg "github.com/mradrianhh/crazy-eights/internal/server"
	"github.com/mradrianhh/crazy-eights/internal/server/screens"
)

var navigator navigation.Navigator
var state navmodels.State
var server network.Server

func main() {
	setup()
	loop()
}

func setup() {
	server = network.NewServer("tcp", "0.0.0.0:1200")
	server.AddHandler(netmodels.USERNAME, receivedUsername)
	go server.Listen()

	navigator = navigation.NewNavigator()

	navigator.AddScreen(screens.GetMainMenuScreen().Identifier, screens.GetMainMenuScreen())

	state.State = serverpkg.MAIN_MENU
}

func loop() {
	for {
		switch state.State {
		case serverpkg.MAIN_MENU:
			err := navigator.RemoveScreensAndPush(&state, screens.MAIN_MENU_SCREEN)
			checkError(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func receivedUsername(request netmodels.Request, conn net.Conn) error {
	fmt.Println("Got a username!")
	return nil
}
