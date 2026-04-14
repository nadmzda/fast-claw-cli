# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands

```powershell
go build -o fastclaw.exe main.go    # Local build
goreleaser release --clean          # Cross-platform release (requires GITHUB_TOKEN)
```

## Architecture

FastClaw CLI is a Go-based CLI tool that wraps a backend API (`https://fast-claw.xyz`) using Cobra for command structure.

### Key Patterns

- **Commands**: Each file in `cmd/` is a Cobra command registered via `init()`. Nested subcommands use `ParentCmd.AddCommand(childCmd)`.
- **API Communication**: `GetClient()` in `cmd/root.go` creates a resty client with the API key. `ExecuteToolAction()` is the common pattern for calling `/api/skill/tool/execute` with action names and parameters.
- **Authentication Priority**: Flag (`-k`) > Env (`FASTCLAW_API_KEY`) > Config File (`~/.fastclaw_config`)
- **Version**: Injected via ldflags at build time (`version` variable in `cmd/root.go`)

### Command Structure

| Command | Description |
|---|---|
| `upload` | Single (<20MB) or multipart upload to R2 storage |
| `vision` | NVIDIA Vision analysis and OCR |
| `image` | Gemini image generation |
| `riverflow` | Fast Riverflow image generation |
| `music` | Suno AI music generation |
| `search` | Google search (news, maps, standard) |
| `scrape` | Web page text extraction |
| `google` | Calendar, Drive, Sheets, Tasks |
| `gmail` | Email management |
| `notion` | Notion page creation/search |
| `social` | Instagram posts, Reddit posts/comments |
| `config` | API key management |
| `update` | Self-update via go-selfupdate |

### File Naming Conventions

Cobra commands follow the pattern: command name is the filename, `init()` registers it, and response types are defined at the bottom of each file (e.g., `ImageResponse` in `cmd/image.go`).