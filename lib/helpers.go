package lib

import (
	"github.com/manifoldco/promptui"

	"math/rand"
	"time"
)

func AskMore() bool {
	prompt := promptui.Prompt{
		Label:     "More?",
		IsConfirm: true,
	}

	result, _ := prompt.Run()

	return result == "y"
}

func RandomAry(len int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(len)
}