package cmd

import (
	"github.com/spf13/cobra"
)

// vpnCmd represents the vpn command
var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "Use this command to interact with the NordVPN web API",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(vpnCmd)
}
