#!/bin/bash

# Usage: ./tag-modules.sh v1.2.3

set -e

TAG="$1"

if [ -z "$TAG" ]; then
  echo "❌ Error: Tag is required (e.g., v1.0.0)"
  exit 1
fi

# Ensure on master
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [ "$CURRENT_BRANCH" != "master" ]; then
  echo "❌ Error: You must be on the 'master' branch (currently on '$CURRENT_BRANCH')"
  exit 1
fi

# Pull latest changes
echo "🔄 Pulling latest changes from origin..."
git pull origin master

# 🔧 List of module paths (relative to repo root)
# Use "." for root module
MODULES=(
  "."              # root module => tag: v1.2.3
  "parser/awsiam"    # => tag: parser/awsiam/v1.2.3
  "parser/clickhouse"         # => tag: parser/clickhouse/v1.2.3
)

# Fail if any tag already exists
for module in "${MODULES[@]}"; do
  if [ "$module" = "." ]; then
    TAG_NAME="$TAG"
  else
    TAG_NAME="${module}/${TAG}"
  fi

  if git rev-parse "$TAG_NAME" >/dev/null 2>&1; then
    echo "❌ Error: Tag '$TAG_NAME' already exists. Aborting."
    exit 1
  fi
done

for module in "${MODULES[@]}"; do
  if [ "$module" = "." ]; then
    TAG_NAME="$TAG"
  else
    TAG_NAME="${module}/${TAG}"
  fi

  echo "🏷️ Tagging $module with $TAG_NAME"
  git tag "$TAG_NAME"
done

echo "📤 Pushing tags..."
for module in "${MODULES[@]}"; do
  if [ "$module" = "." ]; then
    TAG_NAME="$TAG"
  else
    TAG_NAME="${module}/${TAG}"
  fi
  git push origin "$TAG_NAME"
done

echo "✅ All modules tagged and pushed!"
