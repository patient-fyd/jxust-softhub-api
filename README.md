# jxust-softhub-api 开发指南

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://go.dev)
[![Framework](https://img.shields.io/badge/Framework-GoFrame2-green.svg)](https://goframe.org)
[![Database](https://img.shields.io/badge/Database-MySQL-orange.svg)](https://www.mysql.com)

江西应用科技学院软件协会官网 & 后台管理系统后端项目

## 📌 项目介绍

`jxust-softhub-api` 是本系统后端接口服务，基于 GoFrame2 + MySQL 开发，负责为 **官网前端（jxust-softhub-web）** 和 **后台管理系统（jxust-softhub-admin）** 提供 RESTful API，涵盖用户、权限、新闻、活动、项目、资源、审核、统计、互动等核心模块。

同时支持与已有的 **C++ OJ 系统** 和 **搜索引擎系统** 集成，保证系统数据统一、接口高可用、安全可靠。

## 🛠 技术栈

- **开发语言**：Go v1.18+
- **框架**：GoFrame v2.x
- **数据库**：MySQL
- **缓存**：Redis
- **鉴权**：JWT + RBAC 权限模型
- **ORM**：GoFrame ORM
- **API风格**：RESTful
- **日志**：GoFrame 自带 g.Log + 业务审计日志

## 🚀 快速开始

### 1. 环境准备

- Go 1.18+ 
- MySQL 5.7+
- Redis (可选)

### 2. 获取代码

```bash
git clone git@github.com:patient-fyd/jxust-softhub-api.git
cd jxust-softhub-api
```

### 3. 安装依赖

```bash
# 安装gf工具
make cli

# 安装项目依赖
go mod tidy
```

### 4. 配置数据库

```bash
# 导入数据库
mysql -uroot -p'你的密码' < manifest/sql/db.sql
```

修改数据库配置:

```bash
# 修改开发环境配置
vi manifest/config/config.yaml

# 测试/生产环境
vi manifest/config/config.test.yaml  
vi manifest/config/config.prod.yaml
```

### 5. 开发与调试

```bash
# 热更新启动
make run

# 访问swagger文档
浏览器访问: http://localhost:9000/swagger
```

## 📁 工程目录说明

```
├── api         # 接口定义，对外提供服务的输入/输出数据结构，路由定义，数据校验
├── cmd         # 项目启动入口
├── hack        # 项目工具脚本
├── internal    # 内部代码
│   ├── cmd        # 命令行
│   ├── codes      # 业务错误码
│   ├── controller # 控制器层
│   ├── dao        # 数据访问对象
│   ├── logic      # 业务逻辑层
│   ├── model      # 数据模型
│   ├── service    # 服务接口层
├── manifest    # 配置文件目录
├── utility     # 工具类
```

## ✅ 功能模块与开发流程

### 一、用户与权限系统

- 用户注册/登录 (JWT)
- 用户信息管理
- 用户角色管理
- 权限控制 (RBAC)
- 前后台用户分离登录

#### 开发步骤
1. 定义 API 接口 (`api/auth/v1/auth.go` 和 `api/user/v1/user.go`)
2. 生成控制器代码 (`make ctrl`)
3. 实现业务逻辑 (`internal/logic/auth` 和 `internal/logic/user`)
4. 生成服务接口 (`make service`)
5. 在路由中注册接口 (`internal/cmd/apiserver/apiserver.go`)

### 二、新闻动态

- 新闻分类管理
- 新闻内容管理
- 新闻详情接口
- 新闻搜索

#### 开发步骤
1. 定义 API 接口 (`api/news/v1/news.go`)
2. 生成控制器代码 (`make ctrl`)
3. 实现业务逻辑 (`internal/logic/news`)
4. 生成服务接口 (`make service`)
5. 在路由中注册接口

### 三、活动管理

- 活动信息管理
- 活动详情
- 报名接口
- 报名数据管理
- 报名审核
- 报名状态管理

#### 开发步骤
1. 定义 API 接口 (`api/activity/v1/activity.go`)
2. 生成控制器代码 (`make ctrl`)
3. 实现业务逻辑 (`internal/logic/activity`)
4. 生成服务接口 (`make service`)
5. 在路由中注册接口

### 四、项目管理

- 项目基本信息管理
- OJ系统集成
- 搜索引擎系统集成
- 自主项目展示接口

#### 开发步骤
同上，创建对应 API 和实现逻辑

### 五、成员与加入管理

- 成员信息展示
- 核心成员管理
- 协会成员分届管理
- 入会申请及审核

#### 开发步骤
同上，创建对应 API 和实现逻辑

### 六、资源管理

- 资源分类管理
- 文件上传/下载接口
- 多媒体文件支持

#### 开发步骤
同上，创建对应 API 和实现逻辑

### 七、评论 & 消息系统

- 新闻 & 活动评论
- 用户私信
- 通知中心
- 内容审核机制

#### 开发步骤
同上，创建对应 API 和实现逻辑

### 八、统计分析

- 网站访问统计
- 活动报名数据统计
- 用户活跃度统计
- 数据可视化接口

#### 开发步骤
同上，创建对应 API 和实现逻辑

### 九、日志系统

- 操作日志
- 审计日志
- 关键操作报警

#### 开发步骤
同上，创建对应 API 和实现逻辑

## 📘 开发规范

### API 接口规范

- 所有接口遵循 RESTful 风格
- 统一返回格式：`code` + `message` + `traceid` + `data`
- 所有接口需支持 JWT 鉴权和权限校验
- 管理端接口按角色严格控制权限

### 目录命名规范

- 模块名使用单数，如 `user` 而非 `users`
- API 版本控制：`api/模块名称/v1/模块名称.go`
- 控制器：`internal/controller/模块名称/`
- 业务逻辑：`internal/logic/模块名称/`

### 开发流程总结

1. 设计表结构，创建物理表 (已完成)
2. 自动生成数据层代码 (`make dao`)
3. 编写 API 层代码 (数据结构、接口定义)
4. 生成控制器层代码 (`make ctrl`)
5. 编写 Model 层代码 (内部数据结构)
6. 编写 Service 层代码 (业务实现)
7. 生成服务接口 (`make service`)
8. 路由注册
9. 接口测试

## 🪄 版本管理规范

### 分支管理 (Gitflow)

| 分支      | 说明         |
| --------- | ------------ |
| main      | 正式线上分支 |
| develop   | 日常开发分支 |
| feature/* | 功能开发分支 |
| fix/*     | bug 修复分支 |
| hotfix/*  | 紧急修复分支 |
| release/* | 版本发布分支 |

### Commit 规范

| 类型     | 说明                     |
| -------- | ------------------------ |
| feat     | 新增接口/功能           |
| fix      | 修复 Bug                 |
| refactor | 重构代码，不影响接口     |
| style    | 代码风格/格式化          |
| chore    | 构建、脚本、CI/CD        |
| docs     | 文档更新                 |
| test     | 单元测试                 |

#### Commit 示例

```
feat: 完成用户注册接口
fix: 修复 JWT Token 刷新异常
refactor: 重构报名统计接口
docs: 添加活动报名接口文档
```

## 🔧 常用命令

```bash
# 安装 gf 工具
make cli

# 生成 dao 层代码
make dao

# 生成控制器代码
make ctrl

# 生成服务接口代码
make service

# 热更新启动开发环境
make run

# 编译生成二进制文件
make build

# 运行代码检查
make lint
```

## 📝 注意事项

- 不要直接修改 `internal/dao`、`internal/model/do`、`internal/model/entity` 目录下的代码，这些是自动生成的
- 不要修改 `internal/controller/*/xxx_new.go` 文件，这是自动生成的控制器初始化代码
- 开发新功能时，按照GoFrame2框架规范创建对应的模块结构
- 接口开发时，保持前后端接口风格统一
- 推荐使用 Swagger 或 Apifox 维护接口文档
- 推荐使用 GoFrame2 内置中间件（权限、限流、日志）+ 自定义扩展

## 📚 参考链接

- [GoFrame 官方文档](https://goframe.org)
- [JWT 认证](https://goframe.org/pages/viewpage.action?pageId=1114217)
- [GoFrame ORM](https://goframe.org/pages/viewpage.action?pageId=1114245)
- [错误处理](https://goframe.org/pages/viewpage.action?pageId=1114399)

## 开发步骤

1. 确保开发环境已正确设置，所有依赖都已安装
2. 理解项目架构和目录结构
3. 使用GoFrame2框架的规范和模式开发新功能
4. 开发新功能时，遵循项目中定义的设计模式和代码规范
