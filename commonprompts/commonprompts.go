package commonprompts

import (
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

type SelectContent struct {
	Label string
	Items []string
}

type PromptContent struct {
	Label      string
	Validation func(string) error
}

func SelectPrompt(sc SelectContent) (choice string) {
	prompt := promptui.Select{
		Label: sc.Label,
		Items: sc.Items,
	}

	_, choice, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func InputPrompt(pc PromptContent) (result string) {
	prompt := promptui.Prompt{
		Label:       pc.Label,
		Validate:    pc.Validation,
		HideEntered: true,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: %s \n", pc.Label, result)

	return
}

func InputPromptLoop(pc PromptContent, loopTimes string) (inputs []string) {
	numInputs, err := strconv.Atoi(loopTimes)
	if err != nil {
		log.Fatal(err)
	}

	for n := 0; n < numInputs; n++ {
		inputs = append(inputs, InputPrompt(pc))
	}

	return
}
