#!/bin/bash

# 博客接口测试脚本 - 无需认证版本
BASE_URL="http://localhost:9000"

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

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

# 获取博客列表
get_blog_list() {
  echo -e "\n${BLUE}===== 获取博客列表 =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/list?page=1&size=10")
  
  # 检查响应状态
  check_response "$response" "获取博客列表"
  
  # 显示响应内容
  show_response "$response"
}

# 获取推荐博客
get_recommended_blogs() {
  echo -e "\n${BLUE}===== 获取推荐博客 =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/list?page=1&size=5&isRecommend=1")
  
  # 检查响应状态
  check_response "$response" "获取推荐博客"
  
  # 显示响应内容
  show_response "$response"
}

# 按分类获取博客
get_blogs_by_category() {
  category=$1
  echo -e "\n${BLUE}===== 获取分类博客 (分类: $category) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/list?page=1&size=10&category=$category")
  
  # 检查响应状态
  check_response "$response" "按分类获取博客 ($category)"
  
  # 显示响应内容
  show_response "$response"
}

# 按标签获取博客
get_blogs_by_tag() {
  tag=$1
  echo -e "\n${BLUE}===== 获取标签博客 (标签: $tag) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/list?page=1&size=10&tag=$tag")
  
  # 检查响应状态
  check_response "$response" "按标签获取博客 ($tag)"
  
  # 显示响应内容
  show_response "$response"
}

# 搜索博客
search_blogs() {
  keyword=$1
  echo -e "\n${BLUE}===== 搜索博客 (关键词: $keyword) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/list?page=1&size=10&keyword=$keyword")
  
  # 检查响应状态
  check_response "$response" "搜索博客 ($keyword)"
  
  # 显示响应内容
  show_response "$response"
}

# 获取博客详情
get_blog_detail() {
  blog_id=$1
  echo -e "\n${BLUE}===== 获取博客详情 (ID: $blog_id) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/detail?blogId=$blog_id")
  
  # 检查响应状态
  check_response "$response" "获取博客详情 (ID: $blog_id)"
  
  # 显示响应内容
  show_response "$response"
}

# 获取博客分类列表
get_blog_categories() {
  echo -e "\n${BLUE}===== 获取博客分类列表 =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/category/list")
  
  # 检查响应状态
  check_response "$response" "获取博客分类列表"
  
  # 显示响应内容
  show_response "$response"
}

# 获取博客标签列表
get_blog_tags() {
  echo -e "\n${BLUE}===== 获取博客标签列表 =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/tag/list")
  
  # 检查响应状态
  check_response "$response" "获取博客标签列表"
  
  # 显示响应内容
  show_response "$response"
}

# 获取博客评论列表
get_blog_comments() {
  blog_id=$1
  echo -e "\n${BLUE}===== 获取博客评论列表 (博客ID: $blog_id) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/comment/list?blogId=$blog_id&page=1&size=10")
  
  # 检查响应状态
  check_response "$response" "获取博客评论列表 (博客ID: $blog_id)"
  
  # 显示响应内容
  show_response "$response"
}

# 获取作者的博客列表
get_author_blogs() {
  author_id=$1
  echo -e "\n${BLUE}===== 获取作者博客列表 (作者ID: $author_id) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/blog/v1/blog/list?page=1&size=10&authorId=$author_id")
  
  # 检查响应状态
  check_response "$response" "获取作者博客列表 (作者ID: $author_id)"
  
  # 显示响应内容
  show_response "$response"
}

# 测试主函数
run_tests() {
  # 检查服务器是否运行
  check_server
  
  echo -e "\n${BLUE}===== 开始测试博客API (无需认证) =====${NC}"
  
  # 基础API测试
  get_blog_list
  get_recommended_blogs
  get_blog_categories
  get_blog_tags
  
  # 搜索测试
  search_blogs "测试"
  
  # 默认分类测试
  get_blogs_by_category "技术分享"
  
  # 如果提供了博客ID，则测试详情和评论
  if [ -n "$1" ]; then
    get_blog_detail $1
    get_blog_comments $1
  else
    echo -e "\n${YELLOW}提示: 要测试获取博客详情和评论，请提供博客ID作为参数${NC}"
    echo -e "${YELLOW}例如: ./blog_test_public.sh 1${NC}"
  fi
  
  echo -e "\n${GREEN}测试完成!${NC}"
}

# 开始执行测试
run_tests $1 