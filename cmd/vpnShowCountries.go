package cmd

import (
	"github.com/sam7r/nordnm/logger"
	"github.com/sam7r/nordnm/nordvpn"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// vpnShowCountriesCmd represents the vpn command
var vpnShowCountriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Show available countries",
	Long:  `Show all available countries, these can be used to target servers via the list command using --country=ID`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := nordvpn.GetCountries()
		if err != nil {
			logger.Stdout.Errorf("Getting countries failed: %v", err)
		}

		// create new table for output
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "NAME", "CODE"})
		for _, country := range resp {

			// create output rows
			t.AppendRow([]interface{}{
				country.ID,
				country.Name,
				country.Code,
			})
		}
		t.Render()

	},
}

func init() {
	vpnShowCmd.AddCommand(vpnShowCountriesCmd)
}
