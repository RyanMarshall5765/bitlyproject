package cmd

import (
	"fmt"
	"os"

	"github.com/RyanMarshall5765/bitlyproject/commonprompts"
	"github.com/RyanMarshall5765/bitlyproject/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bitlyproject",
	Short: "Program for a boat race",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Average Minutes: Returns the average of all racers finish time.
			-- Enter Manually: Asks for how many racers you wish to input. Then the inputs in the format hh:mm xM, DAY n
			-- Provide File: Takes a text file where each input is on a new line with the format hh:mm xM, DAY n
Format Explanation: hh:mm xM, DAY n
hh is the hour 01-12 , mm is the minute 0-59, x is P or A for AM/PM, and n is the day they finished 1-99.`)
		PickProgram()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func PickProgram() {
	const AVERAGEMINUTES = "Average Minutes"
	selectContent := commonprompts.SelectContent{
		Label: "Please pick a program",
		Items: []string{AVERAGEMINUTES},
	}
	choice := commonprompts.SelectPrompt(selectContent)

	if choice == AVERAGEMINUTES {
		utils.AverageMinutesFlow()
	}
}
