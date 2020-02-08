package cmd

import (
	"fmt"
	"nordnm/logger"
	"nordnm/nmcli"
	"nordnm/nordvpn"
	"nordnm/utils"
	"strings"
	"time"

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
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fetch ovpn file from nord cdn
		hostID = strings.ToLower(hostID)
		techID = strings.ToLower(techID)
		file, err := nordvpn.GetNordVpnConfigFile(hostID, techID)
		if err != nil {
			logger.Stdout.Fatalf("Fetching nord ovpn file failed: %+v, aborting command", err)
		}

		// create connection id and save to temp
		msec := time.Now().UnixNano() / 1000000
		connectionID := fmt.Sprintf("%s.nordvpn.com.%s.%d", hostID, techID, msec)
		filepath := fmt.Sprintf("/tmp/%s.ovpn", connectionID)
		logger.Stdout.Infof("Saving temp file to %s", filepath)
		err = utils.SaveFileToTempLocation(filepath, file, 0664)
		if err != nil {
			logger.Stdout.Fatalf("Saving temp file at location %s failed: %+v", filepath, err)
		}

		// nmcli connection import
		if out, err := nmcli.ImportOvpnConnection(filepath); err != nil {
			logger.Stdout.Fatalf("Importing OVPN file to NetworkManager failed: %s", err)
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
			logger.Stdout.Fatalf("Modifying OVPN file to NetworkManager failed: %s", err)
		} else {
			logger.Stdout.Infof("Modified OVPN file: %v", out)
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
