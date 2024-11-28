#!/bin/sh

# Define color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

bookmark() {
    go run cmd/bookmark-cli/main.go
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Bookmark-cli successfully started${NC}"
    else
        echo -e "${RED}❌ Failed to start Bookmark-cli${NC}"
    fi
}

bookmark