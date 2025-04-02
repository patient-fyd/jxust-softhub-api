# systemctl

## 根据实际情况修改 deploy.sh 部署相关变量

```sh
project_name="jxust-softhub-api"
app_name="${project_name}-api"
systemd_service="${app_name}.service"

# 不要使用个人账号部署和运行服务, 建议建立独立的公共账号部署和运行应用.
deploy_user="vagrant"
# NOTE: 1. 请提前配置发布机到目标服务器之间的ssh key信任:
#    1) ssh-keygen -t rsa -b 4096 -C "vagrant@$(hostname)"
#    2) ssh-copy-id -i ~/.ssh/id_rsa.pub vagrant@$(deploy_server)
# 2. 目标机器上解决sudo执行命令提示输入密码的问题:
#    执行visudo, 末尾添加一行: vagrant ALL=(ALL:ALL) NOPASSWD:ALL
deploy_server="jxust-softhub-api.sre.im"
deploy_dir=/app/$project_name
```

## 在发布机执行部署脚本 deploy.sh

```sh
# 部署测试环境
$ manifest/deploy/systemctl/deploy.sh test

# 部署生产环境
$ manifest/deploy/systemctl/deploy.sh prod
```

## service 文件配置

```ini
[Unit]
Description=jxust-softhub-api-api production environment
Documentation=
After=network.target
Wants=

[Service]
Type=simple
# 运行应用的账号.
User=vagrant
# 程序启动用到的环境变量.
Environment=GF_GERROR_BRIEF=true GF_GCFG_FILE=config.prod.yaml
# 程序所在目录.
WorkingDirectory=/app/jxust-softhub-api
# 使用 sh -c "" 这种方式启动, 避免将stdout和stderr输出到/var/log/message中,
# 统一输出到指定的文件中.
ExecStart=/bin/sh -c '/app/jxust-softhub-api/jxust-softhub-api-api >> jxust-softhub-api-api.log 2>&1'
# 发送SIGTERM信号给所有控制的进程.
KillMode=ctrl-group
# 只能通过 systemctl stop 来关闭服务, 直接kill进程服务还会自动重启.
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

## systemctl 常用操作命令

```sh
# jxust-softhub-api-api.service 配置有变动的时候, 需要重新加载使生效
$ sudo systemctl daemon-reload

# 启动
$ sudo systemctl start jxust-softhub-api-api

# 关闭: 发送 SIGTERM 信号给主(sh)和子进程(jxust-softhub-api-api),
# jxust-softhub-api-api程序可通过捕捉SIGTERM信号来实现优雅关闭.
$ sudo systemctl stop jxust-softhub-api-api

# 重启: 先关闭(SIGTERM), 再启动
$ sudo systemctl restart jxust-softhub-api-api
```
