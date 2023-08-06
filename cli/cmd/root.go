package cmd

import (
	"fmt"
	"os"

	"github.com/Seew0/MeetTuber/api"
	"github.com/spf13/cobra"
)
var greet string = api.Getrandomjoke()
var rootCmd = cobra.Command{
	Use: "mtuber",
	Short: "MeetTuber is a cli bot that can join meets and do wonders for you",
	Long: "MeetTuber is the best bot in the world",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(greet)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}