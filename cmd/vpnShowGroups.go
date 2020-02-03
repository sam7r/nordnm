package cmd

import (
	"nordnm/logger"
	"nordnm/nordvpn"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// vpnShowGroupsCmd represents the vpn command
var vpnShowGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Show available groups",
	Long:  `Show all available groups, these can be used to target specific servers via the list command using --group={IDENTIFIER}`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := nordvpn.GetGroups()
		if err != nil {
			logger.Stdout.Errorf("Getting groups failed: %v", err)
		}

		// create new table for output
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "NAME", "IDENTIFIER"})
		for _, group := range resp {

			// create output rows
			t.AppendRow([]interface{}{
				group.ID,
				group.Title,
				group.Identifier,
			})
		}
		t.Render()

	},
}

func init() {
	vpnShowCmd.AddCommand(vpnShowGroupsCmd)
}
