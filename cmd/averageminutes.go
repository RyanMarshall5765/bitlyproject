package cmd

import (
	"fmt"

	"github.com/RyanMarshall5765/bitlyproject/cfg"
	"github.com/RyanMarshall5765/bitlyproject/commonprompts"
	"github.com/RyanMarshall5765/bitlyproject/promptvalidations"
	"github.com/RyanMarshall5765/bitlyproject/utils"
)

func AverageMinutesFlow() {
	selectContent := commonprompts.SelectContent{
		Label: "Please pick an input method",
		Items: []string{"Enter Manually", "Provided File"},
	}
	choice := commonprompts.SelectPrompt(selectContent)

	if choice == "Provided File" {
		fileNameContent := commonprompts.PromptContent{
			Label:      "Please provide the location of the file",
			Validation: promptvalidations.ValidateFile(),
		}
		f := commonprompts.InputPrompt(fileNameContent)

		times := utils.ReadFileByLine(f, cfg.TimeStampPattern)
		fmt.Println(utils.AverageMinutes(times))

	} else if choice == "Typed Manually" {
		numInputsContent := commonprompts.PromptContent{
			Label:      "How many inputs do you have?",
			Validation: promptvalidations.ValidateRange(1, 50),
		}
		numRacers := commonprompts.InputPrompt(numInputsContent)

		manualInputXTimesContent := commonprompts.PromptContent{
			Label:      "Input racer end date",
			Validation: promptvalidations.ValidateStringPattern(cfg.TimeStampPattern),
		}
		times := commonprompts.InputPromptLoop(manualInputXTimesContent, numRacers)

		fmt.Println(utils.AverageMinutes(times))
	}
}
