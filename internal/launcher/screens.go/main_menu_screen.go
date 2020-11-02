package screens

import (
	"fmt"
	"os"

	"github.com/mradrianhh/Navigator/pkg/navigation"
	"github.com/mradrianhh/Navigator/pkg/navigation/models"
)

// MainMenuScreen ...
type MainMenuScreen struct {
	Identifier string
}

var mainMenuScreen = MainMenuScreen{Identifier: MAIN_MENU_SCREEN}

// GetMainMenuScreen ...
func GetMainMenuScreen() *MainMenuScreen {
	return &mainMenuScreen
}

// Show ...
func (mainMenuScreen MainMenuScreen) Show(state *models.State) error {
	fmt.Print("Launcher Home\n\n0 - Exit\n\n")

	var response int
	_, err := fmt.Scan(&response)
	if err != nil {
		return err
	}

	switch response {
	case 0:
		navigation.Clear()
		os.Exit(0)
	default:
		state.State = "MAIN_MENU"
	}

	return nil
}
