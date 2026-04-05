package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	imagePrompt      string
	imageAspectRatio string
	imageResolution  string
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image [prompt]",
	Short: "Generate high-quality images from text prompts using Gemini",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]
		fmt.Printf("Generating image for: %s ...\n", prompt)

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"prompt":      prompt,
				"aspectRatio": imageAspectRatio,
				"resolution":  imageResolution,
			}).
			SetResult(&ImageResponse{}).
			Post("/api/skill/image/generate")

		if err != nil {
			fmt.Printf("Error: Image generation failed: %v\n", err)
			return
		}

		if resp.IsError() {
			fmt.Printf("Error: Image API responded with status %d: %s\n", resp.StatusCode(), resp.String())
			return
		}

		result := resp.Result().(*ImageResponse)
		fmt.Println("-------------------------------------------")
		fmt.Printf("Success! Image URL: %s\n", result.PublicUrl)
		fmt.Printf("Prompt used: %s\n", result.Prompt)
		fmt.Printf("Cost: %f Credits\n", result.Cost)
		fmt.Println("-------------------------------------------")
	},
}

type ImageResponse struct {
	PublicUrl string  `json:"publicUrl"`
	Prompt    string  `json:"prompt"`
	Cost      float64 `json:"cost"`
}

func init() {
	imageCmd.Flags().StringVarP(&imageAspectRatio, "ratio", "r", "1:1", "Aspect ratio (1:1, 4:3, 3:4, 16:9, 9:16)")
	imageCmd.Flags().StringVar(&imageResolution, "resolution", "1K", "Image resolution (512, 1K, 2K, 4K)")
	RootCmd.AddCommand(imageCmd)
}
