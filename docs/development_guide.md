# jxust-softhub-api 开发指南

## 📖 目录

- [简介](#简介)
- [技术栈](#技术栈)
- [工程目录结构](#工程目录结构)
- [环境管理](#环境管理)
- [多命令管理](#多命令管理)
- [错误码管理](#错误码管理)
- [日志管理](#日志管理)
- [链路跟踪](#链路跟踪)
- [版本管理](#版本管理)
- [Redis工具](#redis工具)
- [开发流程](#开发流程)
- [代码质量](#代码质量)
- [项目部署](#项目部署)
- [常用命令](#常用命令)
- [项目改名](#项目改名)
- [参考资源](#参考资源)

## 简介

`jxust-softhub-api` 是江西应用科技学院软件协会官网和后台管理系统的后端项目，基于 [GoFrameV2](https://github.com/gogf/gf) 框架开发，具有以下特性：

- 优化工程目录结构，支持多个可执行命令
- 多环境管理：开发环境、测试环境、生产环境
- 编译的二进制文件可打印当前应用的版本信息
- 中间件统一拦截响应，规范响应格式，规范业务错误码
- 完善的日志系统：HTTP服务访问日志、HTTP服务错误日志、SQL日志等
- 封装Redis常用工具库
- 自动生成数据库层、服务接口层、控制器层代码
- 完整的增删改查接口示例和完善的开发流程

## 技术栈

- **开发语言**：Go v1.18+
- **框架**：GoFrame v2.x
- **数据库**：MySQL 5.7+
- **缓存**：Redis
- **鉴权**：JWT + RBAC 权限模型
- **ORM**：GoFrame ORM
- **API风格**：RESTful
- **日志**：GoFrame 自带 g.Log + 业务审计日志

## 工程目录结构

```sh
├── CHANGELOG.md  # 版本变更管理
├── Dockerfile  # 用于制作容器镜像
├── Makefile  # 用于项目管理
├── README.md  # 项目文档
├── api  # 对外接口定义: 对外提供服务的输入/输出数据结构定义, 路由path定义, 数据校验等
│   ├── api.go  # 接口模块通用结构
│   └── users  # 用户模块
│       ├── users.go  # 用户模块的api interface, 由make ctrl自动生成
│       └── v1  # 版本控制
│           └── users.go  # 开发者按照规范编写的接口文件, make ctrl会根据本文件自动生成controller代码
├── bin  # make build 和 make build.cli 生成的二进制可执行文件所在目录, 不要提交到仓库
├── cmd  # 项目的可执行文件入口
│   ├── jxust-softhub-api-api  # API服务
│   └── jxust-softhub-api-cli  # 项目的其他可执行服务
├── hack  # 存放项目开发工具、脚本等内容
│   └── config.yaml  # gf 工具的配置文件
│   └── change_project_name.sh  # 将示例项目名称改成你自己的项目名称
├── internal
│   ├── cmd  # 对应外层 cmd 目录
│   ├── codes  # 业务错误码定义维护
│   ├── consts  # 项目所有通用常量定义
│   ├── controller  # 控制器层: 接收/解析用户输入参数的入口
│   ├── dao  # 数据访问对象, 由make dao自动生成
│   ├── logic  # 业务封装: 业务逻辑封装管理
│   ├── model  # 数据结构管理模块
│   └── service  # 业务接口层: 用于业务模块解耦的接口定义层
├── manifest  # 交付清单: 包含应用配置文件, 部署文件等
│   ├── config  # 配置文件存放目录
│   ├── deploy  # 和部署相关的文件
│   └── sql  # SQL脚本文件
├── resource  # 静态资源文件
└── utility  # 通用工具类
```

## 环境管理

### 开发环境

配置文件: `manifest/config/config.yaml`

运行:
```sh
make run
# 或
./jxust-softhub-api-api
```

### 测试环境

配置文件: `manifest/config/config.test.yaml`

运行:
```sh
# 通过环境变量指定配置文件
GF_GCFG_FILE=config.test.yaml GF_GERROR_BRIEF=true ./jxust-softhub-api-api

# 或通过命令行参数指定配置文件
./jxust-softhub-api-api -c config.test.yaml
```

### 生产环境

配置文件: `manifest/config/config.prod.yaml`

运行方式同测试环境，只需更换配置文件即可。

## 多命令管理

### 目录设计

- 命令1: `cmd/jxust-softhub-api-api/jxust-softhub-api-api.go` -> `internal/cmd/apiserver/apiserver.go`
- 命令2: `cmd/jxust-softhub-api-cli/jxust-softhub-api-cli.go` -> `internal/cmd/cli/cli.go`

### 配置文件管理

默认不同命令在相同环境下使用同一个配置文件，也可以使用独立配置文件：

```sh
./jxust-softhub-api-cli -c cli_config.yaml
# 或
GF_GCFG_FILE=cli_config.yaml ./jxust-softhub-api-cli
```

## 错误码管理

### 规范制定

- 统一响应格式
```json
{
  "code": "string",
  "message": "string",
  "traceid": "string",
  "data": null
}
```

- 业务码：统一使用字符串表示，如: `"code": "ValidationFailed"`

- HTTP状态码
  - 200: 成功的响应
  - 202: 部分成功的响应
  - 401: 未通过访问认证
  - 403: 请求的资源未获得授权
  - 404: 请求的资源不存在
  - 400: 其他所有客户端错误
  - 500: 所有服务端错误

### 业务错误码

在 `internal/codes/biz_codes.go` 文件中维护业务错误码：

```go
package codes

var (
    CodeOK          = New(200, "OK", "")
    CodePartSuccess = New(202, "PartSuccess", "part success")
    
    CodePermissionDenied = New(401, "AuthFailed", "authentication failed")
    CodeNotAuthorized    = New(403, "NotAuthorized", "resource is not authorized")
    CodeNotFound         = New(404, "NotFound", "resource does not exist")
    CodeValidationFailed = New(400, "ValidationFailed", "validation failed")
    CodeNotAvailable     = New(400, "NotAvailable", "not available")
    
    CodeInternal = New(500, "InternalError", "an error occurred internally")
)
```

## 日志管理

### HTTP服务日志

#### 配置方法

```yaml
# manifest/config/config.yaml
server:
  # 服务日志(包括访问日志和server错误日志)
  logPath: "logs/" # 日志文件存储目录路径
  logStdout: true # 日志是否输出到终端
  errorStack: true # 当Server捕获到异常时是否记录堆栈信息到日志中
  errorLogEnabled: true # 是否记录异常日志信息到日志中
  errorLogPattern: "error-{Ymd}.log" # 异常错误日志文件格式
  accessLogEnabled: true # 是否记录访问日志(包含异常的访问日志)
  accessLogPattern: "access-{Ymd}.log" # 访问日志文件格式
```

### SQL日志

#### 配置方法

```yaml
# manifest/config/config.yaml
database:
  # sql日志
  logger:
    path: "logs/"
    file: "sql-{Ymd}.log"
    level: "all"
    stdout: true
    ctxKeys: ["user", "mail"]
```

### 开发者打印的通用日志

#### 配置方法

```yaml
# manifest/config/config.yaml
logger:
  path: "logs/" # 日志文件目录
  file: "{Ymd}.log" # 日志文件格式
  level: "all" # DEBU < INFO < NOTI < WARN < ERRO < CRIT
  stStatus: 0 # 是否打印错误堆栈(1: enabled; 0: disabled)
  ctxKeys: ["user", "mail"] # 自动打印Context的变量到日志中
  stdout: true # 日志是否同时输出到终端
```

#### 如何打日志

```go
// jxust-softhub-api-api的日志
g.Log().Info(ctx, "hello world")
g.Log().Errorf(ctx, "hello %s", "world")

// jxust-softhub-api-cli的日志
g.Log("cli").Debug(ctx, "hello world")
g.Log("cli").Warningf(ctx, "hello %s", "world")
```

## 链路跟踪

- 用于链路跟踪的响应Header为: `Trace-Id`，会优先使用客户端传递的请求Header `Trace-Id`的值
- 服务内部调用其他服务的接口，请使用`g.Client()`，会自动注入`Trace-Id`

## 版本管理

### 版本变更流程

1. 更新版本变更文档
```text
## v0.3.0

### Added
- xxx
- xxx

### Changed
- xxx
- xxx

### Fixed
- xxx
- xxx
```

2. 给项目仓库打tag
```sh
git tag v0.3.0
git push --tags
```

3. 使用Makefile编译
```sh
# For jxust-softhub-api-api
make build

# For jxust-softhub-api-cli
make build.cli
```

4. 查看二进制文件版本信息
```sh
bin/darwin_amd64/jxust-softhub-api-api -v
```

## Redis工具

### Redis配置

```yaml
# manifest/config/config.yaml
redis:
  # 默认分组, 调用方式: g.Redis()
  default:
    address: 127.0.0.1:6379
    db: 0
    pass:
    maxConnLifetime: 30m
    idleTimeout: 10m
```

### Redis工具库

- Redis缓存: `internal/pkg/rediscache`
- Redis分布式锁: `internal/pkg/redislock`
- Redis消息队列: `internal/pkg/redismq`
- Redis延迟队列: `internal/pkg/redisdelaymq`
- Redis发布订阅: `internal/pkg/redispubsub`

## 开发流程

### 1. 设计表结构，创建物理表

```sql
-- manifest/sql/users.sql
CREATE DATABASE IF NOT EXISTS `your_database`;
USE `your_database`;

-- Create your table
DROP TABLE IF EXISTS `your_table`;
CREATE TABLE `your_table` (
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `fielda`    varchar(45) NOT NULL COMMENT 'Field A',
    `fieldb`    varchar(45) NOT NULL COMMENT 'Field B',
    `created_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `updated_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_fielda` (`fielda`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

创建物理表：
```sh
mysql -uroot -p'password' < manifest/sql/users.sql
```

### 2. 自动生成数据层相关代码

配置gf工具：
```yaml
# hack/config.yaml
gfcli:
  gen:
    dao:
      - link: "mysql:root:password@tcp(127.0.0.1:3306)/your_database"
        tables: "" # 为空表示生成所有表的代码
        descriptionTag: true 
        noModelComment: true
        jsonCase: "snake"
        clear: true
```

生成代码：
```sh
make dao
```

### 3. 编写API层代码

在`api/users/v1/users.go`中定义接口：

```go
// api/users/v1/users.go
type CreateReq struct {
	g.Meta `path:"/users" method:"post" tags:"UserService" summary:"Create a user record"`
	// ... 字段定义
}
```

### 4. 自动生成controller层框架代码

```sh
make ctrl
```

### 5. 编写model层代码

在`internal/model/`目录下创建model文件：

```go
// internal/model/users.go
type UserCreateInput struct {
	// ... 字段定义
}

type UserCreateOutput struct {
	ID uint `json:"id"`
}
```

### 6. 编写service层代码

编写业务逻辑实现：
```go
// internal/logic/users/users.go
func (s *sUsers) Create(ctx context.Context, in model.UserCreateInput) (*model.UserCreateOutput, error) {
	// ... 业务逻辑
	id, err := dao.Users.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.UserCreateOutput{
		ID: uint(id),
	}, nil
}
```

生成service接口代码：
```sh
make service
```

注册业务实现到服务接口：
```go
// internal/logic/users/users.go
func init() {
	service.RegisterUsers(New())
}
```

### 7. 编写controller层代码

```go
// internal/controller/users/users_v1_create.go
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	data := model.UserCreateInput{
		// ... 字段映射
	}
	result, err := service.Users().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{ID: result.ID}, nil
}
```

### 8. 路由注册

```go
// internal/cmd/apiserver/apiserver.go
var (
    Main = gcmd.Command{
        Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
            s := g.Server()

            s.Group("/v1", func(group *ghttp.RouterGroup) {
                group.Bind(
                    // 对外暴露的接口
                    users.NewV1(),
                )
            })

            s.Run()
            return nil
        },
    }
)
```

### 9. 接口访问测试

```sh
# 运行
make run

# 测试接口
curl -X POST -i 'localhost:9000/v1/users' -d '{"username": "test", "password": "test12345678"}'
```

## 代码质量

### 代码检查工具安装

```sh
# 自动安装并运行golangci-lint
make lint

# 或手动安装
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 使用方法

```sh
# 检查所有代码
golangci-lint run
# 或
make lint

# 查看所有linters
golangci-lint help linters

# 查看已启用的linters
golangci-lint linters

# 使用特定linter
golangci-lint run --no-config --disable-all -E errcheck
```

## 项目部署

### Systemctl部署

1. 设置目标服务器
2. 执行部署脚本：
```sh
# 部署测试环境
./manifest/deploy/systemctl/deploy.sh test

# 部署生产环境
./manifest/deploy/systemctl/deploy.sh prod
```

3. systemctl常用命令：
```sh
sudo systemctl daemon-reload
sudo systemctl start jxust-softhub-api-api
sudo systemctl stop jxust-softhub-api-api
sudo systemctl restart jxust-softhub-api-api
```

### Supervisor部署

1. 在目标服务器上安装supervisor
2. 执行部署脚本：
```sh
# 部署测试环境
./manifest/deploy/supervisor/deploy.sh test

# 部署生产环境
./manifest/deploy/supervisor/deploy.sh prod
```

3. supervisorctl常用命令：
```sh
sudo supervisorctl start jxust-softhub-api-api
sudo supervisorctl stop jxust-softhub-api-api
sudo supervisorctl restart jxust-softhub-api-api
sudo supervisorctl reload
```

### Docker部署

1. 制作镜像：
```sh
make image
```

2. 运行容器：
```sh
# 开发环境
docker run --name jxust-softhub-api -p80:9000 -d jxust-softhub-api-api:tag

# 测试环境
docker run --name jxust-softhub-api -p80:9000 -e GF_GCFG_FILE=config.test.yaml -d jxust-softhub-api-api:tag

# 生产环境
docker run --name jxust-softhub-api -p80:9000 -e GF_GCFG_FILE=config.prod.yaml -d jxust-softhub-api-api:tag
```

## 常用命令

```sh
# 安装最新版gf
make cli

# 运行golangci-lint检查代码
make lint

# 生成数据层代码
make dao

# 生成控制器层代码
make ctrl

# 生成服务接口代码
make service

# 开发环境热启动API服务
make run

# 开发环境热启动CLI服务
make run.cli

# 编译API服务
make build

# 编译CLI服务
make build.cli

# 制作docker镜像
make image

# 制作并推送docker镜像
make image.push
```

## 项目改名

如需将项目名称从`jxust-softhub-api`改为你自己的项目名称：

1. 变更项目目录名称
```sh
mv jxust-softhub-api your-project-name
```

2. 运行变更脚本
```sh
cd your-project-name
hack/change_project_name.sh your-project-name
```

3. 验证
```sh
make build
```

## 参考资源

- [GoFrame官方文档](https://goframe.org)
- [GoFrame ORM文档](https://goframe.org/pages/viewpage.action?pageId=1114245)
- [JWT认证文档](https://goframe.org/pages/viewpage.action?pageId=1114217)
- [错误处理文档](https://goframe.org/pages/viewpage.action?pageId=1114399)
