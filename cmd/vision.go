package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var visionPrompt string

// visionCmd represents the vision command
var visionCmd = &cobra.Command{
	Use:   "vision [image_url_or_path]",
	Short: "Analyze image content using NVIDIA Vision models",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageInput := args[0]
		imageUrl := imageInput

		// 입력값이 로컬 파일인 경우 먼저 업로드
		if _, err := os.Stat(imageInput); err == nil {
			fmt.Printf("Local file detected. Uploading %s first...\n", imageInput)
			
			client := GetClient()
			resp, err := client.R().
				SetFile("file", imageInput).
				SetResult(&UploadResponse{}).
				Post("/api/skill/storage/r2/upload")

			if err != nil || resp.IsError() {
				fmt.Printf("Error: Failed to upload file for vision analysis: %v\n", err)
				return
			}
			
			uploadResult := resp.Result().(*UploadResponse)
			imageUrl = uploadResult.PublicUrl
			fmt.Printf("File uploaded successfully. Public URL: %s\n", imageUrl)
		} else if !strings.HasPrefix(imageInput, "http") {
			fmt.Printf("Error: Invalid image input '%s'. Must be a URL or a valid local file path.\n", imageInput)
			return
		}

		fmt.Println("Analyzing image...")

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"imageUrl": imageUrl,
				"prompt":   visionPrompt,
			}).
			SetResult(&VisionResponse{}).
			Post("/api/skill/vision/analyze")

		if err != nil {
			fmt.Printf("Error: Vision API call failed: %v\n", err)
			return
		}

		if resp.IsError() {
			fmt.Printf("Error: Vision API responded with status %d: %s\n", resp.StatusCode(), resp.String())
			return
		}

		result := resp.Result().(*VisionResponse)
		fmt.Println("-------------------------------------------")
		fmt.Println("Analysis Result:")
		fmt.Println(result.Result)
		fmt.Printf("\nCost: %f Credits\n", result.Cost)
		fmt.Println("-------------------------------------------")
	},
}

type VisionResponse struct {
	Result string  `json:"result"`
	Usage  interface{} `json:"usage"`
	Cost   float64 `json:"cost"`
}

func init() {
	visionCmd.Flags().StringVarP(&visionPrompt, "prompt", "p", "Describe this image in detail.", "Custom analysis prompt")
	RootCmd.AddCommand(visionCmd)
}
