package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "mtuber",
	Short: "MeetTuber is a cli bot to play videos written in golang",
	Long: "MeetTuber is the best bot in the world",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey i am MeetTuber how are you today?")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}