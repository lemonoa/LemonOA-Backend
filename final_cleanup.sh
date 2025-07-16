#!/bin/bash
# Final cleanup script for JWT user ID extraction

echo "=== Final Cleanup of JWT User ID Extraction ==="

# 1. Remove utils imports from controllers that don't use them
echo "Step 1: Removing unused utils imports..."

CONTROLLERS_TO_CHECK=("asset" "meeting" "notice" "notification" "seal" "system" "todo" "workflow")

for controller in "${CONTROLLERS_TO_CHECK[@]}"; do
    if ! grep -q "utils\.GetCurrentUserID" controller/${controller}.go; then
        echo "Removing unused utils import from ${controller}.go"
        sed -i '/github.com\/lemonoa\/LemonOA-Go\/utils/d' controller/${controller}.go
    else
        echo "Keeping utils import in ${controller}.go (it's used)"
    fi
done

# 2. Fix remaining hardcoded patterns for controllers that do use utils
echo "Step 2: Fixing hardcoded patterns in controllers that use utils..."

# Only process files that actually use utils
for controller in "${CONTROLLERS_TO_CHECK[@]}"; do
    if grep -q "utils\.GetCurrentUserID" controller/${controller}.go; then
        echo "Controller ${controller}.go already uses utils correctly"
    elif grep -q "uint(1)" controller/${controller}.go; then
        echo "Controller ${controller}.go needs pattern fixes but doesn't have utils import"
        # Re-add utils import for controllers that need it
        sed -i '/github.com\/lemonoa\/LemonOA-Go\/service/a\
	"github.com/lemonoa/LemonOA-Go/utils"' controller/${controller}.go
        echo "Re-added utils import to ${controller}.go"
    fi
done

echo "=== Cleanup Complete ==="

# Test build
echo "Step 3: Testing build..."
if go build; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed, manual fixes needed"
fi