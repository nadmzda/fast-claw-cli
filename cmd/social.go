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
	Short: "Manage Reddit (Posts, Comments, Communities, and Users)",
	Long: `FastClaw Reddit skill provides full access to the Reddit platform. 
It supports creating content, interacting with communities, and retrieving data for AI analysis.
IMPORTANT: Reddit uses 'fullnames' for identifiers:
- t1_[ID]: Comment (e.g., t1_k12345)
- t3_[ID]: Post/Submission (e.g., t3_s12345)
- t5_[ID]: Subreddit`,
}

var redditPostCmd = &cobra.Command{
	Use:   "post [subreddit] [title] [text]",
	Short: "Create a new text post on Reddit",
	Long:  "Submit a new text-based post (self-post) to a specific subreddit. Do not include 'r/' in the subreddit name.",
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
	Short: "Search across all of Reddit",
	Long:  "Perform a global search for posts matching the query. Returns titles, links, and snippets.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		ExecuteToolAction("REDDIT_SEARCH_ACROSS_SUBREDDITS", map[string]interface{}{
			"search_query": query,
		}, "Search results:")
	},
}

var redditListNewCmd = &cobra.Command{
	Use:   "list-new [subreddit]",
	Short: "Get newest posts from a subreddit",
	Long:  "Retrieve the most recent submissions from a specific subreddit for monitoring or analysis.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		ExecuteToolAction("REDDIT_GET_NEW", map[string]interface{}{
			"subreddit": args[0],
			"limit":     limit,
		}, "Newest posts:")
	},
}

var redditListTopCmd = &cobra.Command{
	Use:   "list-top [subreddit]",
	Short: "Get top posts from a subreddit",
	Long:  "Retrieve the top-rated submissions from a subreddit. You can filter by time (hour, day, week, month, year, all).",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		timeFilter, _ := cmd.Flags().GetString("time")
		ExecuteToolAction("REDDIT_GET_R_TOP", map[string]interface{}{
			"subreddit": args[0],
			"limit":     limit,
			"t":         timeFilter,
		}, "Top posts:")
	},
}

var redditCommentsCmd = &cobra.Command{
	Use:   "comments [post_id]",
	Short: "Get all comments for a specific post",
	Long:  "Retrieve the full comment tree for a given post ID (e.g., '10omtdx' from 't3_10omtdx').",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("REDDIT_RETRIEVE_POST_COMMENTS", map[string]interface{}{
			"article": args[0],
		}, "Post comments:")
	},
}

var redditReplyCmd = &cobra.Command{
	Use:   "reply [thing_id] [text]",
	Short: "Reply to a post or comment",
	Long:  "Add a new comment as a reply to a parent item. thing_id must be a fullname (t1_... for comments, t3_... for posts).",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("REDDIT_POST_REDDIT_COMMENT", map[string]interface{}{
			"thing_id": args[0],
			"text":     args[1],
		}, "Reply sent!")
	},
}

var redditDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete your post or comment",
	Long:  "Permanently remove your own content from Reddit. id must be a fullname starting with t1_ or t3_.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "REDDIT_DELETE_REDDIT_POST"
		if strings.HasPrefix(args[0], "t1_") {
			action = "REDDIT_DELETE_REDDIT_COMMENT"
		}
		ExecuteToolAction(action, map[string]interface{}{
			"id": args[0],
		}, "Content deleted!")
	},
}

var redditUserCmd = &cobra.Command{
	Use:   "user [username]",
	Short: "Get Reddit user profile information",
	Long:  "Retrieve karma, cake day, and profile details for a specific user. Use 'me' to see your own info.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("REDDIT_GET_REDDIT_USER_ABOUT", map[string]interface{}{
			"username": args[0],
		}, "User profile info:")
	},
}

var redditRulesCmd = &cobra.Command{
	Use:   "rules [subreddit]",
	Short: "Fetch rules for a specific subreddit",
	Long:  "Get the official rules and guidelines of a community to ensure content compliance.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("REDDIT_GET_SUBREDDIT_RULES", map[string]interface{}{
			"subreddit": args[0],
		}, "Subreddit rules:")
	},
}

func init() {
	instagramCmd.AddCommand(instagramPostCmd)

	// Reddit Subcommands
	redditListNewCmd.Flags().IntP("limit", "l", 10, "Maximum number of items to retrieve")
	redditListTopCmd.Flags().IntP("limit", "l", 10, "Maximum number of items to retrieve")
	redditListTopCmd.Flags().StringP("time", "t", "all", "Time filter (hour, day, week, month, year, all)")

	redditCmd.AddCommand(redditPostCmd)
	redditCmd.AddCommand(redditSearchCmd)
	redditCmd.AddCommand(redditListNewCmd)
	redditCmd.AddCommand(redditListTopCmd)
	redditCmd.AddCommand(redditCommentsCmd)
	redditCmd.AddCommand(redditReplyCmd)
	redditCmd.AddCommand(redditDeleteCmd)
	redditCmd.AddCommand(redditUserCmd)
	redditCmd.AddCommand(redditRulesCmd)

	socialCmd.AddCommand(instagramCmd)
	socialCmd.AddCommand(redditCmd)
	
	RootCmd.AddCommand(socialCmd)
}
