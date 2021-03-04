package lib

import (
	"github.com/manifoldco/promptui"
)

func AskMore() bool {
	prompt := promptui.Prompt{
		Label:     "More?",
		IsConfirm: true,
	}

	result, _ := prompt.Run()

	return result == "y"
}
