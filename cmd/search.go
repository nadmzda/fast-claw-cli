package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	searchGl  string
	searchHl  string
	searchNum int
)

// searchCmd represents the base search command
var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Perform a Google search via FastClaw",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch("/api/skill/google_search/search", "organic", args)
	},
}

// Sub-search commands
var newsSearchCmd = &cobra.Command{
	Use:   "news [query]",
	Short: "Search Google News",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch("/api/skill/google_search/news", "news", args)
	},
}

var imageSearchCmd = &cobra.Command{
	Use:   "images [query]",
	Short: "Search Google Images",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch("/api/skill/google_search/images", "images", args)
	},
}

var mapSearchCmd = &cobra.Command{
	Use:   "maps [query]",
	Short: "Search Google Maps",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch("/api/skill/google_search/maps", "places", args)
	},
}

func runSearch(endpoint string, resultKey string, args []string) {
	query := strings.Join(args, " ")
	fmt.Printf("Searching for: %s via %s ...\n", query, endpoint)

	client := GetClient()
	resp, err := client.R().
		SetBody(map[string]interface{}{
			"q":   query,
			"gl":  searchGl,
			"hl":  searchHl,
			"num": searchNum,
		}).
		Post(endpoint)

	if err != nil || resp.IsError() {
		fmt.Printf("Error: Search failed: %v %s\n", err, resp.String())
		return
	}

	// 응답 결과 파싱 (동적 필드 처리)
	var body map[string]interface{}
	err = json.Unmarshal(resp.Body(), &body)
	if err != nil {
		fmt.Printf("Error: Failed to parse response: %v\n", err)
		return
	}

	results, ok := body[resultKey].([]interface{})
	if !ok || len(results) == 0 {
		fmt.Printf("No results found in '%s' field.\n", resultKey)
		return
	}

	fmt.Println("-------------------------------------------")
	for i, res := range results {
		item := res.(map[string]interface{})
		fmt.Printf("[%d] %v\n", i+1, item["title"])

		switch resultKey {
		case "organic":
			fmt.Printf("    Link: %v\n", item["link"])
		case "news":
			fmt.Printf("    Source: %v | Date: %v\n", item["source"], item["date"])
			fmt.Printf("    Link: %v\n", item["link"])
		case "images":
			fmt.Printf("    Image: %v\n", item["imageUrl"])
			fmt.Printf("    Context: %v\n", item["link"])
		case "places":
			fmt.Printf("    Address: %v\n", item["address"])
			if r, exists := item["rating"]; exists {
				fmt.Printf("    Rating: %v (%v reviews)\n", r, item["reviews"])
			}
			if p, exists := item["phoneNumber"]; exists {
				fmt.Printf("    Phone: %v\n", p)
			}
			if w, exists := item["website"]; exists {
				fmt.Printf("    Web: %v\n", w)
			}
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------------------")
}

func init() {
	searchCmd.PersistentFlags().StringVar(&searchGl, "gl", "us", "Country code")
	searchCmd.PersistentFlags().StringVar(&searchHl, "hl", "en", "Language code")
	searchCmd.PersistentFlags().IntVar(&searchNum, "num", 5, "Number of results")

	searchCmd.AddCommand(newsSearchCmd)
	searchCmd.AddCommand(imageSearchCmd)
	searchCmd.AddCommand(mapSearchCmd)
	
	RootCmd.AddCommand(searchCmd)
}
