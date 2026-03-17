package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// notionCmd represents the notion command
var notionCmd = &cobra.Command{
	Use:   "notion",
	Short: "Manage Notion workspace",
}

var notionCreateCmd = &cobra.Command{
	Use:   "create [parent_id] [title] [markdown]",
	Short: "Create a new Notion page with Markdown",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("NOTION_CREATE_NOTION_PAGE", map[string]interface{}{
			"parent_id": args[0],
			"title":     args[1],
			"markdown":  args[2],
		}, "Notion page created!")
	},
}

var notionSearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for pages or databases",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		ExecuteToolAction("NOTION_SEARCH_NOTION_PAGE", map[string]interface{}{
			"query": query,
		}, "Search results:")
	},
}

func init() {
	notionCmd.AddCommand(notionCreateCmd)
	notionCmd.AddCommand(notionSearchCmd)
	
	RootCmd.AddCommand(notionCmd)
}
