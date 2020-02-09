package cmd

import (
	"fmt"
	"nordnm/logger"
	"nordnm/nmcli"
	"nordnm/nordvpn"
	"nordnm/utils"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var hostID string
var techID string
var username string
var password string
var dns string

// connCreateCmd represents the vpn command
var connCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new connection",
	Long:  `Create a new connection within NetworkManager, this will only add and not activate the connection, additional settings will be taken from your config file or can be passed in to override`,
	Run: func(cmd *cobra.Command, args []string) {
		// fetch ovpn file from nord cdn
		hostID = strings.ToLower(hostID)
		techID = strings.ToLower(techID)

		// check if host-id already exists?
		networkConnections, err := nmcli.ListConnections(false)
		for _, nmconn := range networkConnections {
			if strings.Contains(nmconn.Name, fmt.Sprintf("%s.", hostID)) {
				fmt.Println("Existing connection found for given host id, aborting")
				logger.Stdout.Infof("Connection for %s already exists", nmconn)
				os.Exit(1)
			}
		}

		// fetching ovpn file
		file, err := nordvpn.GetNordVpnConfigFile(hostID, techID)
		if err != nil {
			fmt.Println("Fetching ovpn file failed, aborting")
			logger.Stdout.Infof("Fetching nord ovpn file failed: %+v, aborting command", err)
			os.Exit(1)
		}

		// create connection id and save to temp
		msec := time.Now().UnixNano() / 1000000
		connectionID := fmt.Sprintf("%s.nordvpn.com.%s.%d", hostID, techID, msec)
		filepath := fmt.Sprintf("/tmp/%s.ovpn", connectionID)
		logger.Stdout.Infof("Saving temp file to %s", filepath)
		err = utils.SaveFileToTempLocation(filepath, file, 0664)
		if err != nil {
			fmt.Println("Temp file save failure, aborting")
			logger.Stdout.Infof("Saving temp file at location %s failed: %+v", filepath, err)
			os.Exit(1)
		}

		// nmcli connection import
		if out, err := nmcli.ImportOvpnConnection(filepath); err != nil {
			fmt.Println("NetworkManager import failure, aborting")
			logger.Stdout.Errorf("Importing OVPN file to NetworkManager failed: %s", err)
			os.Exit(1)
		} else {
			logger.Stdout.Infof("Imported OVPN file: %v", out)
		}

		// nmcli modify connection
		authmode := "encrypted"
		if username != "" && password != "" {
			authmode = "non_encrypted"
		}
		connectionSettings := nmcli.OvpnConnectionDefaults{
			DNS:        dns,
			IgnoreIPV6: true,
			AuthSettings: nmcli.Auth{
				Mode: authmode,
				User: username,
				Pass: password,
			},
		}

		if out, err := nmcli.ModifyConnection(connectionID, connectionSettings); err != nil {
			fmt.Print("A connection was created but unable to modify the details, aborting")
			logger.Stdout.Infof("Modifying OVPN file to NetworkManager failed: %s", err)
			os.Exit(1)
		} else {
			fmt.Println("Connection successfully created")
			logger.Stdout.Infof("Modified OVPN file: %v", out)
			newNetworkConnections, _ := nmcli.ListConnections(false)
			var newConnection nmcli.NetworkConnection
			for _, newnmconn := range newNetworkConnections {
				for _, existingnmconn := range networkConnections {
					if existingnmconn.Name != newnmconn.Name {
						newConnection = newnmconn
					}
				}
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"NAME", "UUID"})
			t.AppendRow([]interface{}{
				newConnection.Name,
				newConnection.UUID,
			})
			t.Render()
		}
	},
}

func init() {
	connCmd.AddCommand(connCreateCmd)
	connCreateCmd.Flags().StringVarP(&hostID, "host", "", "", "The nord vpn host id eg. uk1212 (Required)")
	connCreateCmd.Flags().StringVarP(&techID, "tech", "", "UDP", "The vpn connection technology to use, either TCP or UDP")
	connCreateCmd.MarkFlagRequired("host")

	connCreateCmd.Flags().StringVarP(&username, "username", "", "", "Manually set NordVPN username")
	connCreateCmd.Flags().StringVarP(&password, "password", "", "", "Manually set NordVPN password")

	connCreateCmd.Flags().StringVarP(&dns, "dns", "", "103.86.99.100,103.86.96.100", "Manually set dns address eg. 1.1.1.1,1.0.0.1")
}
