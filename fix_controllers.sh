#!/bin/bash
# Script to fix JWT TODO comments in controllers

CONTROLLERS=("attendance" "todo" "notification" "meeting" "workflow" "asset" "notice" "seal" "system")

for controller in "${CONTROLLERS[@]}"; do
    echo "Fixing controller: $controller"
    
    # Add utils import
    sed -i '/import (/,/)/ {
        /github.com\/lemonoa\/LemonOA-Go\/service/a\
	"github.com/lemonoa/LemonOA-Go/utils"
    }' controller/${controller}.go
    
    # Change TODO comments
    sed -i 's|// TODO: 从JWT中获取当前用户ID|// 从JWT中获取当前用户ID|g' controller/${controller}.go
    
    echo "Fixed imports and TODO comments for $controller"
done

echo "All controllers updated!"