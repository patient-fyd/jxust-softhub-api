#!/bin/bash

# 博客API集成测试脚本
BASE_DIR=$(dirname "$0")

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# 显示帮助信息
print_usage() {
  echo -e "${BLUE}博客API集成测试脚本${NC}"
  echo -e "用法:"
  echo -e "  $0 [选项]"
  echo -e "选项:"
  echo -e "  --public   仅测试无需认证的API"
  echo -e "  --auth     仅测试需要认证的API"
  echo -e "  --all      测试所有API（默认）"
  echo -e "  --help     显示帮助信息"
  echo -e "示例:"
  echo -e "  $0 --public                    # 测试公开API"
  echo -e "  $0 --auth admin 123456         # 测试需认证的API，提供用户名密码"
  echo -e "  $0 --auth admin 123456 1       # 测试需认证的API，使用特定博客ID"
  echo -e "  $0 --all admin 123456          # 测试所有API，提供用户名密码"
  echo -e "  $0 --all admin 123456 1        # 测试所有API，使用特定博客ID"
}

# 运行测试
run_tests() {
  if [ "$mode" = "public" ] || [ "$mode" = "all" ]; then
    echo -e "\n${BLUE}===== 开始测试公开API =====${NC}"
    $BASE_DIR/blog_test_public.sh
    echo -e "\n${GREEN}公开API测试完成${NC}"
  fi
  
  if [ "$mode" = "auth" ] || [ "$mode" = "all" ]; then
    if [ -z "$username" ] || [ -z "$password" ]; then
      echo -e "\n${RED}错误: 测试需认证的API需要提供用户名和密码${NC}"
      echo -e "示例: $0 --auth admin 123456"
      exit 1
    fi
    
    echo -e "\n${BLUE}===== 开始测试需认证的API =====${NC}"
    
    if [ -n "$blog_id" ]; then
      $BASE_DIR/blog_test_auth.sh "$username" "$password" "$blog_id"
    else
      $BASE_DIR/blog_test_auth.sh "$username" "$password"
    fi
    
    echo -e "\n${GREEN}需认证的API测试完成${NC}"
  fi
  
  echo -e "\n${GREEN}所有测试已完成!${NC}"
}

# 参数解析
mode="all"
username=""
password=""
blog_id=""

case "$1" in
  --public)
    mode="public"
    shift
    ;;
  --auth)
    mode="auth"
    shift
    username="$1"
    password="$2"
    blog_id="$3"
    ;;
  --all)
    mode="all"
    shift
    username="$1"
    password="$2"
    blog_id="$3"
    ;;
  --help)
    print_usage
    exit 0
    ;;
  "")
    # 默认使用all模式
    ;;
  *)
    echo -e "${RED}错误: 未知选项 $1${NC}"
    print_usage
    exit 1
    ;;
esac

# 检查文件权限
for file in "$BASE_DIR/blog_test_public.sh" "$BASE_DIR/blog_test_auth.sh"; do
  if [ ! -x "$file" ]; then
    echo -e "${YELLOW}为测试脚本添加执行权限: $file${NC}"
    chmod +x "$file"
  fi
done

# 确认启动测试
echo -e "${BLUE}===============================================================${NC}"
echo -e "${BLUE}           技术博客 API 测试 - 测试模式: $mode${NC}"
echo -e "${BLUE}===============================================================${NC}"

if [ "$mode" = "auth" ] || [ "$mode" = "all" ]; then
  echo -e "${YELLOW}用户名: $username${NC}"
  echo -e "${YELLOW}密码: ******${NC}"
  if [ -n "$blog_id" ]; then
    echo -e "${YELLOW}博客ID: $blog_id${NC}"
  fi
fi

echo -e -n "\n${BLUE}是否开始测试? (y/n) ${NC}"
read -r confirm

if [ "$confirm" = "y" ] || [ "$confirm" = "Y" ]; then
  run_tests
else
  echo -e "${YELLOW}测试已取消${NC}"
  exit 0
fi 