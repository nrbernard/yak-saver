#!/bin/bash
# Setup script for conventional commits git hook

set -e

HOOK_SOURCE="git-hooks/commit-msg"
HOOK_DEST=".git/hooks/commit-msg"

if [ ! -f "$HOOK_SOURCE" ]; then
    echo "❌ Error: $HOOK_SOURCE not found!"
    exit 1
fi

if [ ! -d ".git" ]; then
    echo "❌ Error: Not a git repository!"
    exit 1
fi

cp "$HOOK_SOURCE" "$HOOK_DEST"
chmod +x "$HOOK_DEST"

echo "✅ Conventional commits hook installed successfully!"
echo ""
echo "Your commit messages will now be validated against the Conventional Commits specification."
echo "See https://www.conventionalcommits.org/en/v1.0.0/ for details."

