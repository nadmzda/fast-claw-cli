# FastClaw CLI 🚀
[한국어 문서 (README_KR.md)](./README_KR.md)

FastClaw CLI is an **AI-powered multi-purpose tool** designed to execute AI skills, manage cloud storage, and automate workspaces (Google, Notion, etc.) directly from your terminal.

Built with Go, it is fast, lightweight, and distributed as a single binary for ease of use.

## ✨ Key Features

- 📁 **Smart File Upload**: Automatic detection of file size. Single upload for <20MB, and Multipart upload for larger files.
- 👁️ **NVIDIA Vision Analysis**: Analyze image URLs or upload local images automatically for OCR and detailed descriptions.
- 🎨 **AI Image Generation**: Generate high-quality images using Gemini 3.1 and ultra-fast Riverflow models.
- 🎵 **AI Music Generation**: Create original songs with custom style, lyrics, and vocals using Suno AI.
- 🔍 **Specialized Search**: Tailored outputs for Google News, Images, and Maps (Places, ratings, etc.).
- 🏢 **Workspace Integration**: Full support for Gmail, Calendar, Drive, Sheets, Tasks, and Notion.
- 📱 **Social Automation**: Post to Instagram Business and manage Reddit submissions/comments.

## 🛠 Installation

### 1. Build from Source (Go required)
```powershell
git clone https://github.com/your-username/fast-claw-cli.git
cd fast-claw-cli
go build -o fastclaw.exe main.go
```

### 2. Set Environment Variables (Optional)
To avoid entering your API key every time, set it as an environment variable.
```powershell
$env:FASTCLAW_API_KEY="your_api_key_here"
```

### 3. Self-Update
To update FastClaw CLI to the latest version:
```powershell
./fastclaw update
```

## 🔑 Authentication & Setup

- **API Key**: You need a key from the [FastClaw Dashboard](https://fast-claw.xyz).
- **Service Linking**: Workspace and Social skills require a one-time setup via Magic Links:
  - Gmail: [Authenticate](https://fast-claw.xyz/api/skill/tool/auth-link?appName=gmail)
  - Google Drive: [Authenticate](https://fast-claw.xyz/api/skill/tool/auth-link?appName=drive)
  - Notion: [Authenticate](https://fast-claw.xyz/api/skill/tool/auth-link?appName=notion)

## 📖 Usage Guide

### 1. Upload & Vision Analysis
```powershell
# Upload large files (Auto-multipart)
./fastclaw upload ./my_large_video.mp4

# Local image analysis (Auto-upload then analyze)
./fastclaw vision ./receipt.jpg --prompt "What is the total amount on this receipt?"
```

### 2. AI Image Generation
```powershell
# High-quality image (Gemini)
./fastclaw image "A futuristic Seoul city with flying cars" --ratio 16:9

# With custom resolution (512, 1K, 2K, 4K)
./fastclaw image "A futuristic Seoul city" --ratio 16:9 --resolution 2K

# Fast preview image (Riverflow)
./fastclaw riverflow "Cute robot cat"
```

### 3. AI Music Generation (Suno)
```powershell
# Generate a song (auto lyrics)
./fastclaw music "여름 바다로 떠나는 신나는 여행" --style "K-Pop, Dance, Upbeat" --title "Summer Wave" --email user@example.com --vocal Female

# Check Suno API health
./fastclaw music health
```

### 4. Google Search (Specialized)
```powershell
# Standard search
./fastclaw search "Fast-Claw API"

# Maps search (Outputs address, rating, etc.)
./fastclaw search maps "Best coffee shops in Gangnam" --num 3

# News search
./fastclaw search news "latest AI trends 2026"
```

### 5. Workspace & Social
```powershell
# Add Calendar event (NLP support)
./fastclaw google calendar add "Tomorrow at 2pm meeting with team"

# Send Gmail
./fastclaw gmail send --to "example@gmail.com" --subject "Hello" --body "Sent from CLI"

# Post to Reddit
./fastclaw social reddit post "test" "Hello World" "This is a post from FastClaw CLI"
```

## 📜 Command Overview

| Command | Description |
| :--- | :--- |
| `upload` | Upload files (Single or Multipart) |
| `vision` | Image analysis & OCR |
| `search` | Specialized Google search |
| `gmail` | List and send emails |
| `google` | Manage Calendar, Drive, Sheets, Tasks |
| `image` | High-quality image generation (Gemini) |
| `riverflow` | Fast image generation (Riverflow) |
| `music` | AI music generation (Suno) |
| `scrape` | Extract text from any web page |
| `notion` | Create and search Notion pages |
| `social` | Instagram & Reddit management |
| `update` | Self-update to the latest version |
| `update` | Self-update to the latest version |

## 🏗 Build & Release

This project uses [GoReleaser](https://goreleaser.com) for automated cross-platform builds and GitHub releases.

### Local Build
```powershell
go build -o fastclaw.exe main.go
```

### Release Process
1. Commit changes to `main`
2. Create a version tag (e.g. `v1.1.0`)
3. Run GoReleaser to build and publish

```powershell
git add .
git commit -m "feat: description"
git tag v1.1.0
goreleaser release --clean
git push origin main --tags
```

GoReleaser builds binaries for **linux**, **darwin**, **windows** × **amd64**, **arm64** and publishes them to [GitHub Releases](https://github.com/nadmzda/fast-claw-cli/releases).

> Requires `GITHUB_TOKEN` environment variable set for publishing.

## 🤝 Contributing
Contributions are welcome! Feel free to open issues or submit pull requests to help improve FastClaw CLI.

---
© 2026 Fast-Claw CLI. Built with Go and AI.
