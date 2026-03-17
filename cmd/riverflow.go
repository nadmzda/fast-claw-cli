package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// riverflowCmd represents the riverflow command
var riverflowCmd = &cobra.Command{
	Use:   "riverflow [prompt]",
	Short: "Quickly generate images using Riverflow model",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := strings.Join(args, " ")
		fmt.Printf("Generating fast image for: %s ...\n", prompt)

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"prompt": prompt,
			}).
			SetResult(&RiverflowResponse{}).
			Post("/api/skill/riverflow/generate")

		if err != nil || resp.IsError() {
			fmt.Printf("Error: Riverflow generation failed: %v %s\n", err, resp.String())
			return
		}

		result := resp.Result().(*RiverflowResponse)
		fmt.Println("-------------------------------------------")
		fmt.Printf("Success! Image URL: %s\n", result.PublicUrl)
		fmt.Printf("Cost: %f Credits\n", result.Cost)
		fmt.Println("-------------------------------------------")
	},
}

type RiverflowResponse struct {
	PublicUrl string  `json:"publicUrl"`
	Prompt    string  `json:"prompt"`
	Cost      float64 `json:"cost"`
}

func init() {
	RootCmd.AddCommand(riverflowCmd)
}
