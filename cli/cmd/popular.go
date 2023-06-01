package cmd

import (
	"github.com/Seew0/MeetTuber/api"
	"github.com/spf13/cobra"
)

var popularcmd = &cobra.Command{
	Use: "play-popular",
	Short: "plays a random popular video",
	Long: "this command plays a random popular video ",

	Run: func(cmd *cobra.Command, args []string) {
		GetPopular()
	},
}

func GetPopular(){
	var URL string = api.GetPopularvideoData()
	api.Openbrowser(URL)
}

func init(){
	rootCmd.AddCommand(popularcmd)
}