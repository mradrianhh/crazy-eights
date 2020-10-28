package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mradrianhh/Navigator/pkg/navigation"
	navmodels "github.com/mradrianhh/Navigator/pkg/navigation/models"
	"github.com/mradrianhh/Networker/pkg/network"
	"github.com/mradrianhh/Networker/pkg/network/models"
	clientpkg "github.com/mradrianhh/crazy-eights/internal/client"
	"github.com/mradrianhh/crazy-eights/internal/client/screens"
)

var navigator navigation.Navigator
var state navmodels.State
var client network.Client

func main() {
	setup()
	loop()
}

func setup() {
	client = network.NewClient("tcp", "0.0.0.0:1200")

	navigator = navigation.NewNavigator()

	navigator.AddScreen(screens.GetMainMenuScreen().Identifier, screens.GetMainMenuScreen())

	state.State = clientpkg.MAIN_MENU
}

func loop() {
	for {
		switch state.State {
		case clientpkg.MAIN_MENU:
			err := navigator.RemoveScreensAndPush(&state, screens.MAIN_MENU_SCREEN)
			checkError(err)
		case clientpkg.REQUEST:
			response, _ := client.Request(models.NewRequest(models.USERNAME))
			fmt.Println(response.ResponseCode)
			time.Sleep(1 * time.Second)
			state.State = clientpkg.MAIN_MENU
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
