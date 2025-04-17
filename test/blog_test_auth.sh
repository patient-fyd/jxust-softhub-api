#!/bin/bash

# 博客接口测试脚本 - 需要认证版本
BASE_URL="http://localhost:9000"

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# 全局变量
TOKEN=""
BLOG_ID=""
COMMENT_ID=""

# 检查API服务是否启动
check_server() {
  echo -e "${BLUE}检查API服务是否启动...${NC}"
  
  # 使用curl检测服务器是否响应
  server_response=$(curl -s -o /dev/null -w "%{http_code}" "${BASE_URL}" 2>/dev/null)
  
  if [ "$server_response" = "000" ]; then
    echo -e "${RED}✗ API服务未启动或无法连接，请先启动服务器${NC}"
    echo -e "${YELLOW}提示: 可以通过以下命令启动服务:${NC}"
    echo -e "${YELLOW}go run main.go${NC}"
    exit 1
  else
    echo -e "${GREEN}✓ API服务已启动 (HTTP状态码: $server_response)${NC}"
  fi
}

# 检查响应状态
check_response() {
  response=$1
  endpoint=$2
  
  # 检查响应是否为空或非JSON格式
  if [ -z "$response" ] || ! echo "$response" | jq . &>/dev/null; then
    echo -e "${RED}✗ $endpoint 响应为空或格式错误${NC}"
    echo "原始响应: $response"
    return 1
  fi
  
  # 检查是否包含code字段且值为0
  code=$(echo $response | jq -r '.code')
  if [ "$code" = "0" ]; then
    echo -e "${GREEN}✓ $endpoint 请求成功 (code: $code)${NC}"
    return 0
  else
    message=$(echo $response | jq -r '.msg // "未知错误"')
    echo -e "${RED}✗ $endpoint 请求失败 (code: $code, message: $message)${NC}"
    return 1
  fi
}

# 显示响应内容
show_response() {
  response=$1
  
  # 检查响应是否为空或非JSON格式
  if [ -z "$response" ] || ! echo "$response" | jq . &>/dev/null; then
    echo "原始响应: $response"
    return
  fi
  
  # 尝试获取data部分
  data=$(echo $response | jq '.data')
  if [ "$data" != "null" ] && [ "$data" != "" ]; then
    echo $response | jq '.data'
  else
    # 如果data为空，显示完整响应
    echo $response | jq .
  fi
}

# 用户登录获取token
login() {
  echo -e "\n${BLUE}===== 用户登录 =====${NC}"
  
  # 修改这里的用户名和密码
  username=$1
  password=$2
  
  if [ -z "$username" ] || [ -z "$password" ]; then
    echo -e "${RED}✗ 用户名或密码不能为空${NC}"
    exit 1
  fi
  
  response=$(curl -s -X POST "${BASE_URL}/api/user/v1/login" \
    -H "Content-Type: application/json" \
    -d "{\"userName\":\"$username\",\"password\":\"$password\"}")
  
  # 检查响应状态
  check_response "$response" "用户登录"
  
  # 从响应中提取token
  TOKEN=$(echo $response | jq -r '.data.token')
  
  if [ -z "$TOKEN" ] || [ "$TOKEN" = "null" ]; then
    echo -e "${RED}✗ 获取token失败${NC}"
    exit 1
  else
    echo -e "${GREEN}✓ 获取token成功${NC}"
  fi
}

# 创建博客
create_blog() {
  echo -e "\n${BLUE}===== 创建博客 =====${NC}"
  
  title=$1
  content=$2
  category=$3
  tags=$4
  status=$5
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/create" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"title\": \"$title\",
      \"content\": \"$content\",
      \"summary\": \"这是一篇测试博客的摘要\",
      \"category\": \"$category\",
      \"tags\": \"$tags\",
      \"status\": $status
    }")
  
  # 检查响应状态
  check_response "$response" "创建博客"
  
  # 显示响应内容
  show_response "$response"
  
  # 提取博客ID
  BLOG_ID=$(echo $response | jq -r '.data.blogId')
  
  if [ -z "$BLOG_ID" ] || [ "$BLOG_ID" = "null" ]; then
    echo -e "${RED}✗ 获取博客ID失败${NC}"
  else
    echo -e "${GREEN}✓ 创建博客成功，博客ID: $BLOG_ID${NC}"
  fi
}

# 更新博客
update_blog() {
  echo -e "\n${BLUE}===== 更新博客 =====${NC}"
  
  blog_id=$1
  title=$2
  content=$3
  category=$4
  tags=$5
  status=$6
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/update" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"blogId\": $blog_id,
      \"title\": \"$title\",
      \"content\": \"$content\",
      \"summary\": \"这是一篇测试博客的更新摘要\",
      \"category\": \"$category\",
      \"tags\": \"$tags\",
      \"status\": $status
    }")
  
  # 检查响应状态
  check_response "$response" "更新博客"
  
  # 显示响应内容
  show_response "$response"
}

# 设置博客推荐状态
set_blog_recommend() {
  echo -e "\n${BLUE}===== 设置博客推荐状态 =====${NC}"
  
  blog_id=$1
  is_recommend=$2
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/recommend" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"blogId\": $blog_id,
      \"isRecommend\": $is_recommend
    }")
  
  # 检查响应状态
  check_response "$response" "设置博客推荐状态"
  
  # 显示响应内容
  show_response "$response"
}

# 点赞博客
like_blog() {
  echo -e "\n${BLUE}===== 点赞博客 =====${NC}"
  
  blog_id=$1
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/like" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"blogId\": $blog_id
    }")
  
  # 检查响应状态
  check_response "$response" "点赞博客"
  
  # 显示响应内容
  show_response "$response"
}

# 取消点赞博客
unlike_blog() {
  echo -e "\n${BLUE}===== 取消点赞博客 =====${NC}"
  
  blog_id=$1
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/unlike" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"blogId\": $blog_id
    }")
  
  # 检查响应状态
  check_response "$response" "取消点赞博客"
  
  # 显示响应内容
  show_response "$response"
}

# 创建博客评论
create_blog_comment() {
  echo -e "\n${BLUE}===== 创建博客评论 =====${NC}"
  
  blog_id=$1
  content=$2
  parent_id=$3
  
  request_data="{
    \"blogId\": $blog_id,
    \"content\": \"$content\"
  }"
  
  # 如果有父评论ID，则添加到请求数据中
  if [ -n "$parent_id" ] && [ "$parent_id" != "null" ] && [ "$parent_id" != "0" ]; then
    request_data=$(echo $request_data | jq ". + {\"parentId\": $parent_id}")
  fi
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/comment/create" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "$request_data")
  
  # 检查响应状态
  check_response "$response" "创建博客评论"
  
  # 显示响应内容
  show_response "$response"
  
  # 提取评论ID
  COMMENT_ID=$(echo $response | jq -r '.data.commentId')
  
  if [ -z "$COMMENT_ID" ] || [ "$COMMENT_ID" = "null" ]; then
    echo -e "${RED}✗ 获取评论ID失败${NC}"
  else
    echo -e "${GREEN}✓ 创建评论成功，评论ID: $COMMENT_ID${NC}"
  fi
}

# 获取博客评论列表
get_blog_comments() {
  echo -e "\n${BLUE}===== 获取博客评论列表 =====${NC}"
  
  blog_id=$1
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/comment/list?blogId=$blog_id&page=1&size=10" \
    -H "Authorization: Bearer $TOKEN")
  
  # 检查响应状态
  check_response "$response" "获取博客评论列表"
  
  # 显示响应内容
  show_response "$response"
}

# 删除博客评论
delete_blog_comment() {
  echo -e "\n${BLUE}===== 删除博客评论 =====${NC}"
  
  comment_id=$1
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/comment/delete" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"commentId\": $comment_id
    }")
  
  # 检查响应状态
  check_response "$response" "删除博客评论"
  
  # 显示响应内容
  show_response "$response"
}

# 删除博客
delete_blog() {
  echo -e "\n${BLUE}===== 删除博客 =====${NC}"
  
  blog_id=$1
  
  response=$(curl -s -X POST "${BASE_URL}/api/blog/v1/blog/delete" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d "{
      \"blogId\": $blog_id
    }")
  
  # 检查响应状态
  check_response "$response" "删除博客"
  
  # 显示响应内容
  show_response "$response"
}

# 测试流程
run_tests() {
  # 参数
  username=$1
  password=$2
  blog_id=$3
  
  # 检查服务器状态
  check_server
  
  # 登录获取Token
  login "$username" "$password"
  
  if [ -n "$blog_id" ]; then
    # 使用传入的博客ID进行测试
    BLOG_ID=$blog_id
    echo -e "${YELLOW}使用指定博客ID: $BLOG_ID 进行测试${NC}"
    
    # 无需创建新博客，直接从更新开始测试
    update_blog $BLOG_ID "更新的测试博客" "这是一篇更新的测试博客内容" "技术分享" "测试,接口测试" 1
    set_blog_recommend $BLOG_ID 1
    like_blog $BLOG_ID
    create_blog_comment $BLOG_ID "这是一条测试评论" 0
    get_blog_comments $BLOG_ID
    
    # 如果有评论ID，则删除评论
    if [ -n "$COMMENT_ID" ] && [ "$COMMENT_ID" != "null" ]; then
      delete_blog_comment $COMMENT_ID
    fi
    
    unlike_blog $BLOG_ID
    
    # 注意：不会删除传入的博客ID
    echo -e "${YELLOW}跳过删除博客步骤，因为使用的是指定博客ID${NC}"
  else
    # 创建新博客并测试完整流程
    create_blog "测试博客" "这是一篇测试博客内容" "技术分享" "测试,接口测试" 1
    
    if [ -z "$BLOG_ID" ] || [ "$BLOG_ID" = "null" ]; then
      echo -e "${RED}✗ 获取博客ID失败，无法继续测试${NC}"
      exit 1
    fi
    
    update_blog $BLOG_ID "更新的测试博客" "这是一篇更新的测试博客内容" "技术分享" "测试,接口测试" 1
    set_blog_recommend $BLOG_ID 1
    like_blog $BLOG_ID
    create_blog_comment $BLOG_ID "这是一条测试评论" 0
    get_blog_comments $BLOG_ID
    
    # 如果有评论ID，则尝试回复和删除评论
    if [ -n "$COMMENT_ID" ] && [ "$COMMENT_ID" != "null" ]; then
      create_blog_comment $BLOG_ID "这是一条回复评论" $COMMENT_ID
      delete_blog_comment $COMMENT_ID
    fi
    
    unlike_blog $BLOG_ID
    delete_blog $BLOG_ID
  fi
  
  echo -e "\n${GREEN}博客API测试完成!${NC}"
}

# 打印使用说明
print_usage() {
  echo -e "${BLUE}博客API测试脚本 - 需要认证版本${NC}"
  echo -e "用法:"
  echo -e "  $0 <用户名> <密码> [博客ID]"
  echo -e "参数:"
  echo -e "  用户名: 登录用户名，必填"
  echo -e "  密码: 登录密码，必填"
  echo -e "  博客ID: 可选，指定要测试的博客ID，如果不提供则会创建新博客"
  echo -e "示例:"
  echo -e "  $0 admin 123456"
  echo -e "  $0 admin 123456 1"
}

# 主函数
main() {
  if [ -z "$1" ] || [ -z "$2" ]; then
    print_usage
    exit 1
  fi
  
  # 执行测试
  run_tests "$1" "$2" "$3"
}

# 执行主函数
main "$@" 