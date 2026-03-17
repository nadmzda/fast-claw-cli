package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gmailCmd represents the gmail command
var gmailCmd = &cobra.Command{
	Use:   "gmail",
	Short: "Manage Gmail (List, Send, Search)",
}

var gmailListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent messages from your inbox",
	Example: "  fastclaw gmail list -k YOUR_API_KEY",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching recent Gmail messages...")
		
		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"actionName": "GMAIL_LIST_MESSAGES",
				"parameters": map[string]interface{}{
					"max_results": 10,
				},
			}).
			Post("/api/skill/tool/execute")

		if err != nil || resp.IsError() {
			fmt.Printf("Error: Failed to fetch messages: %v %s\n", err, resp.String())
			fmt.Println("Tip: Check if your Gmail is linked at https://fast-claw.xyz/api/skill/tool/auth-link?appName=gmail")
			return
		}

		fmt.Println("-------------------------------------------")
		fmt.Println(resp.String())
		fmt.Println("-------------------------------------------")
	},
}

var (
	sendTo      string
	sendSubject string
	sendBody    string
	sendHtml    bool
)

var gmailSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send an email (Supports HTML by default)",
	Long: `Send an email to a specified recipient. 
By default, the body is treated as HTML, allowing you to send styled content, 
links, and tables. Use --html=false to send as plain text.`,
	Example: `  # Send a styled HTML email (Default)
  fastclaw gmail send --to "user@example.com" --subject "Hello" --body "<h1>Hello</h1><p>Welcome!</p>"

  # Send a plain text email
  fastclaw gmail send --to "user@example.com" --subject "Quick Note" --body "Just a plain text" --html=false`,
	Run: func(cmd *cobra.Command, args []string) {
		if sendTo == "" || sendSubject == "" || sendBody == "" {
			fmt.Println("Error: --to, --subject, and --body are required.")
			return
		}

		fmt.Printf("Sending email to %s (HTML mode: %v)...\n", sendTo, sendHtml)
		
		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"actionName": "GMAIL_SEND_EMAIL",
				"parameters": map[string]interface{}{
					"to":      sendTo,
					"subject": sendSubject,
					"body":    sendBody,
					"is_html": sendHtml,
				},
			}).
			Post("/api/skill/tool/execute")

		if err != nil || resp.IsError() {
			fmt.Printf("Error: Failed to send email: %v %s\n", err, resp.String())
			return
		}

		fmt.Println("✨ Email sent successfully!")
	},
}

func init() {
	gmailSendCmd.Flags().StringVar(&sendTo, "to", "", "Recipient email address")
	gmailSendCmd.Flags().StringVar(&sendSubject, "subject", "", "Email subject")
	gmailSendCmd.Flags().StringVar(&sendBody, "body", "", "Email body content (supports HTML)")
	gmailSendCmd.Flags().BoolVar(&sendHtml, "html", true, "Send email as HTML (default: true)")
	
	gmailCmd.AddCommand(gmailListCmd)
	gmailCmd.AddCommand(gmailSendCmd)
	RootCmd.AddCommand(gmailCmd)
}
