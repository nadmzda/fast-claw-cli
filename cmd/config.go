package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure FastClaw CLI settings",
}

var setKeyCmd = &cobra.Command{
	Use:   "set-key [API_KEY]",
	Short: "Save FastClaw API Key to local config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := args[0]
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error: Could not find home directory: %v\n", err)
			return
		}

		configPath := filepath.Join(home, ".fastclaw_config")
		err = os.WriteFile(configPath, []byte(apiKey), 0600)
		if err != nil {
			fmt.Printf("Error: Could not save config file: %v\n", err)
			return
		}

		fmt.Printf("Success! API Key saved to %s\n", configPath)
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setKeyCmd)
}
