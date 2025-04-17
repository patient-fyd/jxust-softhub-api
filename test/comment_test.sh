#!/bin/bash

# 评论接口测试脚本
BASE_URL="http://localhost:9000"
TOKEN=""
TEST_POST_ID=1  # 测试用的帖子ID
COMMENT_ID=""   # 用于存储创建的评论ID

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 登录并获取token
login() {
  echo -e "${BLUE}===== 登录获取TOKEN =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/auth/v1/login" \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"admin123"}')
  
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

# 获取帖子评论列表
get_post_comments() {
  post_id=$1
  echo -e "\n${BLUE}===== 获取帖子(ID:${post_id})的评论列表 =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/comment/v1/list?contentType=post&contentId=${post_id}&page=1&size=10" \
    -H "Authorization: Bearer $TOKEN")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
}

# 创建帖子评论
create_post_comment() {
  post_id=$1
  content=$2
  echo -e "\n${BLUE}===== 创建帖子(ID:${post_id})评论 =====${NC}"
  echo -e "评论内容: $content"
  
  response=$(curl -s -X POST "${BASE_URL}/api/comment/v1/create" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{\"contentType\":\"post\",\"contentId\":${post_id},\"content\":\"${content}\"}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
  
  # 提取评论ID
  code=$(echo $response | grep -o '"code":[0-9]*' | cut -d':' -f2)
  if [ "$code" = "0" ]; then
    COMMENT_ID=$(echo $response | grep -o '"commentId":[0-9]*' | cut -d':' -f2)
    echo -e "${GREEN}评论创建成功，评论ID: $COMMENT_ID${NC}"
  else
    message=$(echo $response | grep -o '"message":"[^"]*' | cut -d'"' -f4)
    echo -e "${RED}评论创建失败: $message${NC}"
  fi
}

# 删除评论
delete_comment() {
  comment_id=$1
  echo -e "\n${BLUE}===== 删除评论(ID:${comment_id}) =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/comment/v1/delete" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{\"commentId\":${comment_id}}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
  
  # 检查删除是否成功
  code=$(echo $response | grep -o '"code":[0-9]*' | cut -d':' -f2)
  if [ "$code" = "0" ]; then
    success=$(echo $response | grep -o '"success":\(true\|false\)' | cut -d':' -f2)
    if [ "$success" = "true" ]; then
      echo -e "${GREEN}评论删除成功${NC}"
    else
      echo -e "${RED}评论删除失败${NC}"
    fi
  else
    message=$(echo $response | grep -o '"message":"[^"]*' | cut -d'"' -f4)
    echo -e "${RED}评论删除失败: $message${NC}"
  fi
}

# 测试敏感词过滤
test_sensitive_word() {
  echo -e "\n${BLUE}===== 测试敏感词过滤 =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/comment/v1/create" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{\"contentType\":\"post\",\"contentId\":${TEST_POST_ID},\"content\":\"这是一条包含敏感词的评论\"}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
}

# 测试参数校验 - 内容类型为空
test_empty_content_type() {
  echo -e "\n${BLUE}===== 测试参数校验 - 内容类型为空 =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/comment/v1/create" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{\"contentType\":\"\",\"contentId\":${TEST_POST_ID},\"content\":\"测试内容类型为空\"}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
  
  code=$(echo $response | grep -o '"code":[0-9]*' | cut -d':' -f2)
  if [ "$code" != "0" ]; then
    echo -e "${GREEN}参数校验测试通过，API正确拒绝了内容类型为空的请求${NC}"
  else
    echo -e "${RED}参数校验测试失败，API接受了内容类型为空的请求${NC}"
  fi
}

# 测试参数校验 - 内容ID为0
test_zero_content_id() {
  echo -e "\n${BLUE}===== 测试参数校验 - 内容ID为0 =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/comment/v1/create" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{\"contentType\":\"post\",\"contentId\":0,\"content\":\"测试内容ID为0\"}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
  
  code=$(echo $response | grep -o '"code":[0-9]*' | cut -d':' -f2)
  if [ "$code" != "0" ]; then
    echo -e "${GREEN}参数校验测试通过，API正确拒绝了内容ID为0的请求${NC}"
  else
    echo -e "${RED}参数校验测试失败，API接受了内容ID为0的请求${NC}"
  fi
}

# 测试参数校验 - 评论内容为空
test_empty_comment_content() {
  echo -e "\n${BLUE}===== 测试参数校验 - 评论内容为空 =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/comment/v1/create" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{\"contentType\":\"post\",\"contentId\":${TEST_POST_ID},\"content\":\"\"}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
  
  code=$(echo $response | grep -o '"code":[0-9]*' | cut -d':' -f2)
  if [ "$code" != "0" ]; then
    echo -e "${GREEN}参数校验测试通过，API正确拒绝了评论内容为空的请求${NC}"
  else
    echo -e "${RED}参数校验测试失败，API接受了评论内容为空的请求${NC}"
  fi
}

# 测试未登录时创建评论
test_create_without_login() {
  echo -e "\n${BLUE}===== 测试未登录时创建评论 =====${NC}"
  
  response=$(curl -s -X POST "${BASE_URL}/api/comment/v1/create" \
    -H "Content-Type: application/json" \
    -d "{\"contentType\":\"post\",\"contentId\":${TEST_POST_ID},\"content\":\"未登录测试评论\"}")
  
  # 格式化输出JSON
  echo $response | python3 -m json.tool
  
  code=$(echo $response | grep -o '"code":[0-9]*' | cut -d':' -f2)
  if [ "$code" != "0" ]; then
    echo -e "${GREEN}权限校验测试通过，API正确拒绝了未登录用户的评论创建请求${NC}"
  else
    echo -e "${RED}权限校验测试失败，API接受了未登录用户的评论创建请求${NC}"
  fi
}

# 主流程
main() {
  echo -e "${YELLOW}开始测试评论模块API...${NC}"
  
  # 先测试未登录状态
  test_create_without_login
  
  # 登录
  login
  
  # 获取帖子评论列表
  get_post_comments $TEST_POST_ID
  
  # 参数校验测试
  test_empty_content_type
  test_zero_content_id
  test_empty_comment_content
  
  # 创建评论
  create_post_comment $TEST_POST_ID "这是一条测试评论，由测试脚本自动创建于 $(date)"
  
  # 再次获取评论列表，查看新评论是否已添加
  get_post_comments $TEST_POST_ID
  
  # 测试敏感词过滤（如果系统配置了敏感词）
  test_sensitive_word
  
  # 如果创建评论成功，则删除该评论
  if [ ! -z "$COMMENT_ID" ]; then
    delete_comment $COMMENT_ID
    
    # 再次获取评论列表，确认评论已被删除
    get_post_comments $TEST_POST_ID
  fi
  
  echo -e "\n${YELLOW}评论模块API测试完成${NC}"
}

# 执行主流程
main 