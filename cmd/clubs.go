package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// clubsCmd represents the clubs command
var clubsCmd = &cobra.Command{
	Use:     "clubs",
	Short:   "clubs",
	Example: "futbin clubs --league 13",
	Run: func(cmd *cobra.Command, args []string) {
		leagues, err := leaguesAndClubsClient.Get()
		cobra.CheckErr(err)
		data := make([][]string, 1)
		for _, league := range leagues {
			if flagLeagueID != 0 && flagLeagueID != league.ID {
				continue
			}
			for _, club := range league.Clubs {
				data = append(data, []string{
					strconv.Itoa(club.ID),
					club.Name,
					fmt.Sprintf("%s (%d)", league.Name, league.ID),
				})
			}
		}
		printTable([]string{"ID", "NAME", "LEAGUE (ID)"}, data)
	},
}

func init() {
	rootCmd.AddCommand(clubsCmd)
	clubsCmd.PersistentFlags().IntVar(&flagLeagueID, "league", 0, "League ID")
}
