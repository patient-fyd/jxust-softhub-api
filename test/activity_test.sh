#!/bin/bash

# 活动接口测试脚本
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
  code=$(echo "$response" | jq -r '.code')
  if [ "$code" = "0" ]; then
    echo -e "${GREEN}✓ $endpoint 请求成功 (code: $code)${NC}"
    return 0
  else
    message=$(echo "$response" | jq -r '.msg // "未知错误"')
    echo -e "${RED}✗ $endpoint 请求失败 (code: $code, message: $message)${NC}"
    echo "$response" | jq .
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
  data=$(echo "$response" | jq '.data')
  if [ "$data" != "null" ] && [ "$data" != "" ]; then
    echo "$response" | jq '.data'
  else
    # 如果data为空，显示完整响应
    echo "$response" | jq .
  fi
}

# 获取Token(用于需要认证的API)
token=""
login() {
  echo -e "\n${BLUE}===== 用户登录 =====${NC}"
  
  # 请替换为实际的用户名和密码
  response=$(curl -s -X POST "${BASE_URL}/api/auth/v1/login" \
    -H "Content-Type: application/json" \
    -d '{"username":"wangwu","password":"admin123"}')
  
  # 检查响应状态
  if check_response "$response" "用户登录"; then
    token=$(echo "$response" | jq -r '.data.token')
    if [ -z "$token" ] || [ "$token" = "null" ]; then
      echo -e "${RED}✗ 获取Token失败${NC}"
      return 1
    fi
    echo -e "${GREEN}✓ 登录成功，获取到Token: ${token:0:10}...${NC}"
    return 0
  else
    echo -e "${RED}✗ 登录失败，部分API可能无法测试${NC}"
    return 1
  fi
}

# 获取活动列表
get_activity_list() {
  echo -e "\n${BLUE}===== 获取活动列表 =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/activity/v1/list?page_num=1&page_size=10")
  
  # 检查响应状态
  check_response "$response" "获取活动列表"
  
  # 显示响应内容
  show_response "$response"
}

# 根据状态获取活动列表
get_activity_list_by_status() {
  status=$1
  status_name=$2
  echo -e "\n${BLUE}===== 获取${status_name}活动列表 (状态: $status) =====${NC}"
  
  response=$(curl -s -X GET "${BASE_URL}/api/activity/v1/list?page_num=1&page_size=10&status=$status")
  
  # 检查响应状态
  check_response "$response" "获取${status_name}活动列表"
  
  # 显示响应内容
  show_response "$response"
}

# 获取活动详情
get_activity_detail() {
  activity_id=$1
  echo -e "\n${BLUE}===== 获取活动详情 (ID: $activity_id) =====${NC}"
  
  if [ -z "$activity_id" ]; then
    echo -e "${RED}✗ 活动ID为空，无法获取详情${NC}"
    return 1
  fi
  
  response=$(curl -s -X GET "${BASE_URL}/api/activity/v1/detail/$activity_id")
  
  # 检查响应状态
  check_response "$response" "获取活动详情"
  
  # 显示响应内容
  show_response "$response"
  
  # 返回是否成功
  if [ "$(echo "$response" | jq -r '.code')" = "0" ]; then
    return 0
  else
    return 1
  fi
}

# 创建新活动
create_activity() {
  echo -e "\n${BLUE}===== 创建新活动 =====${NC}"
  
  # 检查是否已登录
  if [ -z "$token" ]; then
    echo -e "${YELLOW}需要先登录才能创建活动${NC}"
    login
    if [ -z "$token" ]; then
      return 1
    fi
  fi
  
  # 创建活动 - 跨平台的日期处理
  current_time=$(date +%s)
  
  # 使用跨平台的方式设置日期
  tomorrow=$(date -d "tomorrow" "+%Y-%m-%d" 2>/dev/null || date -v+1d "+%Y-%m-%d" 2>/dev/null)
  if [ -z "$tomorrow" ]; then
    # 如果上面的命令都失败，使用当前日期加上24小时
    tomorrow=$(date -d "@$(($(date +%s) + 86400))" "+%Y-%m-%d" 2>/dev/null || date "+%Y-%m-%d")
  fi
  
  start_time="${tomorrow} 10:00:00"
  end_time="${tomorrow} 18:00:00"
  title="测试活动-$current_time"
  
  echo -e "${YELLOW}正在创建活动：${title}${NC}"
  echo -e "${YELLOW}开始时间：${start_time}${NC}"
  echo -e "${YELLOW}结束时间：${end_time}${NC}"
  
  # 创建请求体
  request_body='{
    "title": "'"$title"'",
    "description": "这是一个用于测试的活动",
    "startTime": "'"$start_time"'",
    "endTime": "'"$end_time"'",
    "location": "软件学院102教室",
    "maxParticipants": 50
  }'
  
  echo -e "${YELLOW}请求内容：${NC}"
  echo "$request_body" | jq .
  
  response=$(curl -s -X POST "${BASE_URL}/api/activity/v1/create" \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: application/json" \
    -d "$request_body")
  
  # 检查响应状态
  if check_response "$response" "创建新活动"; then
    activity_id=$(echo "$response" | jq -r '.data.activityId')
    if [ -z "$activity_id" ] || [ "$activity_id" = "null" ]; then
      echo -e "${RED}✗ 无法获取新创建的活动ID${NC}"
      return 1
    fi
    echo -e "${GREEN}✓ 活动创建成功，ID: $activity_id${NC}"
    
    # 显示响应内容
    show_response "$response"
    
    # 返回活动ID
    echo "$activity_id"
    return 0
  else
    echo -e "${RED}✗ 活动创建失败${NC}"
    return 1
  fi
}

# 更新活动
update_activity() {
  activity_id=$1
  echo -e "\n${BLUE}===== 更新活动 (ID: $activity_id) =====${NC}"
  
  if [ -z "$activity_id" ]; then
    echo -e "${RED}✗ 活动ID为空，无法更新${NC}"
    return 1
  fi
  
  # 检查是否已登录
  if [ -z "$token" ]; then
    echo -e "${YELLOW}需要先登录才能更新活动${NC}"
    login
    if [ -z "$token" ]; then
      return 1
    fi
  fi
  
  # 更新活动
  request_body='{
    "activityId": '"$activity_id"',
    "title": "更新后的测试活动",
    "description": "这是一个经过更新的测试活动",
    "location": "软件学院103教室",
    "maxParticipants": 60
  }'
  
  response=$(curl -s -X PUT "${BASE_URL}/api/activity/v1/update" \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: application/json" \
    -d "$request_body")
  
  # 检查响应状态
  check_response "$response" "更新活动"
  
  # 显示响应内容
  show_response "$response"
  
  # 返回是否成功
  if [ "$(echo "$response" | jq -r '.code')" = "0" ]; then
    return 0
  else
    return 1
  fi
}

# 报名活动
register_activity() {
  activity_id=$1
  echo -e "\n${BLUE}===== 报名活动 (ID: $activity_id) =====${NC}"
  
  if [ -z "$activity_id" ]; then
    echo -e "${RED}✗ 活动ID为空，无法报名${NC}"
    return 1
  fi
  
  # 检查是否已登录
  if [ -z "$token" ]; then
    echo -e "${YELLOW}需要先登录才能报名活动${NC}"
    login
    if [ -z "$token" ]; then
      return 1
    fi
  fi
  
  # 随机生成学号和联系方式
  student_id="2021$(printf "%06d" $((RANDOM % 100000)))"
  contact="1351234$(printf "%04d" $((RANDOM % 10000)))"
  
  # 报名活动
  request_body='{
    "activityId": '"$activity_id"',
    "name": "测试用户",
    "studentId": "'"$student_id"'",
    "contact": "'"$contact"'"
  }'
  
  response=$(curl -s -X POST "${BASE_URL}/api/activity/v1/register" \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: application/json" \
    -d "$request_body")
  
  # 检查响应状态
  if check_response "$response" "报名活动"; then
    registration_id=$(echo "$response" | jq -r '.data.registrationId')
    if [ -z "$registration_id" ] || [ "$registration_id" = "null" ]; then
      echo -e "${RED}✗ 无法获取报名ID${NC}"
      return 1
    fi
    echo -e "${GREEN}✓ 活动报名成功，报名ID: $registration_id${NC}"
    
    # 显示响应内容
    show_response "$response"
    
    # 返回报名ID
    echo "$registration_id"
    return 0
  else
    echo -e "${RED}✗ 活动报名失败${NC}"
    return 1
  fi
}

# 获取报名列表
get_registration_list() {
  activity_id=$1
  echo -e "\n${BLUE}===== 获取活动报名列表 (活动ID: $activity_id) =====${NC}"
  
  if [ -z "$activity_id" ]; then
    echo -e "${RED}✗ 活动ID为空，无法获取报名列表${NC}"
    return 1
  fi
  
  # 检查是否已登录
  if [ -z "$token" ]; then
    echo -e "${YELLOW}需要先登录才能查看报名列表${NC}"
    login
    if [ -z "$token" ]; then
      return 1
    fi
  fi
  
  # 获取报名列表
  response=$(curl -s -X GET "${BASE_URL}/api/activity/v1/register/list?activityId=$activity_id&page_num=1&page_size=10" \
    -H "Authorization: Bearer $token")
  
  # 检查响应状态
  check_response "$response" "获取活动报名列表"
  
  # 显示响应内容
  show_response "$response"
  
  # 返回是否成功
  if [ "$(echo "$response" | jq -r '.code')" = "0" ]; then
    return 0
  else
    return 1
  fi
}

# 审核通过报名
approve_registration() {
  registration_id=$1
  echo -e "\n${BLUE}===== 审核通过报名 (报名ID: $registration_id) =====${NC}"
  
  if [ -z "$registration_id" ]; then
    echo -e "${RED}✗ 报名ID为空，无法审核${NC}"
    return 1
  fi
  
  # 检查是否已登录
  if [ -z "$token" ]; then
    echo -e "${YELLOW}需要先登录才能审核报名${NC}"
    login
    if [ -z "$token" ]; then
      return 1
    fi
  fi
  
  # 审核通过
  request_body='{
    "registrationId": '"$registration_id"'
  }'
  
  response=$(curl -s -X POST "${BASE_URL}/api/activity/v1/register/approve" \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: application/json" \
    -d "$request_body")
  
  # 检查响应状态
  check_response "$response" "审核通过报名"
  
  # 显示响应内容
  show_response "$response"
  
  # 返回是否成功
  if [ "$(echo "$response" | jq -r '.code')" = "0" ]; then
    return 0
  else
    return 1
  fi
}

# 拒绝报名
reject_registration() {
  registration_id=$1
  echo -e "\n${BLUE}===== 拒绝报名 (报名ID: $registration_id) =====${NC}"
  
  if [ -z "$registration_id" ]; then
    echo -e "${RED}✗ 报名ID为空，无法拒绝${NC}"
    return 1
  fi
  
  # 检查是否已登录
  if [ -z "$token" ]; then
    echo -e "${YELLOW}需要先登录才能拒绝报名${NC}"
    login
    if [ -z "$token" ]; then
      return 1
    fi
  fi
  
  # 拒绝报名
  request_body='{
    "registrationId": '"$registration_id"',
    "reason": "测试拒绝原因"
  }'
  
  response=$(curl -s -X POST "${BASE_URL}/api/activity/v1/register/reject" \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: application/json" \
    -d "$request_body")
  
  # 检查响应状态
  check_response "$response" "拒绝报名"
  
  # 显示响应内容
  show_response "$response"
  
  # 返回是否成功
  if [ "$(echo "$response" | jq -r '.code')" = "0" ]; then
    return 0
  else
    return 1
  fi
}

# 删除活动
delete_activity() {
  activity_id=$1
  echo -e "\n${BLUE}===== 删除活动 (ID: $activity_id) =====${NC}"
  
  if [ -z "$activity_id" ]; then
    echo -e "${RED}✗ 活动ID为空，无法删除${NC}"
    return 1
  fi
  
  # 检查是否已登录
  if [ -z "$token" ]; then
    echo -e "${YELLOW}需要先登录才能删除活动${NC}"
    login
    if [ -z "$token" ]; then
      return 1
    fi
  fi
  
  # 删除活动
  response=$(curl -s -X DELETE "${BASE_URL}/api/activity/v1/delete/$activity_id" \
    -H "Authorization: Bearer $token")
  
  # 检查响应状态
  check_response "$response" "删除活动"
  
  # 显示响应内容
  show_response "$response"
  
  # 返回是否成功
  if [ "$(echo "$response" | jq -r '.code')" = "0" ]; then
    return 0
  else
    return 1
  fi
}

# 测试主函数
run_tests() {
  # 检查服务器是否运行
  check_server
  
  echo -e "\n${BLUE}===== 开始测试活动API =====${NC}"
  
  # 获取Token
  login
  
  # 先测试获取活动列表(无需认证)
  get_activity_list
  
  # 测试按状态获取活动
  get_activity_list_by_status 0 "未开始的"
  get_activity_list_by_status 1 "进行中的"
  get_activity_list_by_status 2 "已结束的"
  
  # 如果提供了活动ID参数，则测试该活动的详情和报名
  if [ -n "$1" ]; then
    activity_id=$1
    
    # 获取活动详情
    if get_activity_detail $activity_id; then
      # 报名活动
      registration_id=$(register_activity $activity_id)
      
      # 获取报名列表
      if [ -n "$registration_id" ]; then
        get_registration_list $activity_id
        
        # 审核报名
        approve_registration $registration_id
        
        # 再次获取报名列表查看变化
        get_registration_list $activity_id
      fi
    fi
  else
    # 如果没有提供活动ID，则创建新活动进行测试
    echo -e "\n${BLUE}===== 创建新活动并进行完整测试 =====${NC}"
    
    # 创建新活动
    activity_id=$(create_activity)
    
    if [ -n "$activity_id" ]; then
      # 获取新创建的活动详情
      get_activity_detail $activity_id
      
      # 更新活动
      if update_activity $activity_id; then
        # 再次获取活动详情查看更新效果
        get_activity_detail $activity_id
      fi
      
      # 报名活动
      registration_id=$(register_activity $activity_id)
      
      # 获取报名列表
      if [ -n "$registration_id" ]; then
        get_registration_list $activity_id
        
        # 审核通过报名
        if approve_registration $registration_id; then
          # 再次获取报名列表查看变化
          get_registration_list $activity_id
        fi
        
        # 创建另一个报名，用于测试拒绝功能
        second_registration_id=$(register_activity $activity_id)
        if [ -n "$second_registration_id" ]; then
          # 拒绝第二个报名
          if reject_registration $second_registration_id; then
            # 再次获取报名列表查看变化
            get_registration_list $activity_id
          fi
        fi
      fi
      
      # 询问是否要删除测试活动
      echo -e "\n${YELLOW}是否删除测试活动? (y/n)${NC}"
      read -r answer
      if [ "$answer" = "y" ] || [ "$answer" = "Y" ]; then
        delete_activity $activity_id
      else
        echo -e "${GREEN}测试活动已保留，ID: $activity_id${NC}"
      fi
    fi
  fi
  
  echo -e "\n${GREEN}测试完成!${NC}"
}

# 开始执行测试
run_tests $1 