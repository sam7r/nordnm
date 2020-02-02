package cmd

import (
	"fmt"
	"nordnm/logger"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var verbose bool
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nordnm",
	Short: "A utility to manage your nord connections using nmcli",
	Long: `Use nordnm to manage your nord vpn connections directly with the Network Manager CLI
	The tool includes commands to:
	- vpn: Interact with the NordVPN API to query available servers
	- conn: Manage, start and stop created Network Manager connections
	- killswitch: Use UFW to enable and disable a killswitch
	- config: manage configuration values
  `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file default is $HOME/.nordnm)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

// initLogger reads in verbose flag to set level of logger
func initLogger() {
	logger.Init(verbose)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".nordnm" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".nordnm")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
