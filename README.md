# 江西应用科技学院软件协会官网后端API服务

## 项目简介

本项目是江西应用科技学院软件协会官网和后台管理系统的后端API服务，基于GoFrameV2框架开发。主要提供新闻管理、会员管理、活动管理、文件管理等功能的接口服务。

## 技术栈

- **开发语言**：Go v1.18+
- **框架**：GoFrame v2.x
- **数据库**：MySQL 5.7+
- **缓存**：Redis
- **鉴权**：JWT + RBAC 权限模型
- **ORM**：GoFrame ORM
- **API风格**：RESTful
- **日志系统**：GoFrame 自带 g.Log + 业务审计日志

## 功能模块

本项目包含以下主要功能模块：

1. **认证模块**
   - 用户注册
   - 用户登录
   - 管理员登录
   - Token刷新

2. **用户模块**
   - 获取用户信息
   - 获取用户列表
   - 分配用户角色

3. **会员申请模块**
   - 提交加入申请
   - 获取申请列表
   - 获取申请详情
   - 审核申请

4. **新闻模块**
   - 获取新闻列表
   - 获取新闻详情
   - 创建新闻
   - 更新新闻
   - 删除新闻

## 项目结构

```sh
├── api             # 对外接口定义
├── cmd             # 命令行工具
├── internal        # 内部实现
│   ├── controller  # 控制器层
│   ├── logic       # 业务逻辑层
│   ├── model       # 数据模型层
│   ├── service     # 服务接口层
│   └── dao         # 数据访问层
├── manifest        # 配置文件
│   ├── config      # 配置项
│   ├── deploy      # 部署配置
│   └── sql         # SQL脚本
└── utility         # 工具库
```

## 开发指南

### 环境要求

- Go 1.18+
- MySQL 5.7+
- Redis

### 初始化数据库

```bash
# 执行SQL脚本
mysql -uroot -p'your_password' < manifest/sql/jxust_softhub.sql
```

### 安装依赖

```bash
go mod tidy
```

### 运行项目

```bash
# 开发环境
make run

# 生产环境
GF_GCFG_FILE=config.prod.yaml ./jxust-softhub-api-api
```

### 代码生成

```bash
# 生成数据访问层代码
make dao

# 生成控制器代码
make ctrl

# 生成服务接口代码
make service
```

## API文档

项目内置Swagger文档，运行项目后访问：

```
http://localhost:8000/swagger
```

## 部署指南

### Docker部署

```bash
# 构建镜像
docker build -t jxust-softhub-api .

# 运行容器
docker run -d -p 8000:8000 \
  -e GF_GCFG_FILE=config.prod.yaml \
  --name jxust-softhub-api \
  jxust-softhub-api
```

## 贡献指南

欢迎提交PR或Issue，共同改进项目质量！

## 版权信息

Copyright © 2024 江西应用科技学院软件协会 