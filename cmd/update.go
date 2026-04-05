package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/creativeprojects/go-selfupdate"
	"github.com/spf13/cobra"
)

const (
	githubOwner = "nadmzda"
	githubRepo  = "fast-claw-cli"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update FastClaw CLI to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Current version: %s\n", version)
		fmt.Println("Checking for updates...")

		ctx := context.Background()
		slug := selfupdate.ParseSlug(githubOwner + "/" + githubRepo)

		latest, found, err := selfupdate.DetectLatest(ctx, slug)
		if err != nil {
			fmt.Printf("Error: Failed to check for updates: %v\n", err)
			os.Exit(1)
		}

		if !found {
			fmt.Println("Error: Repository or release not found")
			os.Exit(1)
		}

		if latest.GreaterThan(version) {
			fmt.Printf("New version available: %s (released: %s)\n", latest.Version(), latest.PublishedAt)
			fmt.Printf("Release notes: %s\n", latest.ReleaseNotes)
			fmt.Println("Downloading and installing update...")

			release, err := selfupdate.UpdateSelf(ctx, version, slug)
			if err != nil {
				fmt.Printf("Error: Failed to update: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Successfully updated to version %s\n", release.Version())
		} else {
			fmt.Println("You are already running the latest version.")
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
