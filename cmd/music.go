package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	musicStyle    string
	musicTitle    string
	musicEmail    string
	musicVocal    string
	musicAutoMode bool
)

var musicCmd = &cobra.Command{
	Use:   "music [lyrics_or_theme]",
	Short: "Generate original music using Suno AI",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lyricsPrompt := args[0]
		fmt.Printf("Generating music: %s ...\n", musicTitle)

		client := GetClient()
		body := map[string]interface{}{
			"stylePrompt":  musicStyle,
			"lyricsPrompt": lyricsPrompt,
			"title":        musicTitle,
			"email":        musicEmail,
			"isAutoMode":   musicAutoMode,
		}
		if musicVocal != "" {
			body["vocalGender"] = musicVocal
		}

		resp, err := client.R().
			SetBody(body).
			SetResult(&MusicResponse{}).
			Post("/api/skill/music/create")

		if err != nil {
			fmt.Printf("Error: Music generation failed: %v\n", err)
			return
		}

		if resp.IsError() {
			fmt.Printf("Error: Music API responded with status %d: %s\n", resp.StatusCode(), resp.String())
			return
		}

		result := resp.Result().(*MusicResponse)
		fmt.Println("-------------------------------------------")
		fmt.Printf("Success! Song: %s\n", musicTitle)
		fmt.Printf("Cost: %.0f Credits\n", result.Cost)
		fmt.Println("-------------------------------------------")
	},
}

var musicHealthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check Suno API health status",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetClient()
		resp, err := client.R().
			Get("/api/skill/music/health")

		if err != nil {
			fmt.Printf("Error: Health check failed: %v\n", err)
			return
		}

		if resp.IsError() {
			fmt.Printf("Error: Health API responded with status %d: %s\n", resp.StatusCode(), resp.String())
			return
		}

		fmt.Println("-------------------------------------------")
		fmt.Println("Music service is healthy")
		fmt.Println(resp.String())
		fmt.Println("-------------------------------------------")
	},
}

type MusicResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Cost    float64     `json:"cost"`
}

func init() {
	musicCmd.Flags().StringVarP(&musicStyle, "style", "s", "", "Music genre, mood, BPM (e.g. \"K-Pop, Dance, Upbeat\")")
	musicCmd.Flags().StringVarP(&musicTitle, "title", "t", "", "Song title")
	musicCmd.Flags().StringVarP(&musicEmail, "email", "e", "", "Email to receive the result")
	musicCmd.Flags().StringVarP(&musicVocal, "vocal", "v", "", "Vocal gender (Male, Female, or empty)")
	musicCmd.Flags().BoolVar(&musicAutoMode, "auto", true, "Auto lyrics generation")

	musicCmd.MarkFlagRequired("style")
	musicCmd.MarkFlagRequired("title")
	musicCmd.MarkFlagRequired("email")

	musicCmd.AddCommand(musicHealthCmd)
	RootCmd.AddCommand(musicCmd)
}
