#!/bin/bash
set -e

REPO="nadmzda/fast-claw-cli"
echo "🔍 Fetching latest version info from GitHub..."
LATEST_RELEASE=$(curl -s https://api.github.com/repos/$REPO/releases/latest)
TAG_NAME=$(echo "$LATEST_RELEASE" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
VERSION_NUM=${TAG_NAME#v}

ARCH_RAW=$(uname -m)

if [ "$ARCH_RAW" = "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH_RAW" = "arm64" ]; then
    ARCH="arm64"
else
    echo "❌ Unsupported architecture: $ARCH_RAW"
    exit 1
fi

OS_TYPE="darwin"
ASSET_NAME="fast-claw-cli_${VERSION_NUM}_${OS_TYPE}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$TAG_NAME/$ASSET_NAME"

echo "🚀 Downloading $TAG_NAME for macOS/$ARCH..."
curl -L "$DOWNLOAD_URL" -o fastclaw.tar.gz

echo "📦 Extracting files..."
tar -xzf fastclaw.tar.gz
chmod +x fastclaw

if command -v brew &> /dev/null; then
    BIN_DIR="$(brew --prefix)/bin"
else
    BIN_DIR="/usr/local/bin"
fi

mkdir -p "$BIN_DIR"

if [ -w "$BIN_DIR" ]; then
    mv fastclaw "$BIN_DIR/"
else
    echo "⚠️ Requires sudo to move binary to $BIN_DIR"
    sudo mv fastclaw "$BIN_DIR/"
fi

rm fastclaw.tar.gz
echo "✨ FastClaw CLI $TAG_NAME installed successfully to $BIN_DIR/fastclaw"
