package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape [url]",
	Short: "Extract high-quality text content and metadata from any public URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		targetUrl := args[0]
		fmt.Printf("Scraping URL: %s ...\n", targetUrl)

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"url": targetUrl,
			}).
			SetResult(&ScrapeResponse{}).
			Post("/api/skill/web_scrape/scrape")

		if err != nil {
			fmt.Printf("Error: Scraping failed: %v\n", err)
			return
		}

		if resp.IsError() {
			fmt.Printf("Error: Scrape API responded with status %d: %s\n", resp.StatusCode(), resp.String())
			return
		}

		result := resp.Result().(*ScrapeResponse)
		fmt.Println("-------------------------------------------")
		fmt.Printf("Title: %s\n", result.Metadata.Title)
		if result.Metadata.Description != "" {
			fmt.Printf("Description: %s\n", result.Metadata.Description)
		}
		fmt.Println("-------------------------------------------")
		fmt.Println("Content:")
		
		// 텍스트가 너무 길면 CLI 출력을 위해 적절히 처리 (예: 1000자까지만 표시)
		content := result.Text
		if len(content) > 1000 {
			content = content[:1000] + "... (truncated)"
		}
		fmt.Println(content)
		fmt.Println("-------------------------------------------")
	},
}

type ScrapeResponse struct {
	Text     string `json:"text"`
	Metadata struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"metadata"`
	Credits float64 `json:"credits"`
}

func init() {
	RootCmd.AddCommand(scrapeCmd)
}
