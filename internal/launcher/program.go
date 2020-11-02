package launcher

import (
	"fmt"
	"os"

	"github.com/mradrianhh/Navigator/pkg/navigation"
	"github.com/mradrianhh/Navigator/pkg/navigation/models"
	"github.com/mradrianhh/crazy-eights/internal/launcher/screens.go"
)

var state models.State
var navigator navigation.Navigator

// Run the launcher-program.
func Run() {
	fmt.Println("Running launcher")
	setup()
	loop()
}

// Setup the launcher-program.
func setup() {
	navigator = navigation.NewNavigator()
	state.State = "MAIN_MENU"

	navigator.AddScreen(screens.GetMainMenuScreen().Identifier, screens.GetMainMenuScreen())
}

// Launcher main-loop.
func loop() {
	for {
		switch state.State {
		case "MAIN_MENU":
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
