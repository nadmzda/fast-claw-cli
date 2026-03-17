package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// googleCmd represents the google command
var googleCmd = &cobra.Command{
	Use:   "google",
	Short: "Manage Google Workspace services (Calendar, Drive, Sheets, Tasks)",
}

// --- Calendar ---
var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Manage Google Calendar",
}

var calendarAddCmd = &cobra.Command{
	Use:   "add [event_text]",
	Short: "Quickly add an event using natural language",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		ExecuteToolAction("GOOGLECALENDAR_QUICK_ADD", map[string]interface{}{
			"text": text,
		}, "Calendar event added!")
	},
}

// --- Drive ---
var driveCmd = &cobra.Command{
	Use:   "drive",
	Short: "Manage Google Drive",
}

var driveSearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for files in Google Drive",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		ExecuteToolAction("GOOGLEDRIVE_FIND_FILE", map[string]interface{}{
			"q": fmt.Sprintf("name contains '%s'", query),
		}, "Search results:")
	},
}

// --- Sheets ---
var sheetsCmd = &cobra.Command{
	Use:   "sheets",
	Short: "Manage Google Sheets",
}

var sheetsReadCmd = &cobra.Command{
	Use:   "read [spreadsheet_id] [range]",
	Short: "Read values from a spreadsheet",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteToolAction("GOOGLESHEETS_VALUES_GET", map[string]interface{}{
			"spreadsheet_id": args[0],
			"range":          args[1],
		}, "Sheet Data:")
	},
}

// --- Tasks ---
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Manage Google Tasks",
}

var tasksAddCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task to default list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		ExecuteToolAction("GOOGLETASKS_INSERT_TASK", map[string]interface{}{
			"tasklist_id": "@default",
			"title":       title,
		}, "Task created!")
	},
}

func init() {
	calendarCmd.AddCommand(calendarAddCmd)
	driveCmd.AddCommand(driveSearchCmd)
	sheetsCmd.AddCommand(sheetsReadCmd)
	tasksCmd.AddCommand(tasksAddCmd)

	googleCmd.AddCommand(calendarCmd)
	googleCmd.AddCommand(driveCmd)
	googleCmd.AddCommand(sheetsCmd)
	googleCmd.AddCommand(tasksCmd)
	
	RootCmd.AddCommand(googleCmd)
}
