package cmd

import (
	"github.com/sam7r/nordnm/nordvpn"
	"github.com/sam7r/nordnm/utils"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// vpnListCmd represents the vpnList command
var vpnListCmd = &cobra.Command{
	Use:   "list",
	Short: "List NordVPN servers",
	Long:  `List all available NordVPN servers`,
	Run: func(cmd *cobra.Command, args []string) {

		flagFilters := nordvpn.RecommendationFilters{
			CountryID:     uint8(viper.GetInt("preferences.countryCode")),
			TechnologyID:  viper.GetString("preferences.technologyIdentifier"),
			ServerGroupID: viper.GetString("preferences.groupIdentifier"),
			Limit:         uint8(viper.GetInt("preferences.limit")),
		}

		resp, err := nordvpn.GetRecommendations(flagFilters)
		if err != nil {
			utils.Logger.Errorf("Getting recommendations failed: %v", err)
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

	vpnListCmd.Flags().StringP("group", "g", "", "Server group ID i.e legacy_double_vpn")
	vpnListCmd.Flags().Uint8P("country", "c", 0, "Country ID i.e. 227 (GB)")
	vpnListCmd.Flags().StringP("technology", "t", "", "Technology identifier i.e. openvpn_udp")
	vpnListCmd.Flags().Uint8P("limit", "l", 10, "Limit the number of results returned")

	viper.BindPFlag("preferences.countryCode", vpnListCmd.Flag("country"))
	viper.BindPFlag("preferences.groupIdentifier", vpnListCmd.Flag("group"))
	viper.BindPFlag("preferences.technologyIdentifier", vpnListCmd.Flag("technology"))
	viper.BindPFlag("preferences.limit", vpnListCmd.Flag("limit"))
}
