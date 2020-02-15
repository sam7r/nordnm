package cmd

import (
	"github.com/sam7r/nordnm/logger"
	"github.com/sam7r/nordnm/nordvpn"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// vpnShowTechCmd represents the vpn command
var vpnShowTechCmd = &cobra.Command{
	Use:   "tech",
	Short: "Show available technologies",
	Long:  `Show all available technologies, these can be used to target servers via the list command using --technology=IDENTIFIER`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := nordvpn.GetTechnologies()
		if err != nil {
			logger.Stdout.Errorf("Getting technologies failed: %v", err)
		}

		// create new table for output
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "NAME", "IDENTIFIER"})
		for _, technology := range resp {

			// create output rows
			t.AppendRow([]interface{}{
				technology.ID,
				technology.Name,
				technology.Identifier,
			})
		}
		t.Render()

	},
}

func init() {
	vpnShowCmd.AddCommand(vpnShowTechCmd)
}
