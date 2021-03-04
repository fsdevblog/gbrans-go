package main

import (
	"fmt"
	"gbrains-go/lesson2"
	"gbrains-go/lib"
	"os"

	"github.com/manifoldco/promptui"
)

const (
	LessonSecond = "Second lesson"
	Exit         = "Exit"
	Back         = "Back"
)

func selectPrompt(label string, items interface{}) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func catchPromptError(err error) {
	fmt.Printf("Prompt failed: %v\n", err)
	os.Exit(1)
}

func main() {
	for { // Lessons loop
		lessonQ, err := selectPrompt("Choose the lesson", []string{LessonSecond, Exit})

		if err != nil {
			catchPromptError(err)
		}

		if lessonQ == LessonSecond {
			secondLessonTasks()
		} else if lessonQ == Exit {
			break
		}
	}
}

func secondLessonTasks() {
	taskQ, err := selectPrompt("Choose the task", []string{"1", "2", "3", Back})

	if err != nil {
		catchPromptError(err)
	}

	taskLoop := true
	for taskLoop { // tasks loop
		switch taskQ {
		case "1":
			fmt.Println("Rectangle area. Enter the lengths of the sides of the rectangle")
			lesson2.RectangleArea()
		case "2":
			fmt.Println("Circumference and diameter. Please input area of a circle")
			lesson2.DiAndLengthOfCircle()
		case "3":
			fmt.Println("Digits to the words")
			lesson2.DigitsToWords()
		case Back:
			taskLoop = false
		}

		if taskQ != Back && !lib.AskMore() {
			taskLoop = false
		}
	}
}
