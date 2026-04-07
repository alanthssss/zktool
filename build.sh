#!/bin/bash

set -e

APP_NAME="seed"
BUILD_DIR="build"
PLATFORMS=(
  "linux/amd64"
  "windows/amd64"
  "darwin/amd64"
  "darwin/arm64"
)

mkdir -p "$BUILD_DIR"
echo "📦 Building $APP_NAME for multiple platforms..."

for platform in "${PLATFORMS[@]}"; do
  IFS="/" read -r GOOS GOARCH <<< "$platform"
  output_name="$GOOS-$GOARCH/$APP_NAME"
  [[ "$GOOS" == "windows" ]] && output_name+=".exe"

  echo "🔧 Building for $GOOS/$GOARCH -> $output_name"
  GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -o "$BUILD_DIR/$output_name" ./main.go
	done

echo "✅ All builds complete. Files are in the $BUILD_DIR directory."
