package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// gmailCmd represents the gmail command
var gmailCmd = &cobra.Command{
	Use:   "gmail",
	Short: "Manage Gmail (List, Send, Search, Threads, Labels, and more)",
	Long: `Gmail management tool supporting full email lifecycle operations.

Authentication:
  Gmail requires OAuth2 linking. If not linked, run:
  ./fastclaw gmail auth

Quick Examples:
  # List recent messages
  ./fastclaw gmail list

  # Search emails
  ./fastclaw gmail search "from:example.com is:unread"

  # Send an email
  ./fastclaw gmail send --to "user@example.com" --subject "Hello" --body "Content"

  # View a thread
  ./fastclaw gmail thread <thread_id>

  # Reply to a thread
  ./fastclaw gmail reply <thread_id> --body "My reply text"

  # Manage drafts
  ./fastclaw gmail drafts
  ./fastclaw gmail draft send <draft_id>

  # Label management
  ./fastclaw gmail labels
  ./fastclaw gmail label create "Work Projects"
  ./fastclaw gmail modify <message_id> --add-label "INBOX" --remove-label "UNREAD"
`,
}

// gmailListCmd lists recent messages from inbox
var gmailListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent messages from your inbox",
	Long: `Retrieves recent email messages from the authenticated user's inbox.

This command fetches messages with basic metadata (id, threadId, subject, sender, date).
Use "gmail search" for filtered queries or "gmail message" for full message details.

Examples:
  # Get last 10 messages
  ./fastclaw gmail list

  # Get last 20 messages
  ./fastclaw gmail list --max 20

Authentication:
  Requires Gmail OAuth2 link. If not linked, prompt user to authenticate at:
  https://fast-claw.xyz/api/skill/tool/auth-link?appName=gmail
`,
	Example: "  fastclaw gmail list\n  fastclaw gmail list --max 20",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching recent Gmail messages...")

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"actionName": "GMAIL_FETCH_EMAILS",
				"parameters": map[string]interface{}{
					"max_results": gmailMaxResults,
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

// gmailSearchCmd searches Gmail messages
var gmailSearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search Gmail messages using advanced query syntax",
	Long: `Search Gmail using the same powerful syntax as Gmail's search box.

QUERY SYNTAX:
  Operators:
    from:<address>    - Messages from specific sender
    to:<address>      - Messages to specific recipient
    subject:<text>    - Messages with text in subject
    label:<name>      - Messages with specific label (use label ID for custom labels)
    is:<state>        - State: unread, read, starred, important
    has:<property>    - Property: attachment, drive, photo
    after:<date>      - Messages after date (YYYY/MM/DD format, UTC)
    before:<date>     - Messages before date (YYYY/MM/DD format, UTC)
    in:<folder>       - Messages in specific folder
    category:<cat>    - Category: primary, social, promotions, updates, forums

  Logical Operators:
    AND (default)     - Both conditions must match
    OR                - Either condition matches
    NOT or -          - Exclude results

  Special:
    "exact phrase"    - Exact phrase match (use quotes)
    is:snoozed        - Snoozed messages
    is:muted          - Muted conversations

EXAMPLES:
  # Unread from specific sender
  ./fastclaw gmail search "from:boss@company.com is:unread"

  # Messages with attachments
  ./fastclaw gmail search "has:attachment subject:report"

  # Date range
  ./fastclaw gmail search "after:2024/01/01 before:2024/03/01"

  # Complex search
  ./fastclaw gmail search "from:newsletter@tech.com OR from:alerts@social.com -is:read label:CATEGORY_PROMOTIONS"

IMPORTANT:
  - Gmail search is case-insensitive
  - Use 'is:' for system states, 'label:' for custom/user-created labels
  - label:snoozed is WRONG - use is:snoozed instead
  - Date operators are UTC-based
`,
	Args: cobra.MinimumNArgs(1),
	Example: `  # Unread messages from sender
  fastclaw gmail search "from:important@sender.com is:unread"

  # Messages with attachments
  fastclaw gmail search "has:attachment subject:invoice"

  # Date range search
  fastclaw gmail search "after:2024/01/01 before:2024/12/31 subject:taxes"

  # Exclude read messages
  fastclaw gmail search "label:Work -is:read"`,
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		fmt.Printf("Searching Gmail: %s\n", query)

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"actionName": "GMAIL_FETCH_EMAILS",
				"parameters": map[string]interface{}{
					"query":       query,
					"max_results": gmailMaxResults,
				},
			}).
			Post("/api/skill/tool/execute")

		if err != nil || resp.IsError() {
			fmt.Printf("Error: Search failed: %v %s\n", err, resp.String())
			return
		}

		fmt.Println("-------------------------------------------")
		fmt.Println(resp.String())
		fmt.Println("-------------------------------------------")
	},
}

// gmailThreadCmd retrieves all messages in a thread
var gmailThreadCmd = &cobra.Command{
	Use:   "thread [thread_id]",
	Short: "Get all messages in a Gmail thread",
	Long: `Retrieves the complete message tree of a Gmail thread using its thread ID.

A thread ID is different from a message ID. Thread IDs are hexadecimal strings
(typically 15-16 characters like '19bf77729bcb3a44') and represent a conversation.

HOW TO GET A THREAD ID:
  1. From 'gmail list' or 'gmail search' - the response includes threadId field
  2. From 'gmail threads' command
  3. From other Gmail API responses

THREAD vs MESSAGE:
  - Thread ID (thread_id): Groups all messages in a conversation
  - Message ID (message_id): Individual email identifier

IMPORTANT:
  - Do NOT confuse thread_id with message_id
  - Do NOT use web UI legacy IDs (e.g., 'FMfcgzQfBZdVqKZcSVBhqwWLKWCtDdWQ')
  - Messages in a thread may have different labels (some archived, some not)
  - Message order in response is NOT guaranteed - sort by 'internalDate' to find oldest/newest
  - This command uses the thread_id, NOT message_id - passing a message_id will fail

EXAMPLES:
  # Get thread messages
  ./fastclaw gmail thread 19bf77729bcb3a44

  # With pagination (for long threads)
  ./fastclaw gmail thread 19bf77729bcb3a44 --page-token <nextPageToken>

AUTHENTICATION:
  Requires Gmail OAuth2 link. If not linked, user will be prompted to authenticate.

ERRORS:
  - "Invalid id value" - Usually means a message_id was used instead of thread_id
  - "thread not found" - Thread doesn't exist or is inaccessible
`,
	Args: cobra.ExactArgs(1),
	Example: `  # Get messages in a thread
  fastclaw gmail thread 19bf77729bcb3a44

  # Paginate through long threads
  fastclaw gmail thread 19bf77729bcb3a44 --page-token eyJwaWQiOjB9`,
	Run: func(cmd *cobra.Command, args []string) {
		threadID := args[0]
		pageToken, _ := cmd.Flags().GetString("page-token")

		fmt.Printf("Fetching thread: %s\n", threadID)

		params := map[string]interface{}{
			"thread_id": threadID,
		}
		if pageToken != "" {
			params["page_token"] = pageToken
		}

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"actionName": "GMAIL_FETCH_MESSAGE_BY_THREAD_ID",
				"parameters": params,
			}).
			Post("/api/skill/tool/execute")

		if err != nil || resp.IsError() {
			fmt.Printf("Error: Failed to fetch thread: %v %s\n", err, resp.String())
			return
		}

		fmt.Println("-------------------------------------------")
		fmt.Printf("Thread ID: %s\n", threadID)
		fmt.Println(resp.String())
		fmt.Println("-------------------------------------------")
	},
}

// gmailMessageCmd retrieves a specific email message
var gmailMessageCmd = &cobra.Command{
	Use:   "message [message_id]",
	Short: "Get a specific email message by its ID",
	Long: `Retrieves a single email message by its unique message ID.

MESSAGE ID vs THREAD ID:
  - message_id: 15-16 character hexadecimal string identifying ONE email
  - thread_id: The conversation ID grouping related emails

HOW TO GET MESSAGE ID:
  1. From 'gmail list' or 'gmail search' - response includes 'messageId' field
  2. From 'gmail threads' command
  3. From other Gmail API responses

IMPORTANT:
  - Do NOT confuse message_id with thread_id - they are different
  - Do NOT use email subjects, dates, or sender names as IDs
  - Do NOT use UUIDs (32-character strings) - those are NOT valid Gmail IDs
  - Format validation only checks structure - the ID must also exist and be accessible
  - Spam and trash messages are excluded unless explicitly requested

FORMAT OPTIONS:
  metadata (default) - Headers and metadata only, no body. Fast and sufficient for most uses.
  full               - Complete MIME structure with all headers and body. Heavy payload.
  minimal            - Lightest: ID, thread ID, labels only
  raw                - Raw RFC 2822 formatted message as base64url string

EXAMPLES:
  # Get message with metadata (default)
  ./fastclaw gmail message 18c5f5d1a2b3c4d

  # Get full message with body
  ./fastclaw gmail message 18c5f5d1a2b3c4d --format full

  # Get minimal info
  ./fastclaw gmail message 18c5f5d1a2b3c4d --format minimal

AUTHENTICATION:
  Requires Gmail OAuth2 link.
`,
	Args: cobra.ExactArgs(1),
	Example: `  # Get message with metadata
  fastclaw gmail message 18c5f5d1a2b3c4d

  # Get full message with body content
  fastclaw gmail message 18c5f5d1a2b3c4d --format full

  # Minimal retrieval (ID and labels only)
  fastclaw gmail message 18c5f5d1a2b3c4d --format minimal`,
	Run: func(cmd *cobra.Command, args []string) {
		messageID := args[0]
		format, _ := cmd.Flags().GetString("format")

		fmt.Printf("Fetching message: %s\n", messageID)

		params := map[string]interface{}{
			"message_id": messageID,
		}
		if format != "" {
			params["format"] = format
		}

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"actionName": "GMAIL_FETCH_MESSAGE_BY_MESSAGE_ID",
				"parameters": params,
			}).
			Post("/api/skill/tool/execute")

		if err != nil || resp.IsError() {
			fmt.Printf("Error: Failed to fetch message: %v %s\n", err, resp.String())
			return
		}

		fmt.Println("-------------------------------------------")
		fmt.Printf("Message ID: %s\n", messageID)
		fmt.Println(resp.String())
		fmt.Println("-------------------------------------------")
	},
}

// gmailReplyCmd replies to an existing thread
var gmailReplyCmd = &cobra.Command{
	Use:   "reply [thread_id]",
	Short: "Reply to an existing email thread",
	Long: `Sends a reply message within an existing Gmail thread.

This command replies to ALL participants in the thread and adds the reply as the
newest message in the conversation. The reply uses the original thread's subject.

THREAD vs NEW THREAD:
  - Using a valid thread_id: Reply appears in existing conversation
  - Providing a subject: Creates a NEW thread instead of replying

HOW TO GET THREAD ID:
  - From 'gmail list' or 'gmail search' - response includes 'threadId' field
  - From 'gmail threads' command
  - From message metadata

IMPORTANT:
  - Do NOT provide a custom subject - it will start a NEW conversation
  - Thread ID must be valid and accessible, not a message_id
  - At least one recipient (To), CC, or BCC must exist in the original message
  - Total message size including attachments must stay under 25MB
  - For large files, use Google Drive shareable links instead

MESSAGE BODY:
  - HTML supported: Use --html=true (default) for styled content
  - Plain text: Use --html=false
  - CC/BCC: Optional recipients to receive copies

EXAMPLES:
  # Reply with plain text
  ./fastclaw gmail reply 19bf77729bcb3a44 --body "Thank you for the update!"

  # Reply with HTML content
  ./fastclaw gmail reply 19bf77729bcb3a44 --body "<b>Thanks!</b><p>I will review and get back to you.</p>"

  # Reply with CC
  ./fastclaw gmail reply 19bf77729bcb3a44 --body "See you at 3pm" --cc colleague@company.com

  # Reply to own message in thread
  ./fastclaw gmail reply 19bf77729bcb3a44 --body "Confirmed, I'm attending." --to me@gmail.com

AUTHENTICATION:
  Requires Gmail OAuth2 link.
`,
	Args: cobra.ExactArgs(1),
	Example: `  # Plain text reply
  fastclaw gmail reply 19bf77729bcb3a44 --body "Got it, thanks!"

  # HTML formatted reply
  fastclaw gmail reply 19bf77729bcb3a44 --body "<p>I'll review and respond shortly.</p>" --html

  # Reply with CC
  fastclaw gmail reply 19bf77729bcb3a44 --body "Team meeting confirmed" --cc manager@company.com

  # Reply with BCC
  fastclaw gmail reply 19bf77729bcb3a44 --body "Sending to archive" --bcc archive@company.com`,
	Run: func(cmd *cobra.Command, args []string) {
		threadID := args[0]

		if replyBody == "" {
			fmt.Println("Error: --body is required for replying")
			return
		}

		fmt.Printf("Replying to thread: %s\n", threadID)

		params := map[string]interface{}{
			"thread_id":    threadID,
			"message_body": replyBody,
			"is_html":      replyHtml,
		}

		if len(replyCc) > 0 {
			params["cc"] = replyCc
		}
		if len(replyBcc) > 0 {
			params["bcc"] = replyBcc
		}
		if replyTo != "" {
			params["recipient_email"] = replyTo
		}

		client := GetClient()
		resp, err := client.R().
			SetBody(map[string]interface{}{
				"actionName": "GMAIL_REPLY_TO_THREAD",
				"parameters": params,
			}).
			Post("/api/skill/tool/execute")

		if err != nil || resp.IsError() {
			fmt.Printf("Error: Failed to send reply: %v %s\n", err, resp.String())
			return
		}

		fmt.Println("Reply sent successfully!")
	},
}

var (
	sendTo           string
	sendSubject      string
	sendBody         string
	sendHtml         bool
	gmailMaxResults  int
	replyBody        string
	replyHtml        bool
	replyCc          []string
	replyBcc         []string
	replyTo          string
)

var gmailSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a new email message",
	Long: `Sends a new email to one or more recipients. This creates a NEW thread,
not a reply to an existing conversation. For replying, use 'gmail reply'.

RECIPIENTS:
  - --to (required): Primary recipient email address
  - --cc: Carbon Copy recipients
  - --bcc: Blind Carbon Copy recipients

BODY CONTENT:
  - HTML mode (default): Body is treated as HTML - use for styled content, links, tables
  - Plain text: Use --html=false for plain text emails

SUBJECT:
  - Required for sending (except when replying to a thread)
  - A subject with a thread_id will create NEW thread, not reply

ATTACHMENTS:
  - Not supported via CLI flags directly
  - For attachments, first upload to storage, then use the s3key reference

SIZE LIMITS:
  - Total message size must be under ~25MB after base64 encoding
  - For larger files, use Google Drive shareable links

EXAMPLES:
  # Send HTML email (default)
  ./fastclaw gmail send --to "user@example.com" --subject "Meeting Request" --body "<h1>Hi</h1><p>Let's meet tomorrow.</p>"

  # Send plain text email
  ./fastclaw gmail send --to "user@example.com" --subject "Quick Note" --body "Just a plain text message" --html=false

  # Send with CC
  ./fastclaw gmail send --to "main@recipient.com" --cc "copy@recipient.com" --subject "Report" --body "<p>See attached report.</p>"

AUTHENTICATION:
  Requires Gmail OAuth2 link.
`,
	Example: `  # Send HTML email
  fastclaw gmail send --to "user@example.com" --subject "Hello" --body "<h1>Hello</h1><p>Welcome!</p>"

  # Send plain text
  fastclaw gmail send --to "user@example.com" --subject "Quick Note" --body "Plain text message" --html=false

  # Send with CC and BCC
  fastclaw gmail send --to "main@example.com" --cc "copy@example.com" --bcc "hidden@example.com" --subject "Info" --body "Details here"`,
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

		fmt.Println("Email sent successfully!")
	},
}

func init() {
	gmailCmd.PersistentFlags().IntVar(&gmailMaxResults, "max", 10, "Maximum number of messages to retrieve")

	gmailSendCmd.Flags().StringVar(&sendTo, "to", "", "Recipient email address")
	gmailSendCmd.Flags().StringVar(&sendSubject, "subject", "", "Email subject")
	gmailSendCmd.Flags().StringVar(&sendBody, "body", "", "Email body content (supports HTML)")
	gmailSendCmd.Flags().BoolVar(&sendHtml, "html", true, "Send email as HTML (default: true)")

	gmailReplyCmd.Flags().StringVar(&replyBody, "body", "", "Reply message body content")
	gmailReplyCmd.Flags().BoolVar(&replyHtml, "html", false, "Send reply as HTML (default: false)")
	gmailReplyCmd.Flags().StringArrayVar(&replyCc, "cc", []string{}, "CC recipient email addresses")
	gmailReplyCmd.Flags().StringArrayVar(&replyBcc, "bcc", []string{}, "BCC recipient email addresses")
	gmailReplyCmd.Flags().StringVar(&replyTo, "to", "", "Reply to specific recipient (optional)")

	gmailThreadCmd.Flags().String("page-token", "", "Page token for paginating through long threads")
	gmailMessageCmd.Flags().String("format", "metadata", "Message format: metadata (default), full, minimal, raw")

	gmailCmd.AddCommand(gmailListCmd)
	gmailCmd.AddCommand(gmailSendCmd)
	gmailCmd.AddCommand(gmailSearchCmd)
	gmailCmd.AddCommand(gmailThreadCmd)
	gmailCmd.AddCommand(gmailMessageCmd)
	gmailCmd.AddCommand(gmailReplyCmd)
	RootCmd.AddCommand(gmailCmd)
}
