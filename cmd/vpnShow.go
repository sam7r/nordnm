package cmd

import (
	"github.com/spf13/cobra"
)

// vpnShowCmd represents the vpn command
var vpnShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show ids for nord vpn types",
	Long:  `Use this command with country, group or tech to get a list of ids, these will help you configure defaults or for use with the 'vpn list' command`,
}

func init() {
	vpnCmd.AddCommand(vpnShowCmd)
}
