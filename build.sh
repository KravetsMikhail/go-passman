#!/bin/bash

# Build script for go-passman
# Usage: ./build.sh [target] [version]
# Targets: linux, darwin, windows, all, release, clean

set -e

BINARY_NAME="go-passman"
OUTPUT_DIR="dist"
VERSION="${2:-0.3.0}"   # used by release target

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

build_target() {
    local GOOS=$1
    local GOARCH=$2
    local OUTPUT_NAME=$3
    
    echo -e "${BLUE}Building for ${GOOS}/${GOARCH} (slim)...${NC}"
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -trimpath -o "${OUTPUT_DIR}/${OUTPUT_NAME}"
    echo -e "${GREEN}✓ Built: ${OUTPUT_DIR}/${OUTPUT_NAME}${NC}"
}

clean() {
    echo -e "${BLUE}Cleaning...${NC}"
    rm -rf dist/
    echo -e "${GREEN}✓ Clean complete${NC}"
}

help() {
    echo "Usage: $0 [target] [version]"
    echo ""
    echo "Targets:"
    echo "  linux       - Build for Linux (x86_64)"
    echo "  darwin      - Build for macOS (Intel)"
    echo "  darwin-arm  - Build for macOS (Apple Silicon)"
    echo "  windows     - Build for Windows"
    echo "  all         - Build for all platforms"
    echo "  release     - Build for GitHub release (all platforms, version in name)"
    echo "  release 0.3.0 - Same with explicit version"
    echo "  clean       - Clean build artifacts"
    echo "  help        - Show this help message"
}

# Create output directory
mkdir -p "$OUTPUT_DIR"

case "${1:-help}" in
    linux)
        build_target "linux" "amd64" "${BINARY_NAME}-linux-amd64"
        ;;
    darwin)
        build_target "darwin" "amd64" "${BINARY_NAME}-darwin-amd64"
        ;;
    darwin-arm)
        build_target "darwin" "arm64" "${BINARY_NAME}-darwin-arm64"
        ;;
    windows)
        build_target "windows" "amd64" "${BINARY_NAME}-windows-amd64.exe"
        ;;
    all)
        build_target "linux" "amd64" "${BINARY_NAME}-linux-amd64"
        build_target "darwin" "amd64" "${BINARY_NAME}-darwin-amd64"
        build_target "darwin" "arm64" "${BINARY_NAME}-darwin-arm64"
        build_target "windows" "amd64" "${BINARY_NAME}-windows-amd64.exe"
        echo -e "${GREEN}✓ All builds complete!${NC}"
        ;;
    release)
        RELEASE_SUBDIR="release-${VERSION}"
        mkdir -p "${OUTPUT_DIR}/${RELEASE_SUBDIR}"
        echo -e "${BLUE}Building release ${VERSION} for GitHub...${NC}"
        build_target "windows" "amd64" "${RELEASE_SUBDIR}/${BINARY_NAME}-${VERSION}-windows-amd64.exe"
        build_target "windows" "arm64" "${RELEASE_SUBDIR}/${BINARY_NAME}-${VERSION}-windows-arm64.exe"
        build_target "linux" "amd64" "${RELEASE_SUBDIR}/${BINARY_NAME}-${VERSION}-linux-amd64"
        build_target "linux" "arm64" "${RELEASE_SUBDIR}/${BINARY_NAME}-${VERSION}-linux-arm64"
        build_target "linux" "386" "${RELEASE_SUBDIR}/${BINARY_NAME}-${VERSION}-linux-386"
        build_target "darwin" "amd64" "${RELEASE_SUBDIR}/${BINARY_NAME}-${VERSION}-darwin-amd64"
        build_target "darwin" "arm64" "${RELEASE_SUBDIR}/${BINARY_NAME}-${VERSION}-darwin-arm64"
        echo -e "${GREEN}✓ Release builds: ${OUTPUT_DIR}/${RELEASE_SUBDIR}/${NC}"
        echo "Upload the contents to GitHub Releases."
        ;;
    clean)
        clean
        ;;
    help)
        help
        ;;
    *)
        echo -e "${RED}Unknown target: $1${NC}"
        help
        exit 1
        ;;
esac
