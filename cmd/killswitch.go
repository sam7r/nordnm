package cmd

import (
	"fmt"

	"github.com/sam7r/nordnm/logger"
	"github.com/sam7r/nordnm/ufw"
	"github.com/spf13/cobra"
)

var dryRun bool

// killswitchCmd represents the killswitch command
var killswitchCmd = &cobra.Command{
	Use:   "killswitch",
	Short: "Use this command to change the UFW rules",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(killswitchCmd)
	killswitchCmd.AddCommand(killswitchEnableCmd)
	killswitchCmd.AddCommand(killswitchDisableCmd)

	killswitchEnableCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Runs UFW rule set changes in dry run mode")
	killswitchDisableCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Runs UFW rule set changes in dry run mode")
}

var killswitchEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable UFW rules acting as a killswitch",
	Long:  "Enables predefined rules to block io on default rule set and only allow out via tun0",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := ufw.EnableKillswitch(dryRun)
		if err != nil {
			logger.Stdout.Infof("An error occured enabling the killswitch: %e", err)
		}
		fmt.Printf("%s", out)
	},
}

var killswitchDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Removes UFW rules acting as a killswitch",
	Long:  "Removed predefined rules to block io on default rule set and only allow out via tun0",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := ufw.DisableKillswitch(dryRun)
		if err != nil {
			logger.Stdout.Infof("An error occured removing the killswitch: %e", err)
		}
		fmt.Println(out)
	},
}
