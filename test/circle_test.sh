#!/bin/bash

# 圈子接口测试脚本
BASE_URL="http://localhost:9000"
TOKEN=""

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 登录并获取token
login() {
  echo -e "${BLUE}===== 登录获取TOKEN =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/auth/v1/login" \
    -H "Content-Type: application/json" \
    -d '{"username":"wangwu","password":"admin123"}')
  
  # 检查登录是否成功
  code=$(echo $response | grep -o '"code":[0-9]*' | cut -d':' -f2)
  
  if [ "$code" = "0" ]; then
    # 提取token
    TOKEN=$(echo $response | grep -o '"token":"[^"]*' | cut -d'"' -f4)
    echo -e "${GREEN}登录成功，TOKEN: ${TOKEN:0:15}...${NC}"
  else
    message=$(echo $response | grep -o '"message":"[^"]*' | cut -d'"' -f4)
    echo -e "${RED}登录失败: $message${NC}"
    exit 1
  fi
}

# 获取圈子列表
get_circle_list() {
  echo -e "\n${BLUE}===== 获取圈子列表 =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/circle/v1/list?page=1&size=10" \
    -H "Authorization: Bearer $TOKEN")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
}

# 获取圈子详情
get_circle_detail() {
  circle_id=$1
  echo -e "\n${BLUE}===== 获取圈子详情 (ID: $circle_id) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/circle/v1/detail?circleId=$circle_id" \
    -H "Authorization: Bearer $TOKEN")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
}

# 关注/取消关注圈子
toggle_circle_join() {
  circle_id=$1
  echo -e "\n${BLUE}===== 关注/取消关注圈子 (ID: $circle_id) =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/circle/v1/join" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{\"circleId\":$circle_id}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
}

# 主流程
main() {
  login
  get_circle_list
  
  # 假设有ID为1的圈子可供测试
  circle_id=1
  get_circle_detail $circle_id
  
  # 测试关注/取消关注
  toggle_circle_join $circle_id
  
  # 再次查看详情，确认关注状态变化
  get_circle_detail $circle_id
  
  # 再次执行取消关注
  toggle_circle_join $circle_id
  
  # 最终查看详情，确认关注状态恢复
  get_circle_detail $circle_id
}

# 执行主流程
main 