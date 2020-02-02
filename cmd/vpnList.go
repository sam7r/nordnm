package cmd

import (
	"nordnm/logger"
	"nordnm/nordvpn"

	"github.com/spf13/cobra"
)

// vpnListCmd represents the vpnList command
var vpnListCmd = &cobra.Command{
	Use:   "list",
	Short: "List NordVPN servers",
	Long:  `List all available NordVPN servers`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := nordvpn.GetRecommendations(nordvpn.RecommendationFilters{})
		if err != nil {
			logger.STDout.Errorf("Getting recommendations failed: %v", err)
		}
		logger.STDout.Info(resp)
	},
}

func init() {
	vpnCmd.AddCommand(vpnListCmd)

	// Local flag -- looks for connection files already present within the system
	vpnListCmd.Flags().BoolP("local", "l", false, "List VPN connections stored locally")
}
