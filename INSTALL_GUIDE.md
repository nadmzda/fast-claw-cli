# Installation Guide for FastClaw CLI

This guide provides automated and manual installation steps for **FastClaw CLI**, a powerful tool to interact with FastClaw API skills.

## 🔗 Project Resources
- **GitHub Repository:** [https://github.com/nadmzda/fast-claw-cli](https://github.com/nadmzda/fast-claw-cli)
- **Latest Releases:** [https://github.com/nadmzda/fast-claw-cli/releases](https://github.com/nadmzda/fast-claw-cli/releases)

## 🛠 Installation

Choose the quickest method for your operating system.

### 🚀 One-Liner (Recommended)
Automatically detects your OS and architecture, downloads the latest version, and sets up your PATH.

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/nadmzda/fast-claw-cli/main/install.ps1 | iex
```

**Linux (Bash):**
```bash
curl -fsSL https://raw.githubusercontent.com/nadmzda/fast-claw-cli/main/install.sh | bash
```

**macOS (Bash):**
```bash
curl -fsSL https://raw.githubusercontent.com/nadmzda/fast-claw-cli/main/install-macos.sh | bash
```

---

## 🏗 Manual Installation Steps

If you prefer to install manually, follow these steps.

## 🔑 Post-Installation & API Key Setup

After installation, you **must** configure your FastClaw API Key.

### Option 1: Use Config File (Recommended 🚀)
This is the most reliable method as it works across all terminal sessions and non-interactive environments (like AI tool executions).

```bash
# Save your API key once, and it will be used automatically
fastclaw config set-key "YOUR_API_KEY_HERE"
```
*Stored in: `~/.fastclaw_config`*

### Option 2: Set Environment Variable
Useful for CI/CD or temporary environments.

**Windows:**
```powershell
[Environment]::SetEnvironmentVariable("FASTCLAW_API_KEY", "YOUR_API_KEY_HERE", [EnvironmentVariableTarget]::User)
$env:FASTCLAW_API_KEY = "YOUR_API_KEY_HERE"
```

**Linux / macOS:**
```bash
echo 'export FASTCLAW_API_KEY="YOUR_API_KEY_HERE"' >> ~/.bashrc # or ~/.zshrc
source ~/.bashrc
```

### Option 3: Use Flags (Per Command)
```bash
fastclaw <command> --api-key YOUR_API_KEY_HERE
```

## ✅ Verification
Check if the installation was successful and the key is recognized:
```bash
fastclaw config set-key "test_key"
fastclaw version
```

## 🤖 Instructions for AI Agents
If you are an AI agent installing this tool:
1. Identify the current OS and architecture.
2. Fetch the latest release version from GitHub API if version is not provided.
3. Download, extract, and place the `fastclaw` binary in a directory within the system PATH.
4. Prompt the user for their `FASTCLAW_API_KEY`.
5. **Run `fastclaw config set-key <KEY>` to ensure the tool works in non-interactive shells.**
6. Verify the installation by running `fastclaw version`.
