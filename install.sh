#!/bin/bash
set -e

REPO="nadmzda/fast-claw-cli"
echo "🔍 Fetching latest version info from GitHub..."
LATEST_RELEASE=$(curl -s https://api.github.com/repos/$REPO/releases/latest)
TAG_NAME=$(echo "$LATEST_RELEASE" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
VERSION_NUM=${TAG_NAME#v}

OS_TYPE=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH_RAW=$(uname -m)

if [ "$ARCH_RAW" = "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH_RAW" = "aarch64" ] || [ "$ARCH_RAW" = "arm64" ]; then
    ARCH="arm64"
else
    echo "❌ Unsupported architecture: $ARCH_RAW"
    exit 1
fi

ASSET_NAME="fast-claw-cli_${VERSION_NUM}_${OS_TYPE}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$TAG_NAME/$ASSET_NAME"

echo "🚀 Downloading $TAG_NAME for $OS_TYPE/$ARCH..."
curl -L "$DOWNLOAD_URL" -o fastclaw.tar.gz

echo "📦 Extracting files..."
tar -xzf fastclaw.tar.gz
chmod +x fastclaw

if [ -w "/usr/local/bin" ]; then
    mv fastclaw /usr/local/bin/
else
    echo "⚠️ Requires sudo to move binary to /usr/local/bin"
    sudo mv fastclaw /usr/local/bin/
fi

rm fastclaw.tar.gz
echo "✨ FastClaw CLI $TAG_NAME installed successfully to /usr/local/bin/fastclaw"
