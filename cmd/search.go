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
	Short: "Perform a Google search via FastClaw (Integrated Search)",
	Long: `FastClaw Search skill provides access to Google Search with localized options.
AI AGENTS: Always consider 'gl' (country) and 'hl' (language) flags for better local results.
- For Korean results: Use '--gl kr --hl ko'
- For US results: Use '--gl us --hl en' (default)
- Increase '--num' (or '-n') up to 10 for more comprehensive analysis.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch("/api/skill/google_search/search", "organic", args)
	},
}

// Sub-search commands
var newsSearchCmd = &cobra.Command{
	Use:   "news [query]",
	Short: "Search Google News with timestamps",
	Long: `Retrieve the latest news articles for a query. 
Returns: Title, Link, Source, and Publication Date.
AI AGENTS: Use this for monitoring trending topics or recent events.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch("/api/skill/google_search/news", "news", args)
	},
}

var imageSearchCmd = &cobra.Command{
	Use:   "images [query]",
	Short: "Search Google Images for URLs",
	Long: `Find image URLs and their context pages. 
Returns: Title, ImageUrl, and the original link.
AI AGENTS: Use this when you need visual references or assets.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runSearch("/api/skill/google_search/images", "images", args)
	},
}

var mapSearchCmd = &cobra.Command{
	Use:   "maps [query]",
	Short: "Search Google Maps (Places, Ratings, Contact)",
	Long: `Retrieve detailed information about locations and businesses. 
Returns: Business Name, Address, Rating, Review Count, Phone, and Website.
AI AGENTS: Use this for lead generation, business research, or location-based services.`,
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
	searchCmd.PersistentFlags().StringVar(&searchGl, "gl", "us", "Country code (e.g., 'kr', 'us', 'jp')")
	searchCmd.PersistentFlags().StringVar(&searchHl, "hl", "en", "Language code (e.g., 'ko', 'en', 'ja')")
	searchCmd.PersistentFlags().IntVarP(&searchNum, "num", "n", 5, "Number of results to return (max: 10)")

	searchCmd.AddCommand(newsSearchCmd)
	searchCmd.AddCommand(imageSearchCmd)
	searchCmd.AddCommand(mapSearchCmd)
	
	RootCmd.AddCommand(searchCmd)
}
