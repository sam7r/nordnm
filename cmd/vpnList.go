package cmd

import (
	"nordnm/logger"
	"nordnm/nordvpn"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var listFlagFilters nordvpn.RecommendationFilters

// vpnListCmd represents the vpnList command
var vpnListCmd = &cobra.Command{
	Use:   "list",
	Short: "List NordVPN servers",
	Long:  `List all available NordVPN servers`,
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := nordvpn.GetRecommendations(listFlagFilters)
		if err != nil {
			logger.Stdout.Errorf("Getting recommendations failed: %v", err)
		}

		// create new table for output
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"", "ID", "NAME", "HOSTNAME", "LOAD", "TECHNOLOGIES"})
		for number, recommendation := range resp {

			// format list of technologies
			technologies := make([]string, 0, len(recommendation.Technologies))
			for _, tech := range recommendation.Technologies {
				technologies = append(technologies, tech.Name)
			}

			// create output rows
			t.AppendRow([]interface{}{
				number + 1,
				recommendation.ID,
				recommendation.Name,
				recommendation.Hostname,
				recommendation.Load,
				strings.Join(technologies, "\n"),
			})
		}
		t.Render()

	},
}

func init() {
	vpnCmd.AddCommand(vpnListCmd)

	vpnListCmd.Flags().StringVarP(&listFlagFilters.ServerGroupID, "group", "g", "", "Server group ID i.e legacy_double_vpn")
	vpnListCmd.Flags().Uint8VarP(&listFlagFilters.CountryID, "country", "c", 0, "Country ID i.e. 227 (GB)")
	vpnListCmd.Flags().StringVarP(&listFlagFilters.TechnologyID, "technology", "t", "", "Technology identifier i.e. openvpn_udp")
	vpnListCmd.Flags().Uint8VarP(&listFlagFilters.Limit, "limit", "l", 10, "Limit the number of results returned")
}
