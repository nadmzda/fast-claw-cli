package cmd

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var (
	apiKey  string
	verbose bool
	version = "dev" // LDFlags를 통해 빌드 시 주입됨
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "fastclaw",
	Version: version,
	Short:   "FastClaw CLI tool to access AI skills and storage",
	Long: `FastClaw CLI is a powerful tool to interact with FastClaw API skills 
including file storage, vision analysis, and web search.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Persistent flags (available to all subcommands)
	RootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", os.Getenv("FASTCLAW_API_KEY"), "FastClaw API Key (env: FASTCLAW_API_KEY)")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Verbose output")
}

// GetClient returns a resty client configured with the API Key
func GetClient() *resty.Client {
	client := resty.New()
	client.SetBaseURL("https://fast-claw.xyz")
	if apiKey != "" {
		client.SetHeader("X-API-KEY", apiKey)
	}
	return client
}

// ExecuteToolAction runs a generic tool action via /api/skill/tool/execute
func ExecuteToolAction(actionName string, params map[string]interface{}, successMsg string) {
	client := GetClient()
	resp, err := client.R().
		SetBody(map[string]interface{}{
			"actionName": actionName,
			"parameters": params,
		}).
		Post("/api/skill/tool/execute")

	if err != nil || resp.IsError() {
		fmt.Printf("Error: Action '%s' failed: %v %s\n", actionName, err, resp.String())
		return
	}

	fmt.Println("-------------------------------------------")
	fmt.Println(successMsg)
	fmt.Println(resp.String())
	fmt.Println("-------------------------------------------")
}

// SearchResponse represents common Google search response
type SearchResponse struct {
	Organic []struct {
		Title   string `json:"title"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"organic"`
}
