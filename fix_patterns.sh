#!/bin/bash
# Script to fix JWT hardcoded userID patterns

fix_createdby_pattern() {
    local file=$1
    local object_name=$2
    
    # Replace the pattern for CreatedBy
    sed -i "/从JWT中获取当前用户ID/{
        n
        s/${object_name}\.CreatedBy = uint(1)/userID, err := utils.GetCurrentUserID(ctx)\n\tif err != nil {\n\t\tctx.JSON(http.StatusUnauthorized, gin.H{\"error\": err.Error()})\n\t\treturn\n\t}\n\t${object_name}.CreatedBy = userID/
    }" "$file"
}

fix_approver_pattern() {
    local file=$1
    
    # Replace the pattern for approverID
    sed -i "/从JWT中获取当前用户ID/{
        n
        s/approverID := uint(1)/approverID, err := utils.GetCurrentUserID(ctx)\n\tif err != nil {\n\t\tctx.JSON(http.StatusUnauthorized, gin.H{\"error\": err.Error()})\n\t\treturn\n\t}/
    }" "$file"
}

# Fix attendance controller
echo "Fixing attendance controller..."
fix_createdby_pattern "controller/attendance.go" "rule"
fix_approver_pattern "controller/attendance.go"

echo "Attendance controller fixed!"