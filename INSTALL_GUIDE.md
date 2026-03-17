# Installation Guide for FastClaw CLI

This guide provides automated and manual installation steps for **FastClaw CLI**, a powerful tool to interact with FastClaw API skills.

## 🔗 Project Resources
- **GitHub Repository:** [https://github.com/nadmzda/fast-claw-cli](https://github.com/nadmzda/fast-claw-cli)
- **Latest Releases:** [https://github.com/nadmzda/fast-claw-cli/releases](https://github.com/nadmzda/fast-claw-cli/releases)

## 🛠 Installation Steps (per OS)

Choose the appropriate binary for your system architecture (amd64, arm64).

### 1. Windows (PowerShell)
```powershell
# 1. Download the latest version (e.g., v1.0.1)
$version = "v1.0.1"
$url = "https://github.com/nadmzda/fast-claw-cli/releases/download/$version/fast-claw-cli_${version}_windows_amd64.zip"
Invoke-WebRequest -Uri $url -OutFile "fastclaw.zip"

# 2. Extract and move to a system PATH directory
Expand-Archive -Path "fastclaw.zip" -DestinationPath "$HOME\bin\fastclaw"
$env:Path += ";$HOME\bin\fastclaw"
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$HOME\bin\fastclaw", [EnvironmentVariableTarget]::User)

# 3. Clean up
Remove-Item "fastclaw.zip"
```

### 2. Linux / macOS (Terminal)
```bash
# 1. Download the latest version (e.g., v1.0.1)
VERSION="v1.0.1"
# For Linux amd64:
URL="https://github.com/nadmzda/fast-claw-cli/releases/download/${VERSION}/fast-claw-cli_${VERSION}_linux_amd64.tar.gz"
# For macOS arm64 (M1/M2/M3):
# URL="https://github.com/nadmzda/fast-claw-cli/releases/download/${VERSION}/fast-claw-cli_${VERSION}_darwin_arm64.tar.gz"

curl -L $URL -o fastclaw.tar.gz

# 2. Extract and move to /usr/local/bin
tar -xzf fastclaw.tar.gz
chmod +x fastclaw
sudo mv fastclaw /usr/local/bin/

# 3. Clean up
rm fastclaw.tar.gz
```

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
