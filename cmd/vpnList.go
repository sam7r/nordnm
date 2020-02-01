package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nordnm/nordvpn"
)

// vpnListCmd represents the vpnList command
var vpnListCmd = &cobra.Command{
	Use:   "list",
	Short: "List NordVPN servers",
	Long:  `List all available NordVPN servers`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := nordvpn.GetRecommendations(nordvpn.RecommendationFilters{})
		if err != nil {
			// log full error to file
			// print and panic with user friendly error message
		}

		fmt.Println(resp)
	},
}

func init() {
	vpnCmd.AddCommand(vpnListCmd)

	// Local flag -- looks for connection files already present within the system
	vpnListCmd.Flags().BoolP("local", "l", false, "List VPN connections stored locally")
}
