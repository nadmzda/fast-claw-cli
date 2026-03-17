package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// socialCmd represents the social command
var socialCmd = &cobra.Command{
	Use:   "social",
	Short: "Manage Social Media accounts (Instagram, Reddit)",
}

// --- Instagram ---
var instagramCmd = &cobra.Command{
	Use:   "instagram",
	Short: "Manage Instagram (Business)",
}

var instagramPostCmd = &cobra.Command{
	Use:   "post [image_url] [caption]",
	Short: "Create a new image post",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("INSTAGRAM_CREATE_POST", map[string]interface{}{
			"image_url": args[0],
			"caption":   args[1],
		}, "Instagram post created!")
	},
}

// --- Reddit ---
var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "Manage Reddit",
}

var redditPostCmd = &cobra.Command{
	Use:   "post [subreddit] [title] [text]",
	Short: "Create a new text post on Reddit",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("REDDIT_CREATE_REDDIT_POST", map[string]interface{}{
			"subreddit": args[0],
			"title":     args[1],
			"text":      args[2],
			"kind":      "self",
		}, "Reddit post created!")
	},
}

var redditSearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search across Reddit",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		ExecuteToolAction("REDDIT_SEARCH_ACROSS_SUBREDDITS", map[string]interface{}{
			"search_query": query,
		}, "Search results:")
	},
}

func init() {
	instagramCmd.AddCommand(instagramPostCmd)
	redditCmd.AddCommand(redditPostCmd)
	redditCmd.AddCommand(redditSearchCmd)

	socialCmd.AddCommand(instagramCmd)
	socialCmd.AddCommand(redditCmd)
	
	RootCmd.AddCommand(socialCmd)
}
