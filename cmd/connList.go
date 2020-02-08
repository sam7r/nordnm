package cmd

import (
	"nordnm/logger"
	"nordnm/nmcli"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var showActiveConn bool
var showAllConn bool

// connListCmd represents the vpn command
var connListCmd = &cobra.Command{
	Use:   "list",
	Short: "List managed VPN connections",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := nmcli.ListConnections(showActiveConn)
		if showAllConn != true {
			out.FilterByType("vpn")
			logger.Stdout.Info(out)
		}

		if err != nil {
			logger.Stdout.Infof("Showing connections failed: %v", err)
		}

		// TODO: add active state to the table

		// create new table for output
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"NAME", "UUID"})
		for _, conn := range out {
			// create output rows
			t.AppendRow([]interface{}{
				conn.Name,
				conn.UUID,
			})
		}
		t.Render()
	},
}

func init() {
	connCmd.AddCommand(connListCmd)
	connListCmd.Flags().BoolVarP(&showActiveConn, "active", "", false, "show only active connections")
	connListCmd.Flags().BoolVarP(&showAllConn, "all", "", false, "show all connections, not just vpn")

}
