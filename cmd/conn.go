package cmd

import (
	"github.com/spf13/cobra"
)

// connCmd represents the vpn command
var connCmd = &cobra.Command{
	Use:   "conn",
	Short: "Use this command to interact with the NetworkManager CLI",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(connCmd)
}
