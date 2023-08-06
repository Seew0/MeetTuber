package cmd

import (
	"log"

	"github.com/Seew0/MeetTuber/bot"
	"github.com/spf13/cobra"
)
var joinCMD = &cobra.Command{
	Use: "join",
	Run: func(cmd *cobra.Command, args []string) {
		JoinFunc(args)
	},
}

func JoinFunc(args []string){
	url := args[1]
	switch meetOption := args[0];meetOption{
	case "jitsi-meet":
		bot.JoinJitsiMeet(url)
	default:
		log.Fatalln(meetOption + "is not supported the bot works only with google-meet and jitsi-meet")
	}
}
func init(){
	rootCmd.AddCommand(joinCMD)
}