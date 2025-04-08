# 帖子创建接口排查流程

## 问题描述

帖子创建功能(`POST /api/post/v1/create`)在测试过程中出现无法成功创建的情况，返回"an error occurred internally"错误，具体表现为返回500状态码，错误信息为"runtime error: invalid memory address or nil pointer dereference"。

## 排查流程

本文档记录排查流程、问题原因及解决方案，用于指导系统开发和维护人员解决类似问题。

## 一、功能测试用例设计（黑盒）

| 测试编号 | 场景描述 | 输入 | 预期结果 | 实际结果 |
|---------|---------|------|---------|---------|
| T01 | 正常创建帖子（含图片+圈子+话题） | 有效 Content、Images、CircleId、TopicId | 成功返回 PostId | 失败：内部错误 |
| T02 | 创建帖子（无图片） | Content + TopicId + CircleId | 成功返回 PostId | 失败：内部错误 |
| T03 | 创建帖子（无圈子） | Content + TopicId | 成功返回 PostId | 未测试 |
| T04 | 创建帖子（无话题） | Content + CircleId | 成功返回 PostId | 未测试 |
| T05 | 创建帖子（图片为空数组） | Content + [] | 成功 | 失败：内部错误 |
| T06 | 创建帖子（非法图片格式） | ImageURL 为 "not-a-url" | 后端验证失败或成功 | 未测试 |
| T07 | 未登录状态提交帖子 | Header 无 token | 报错："用户未登录" | 未测试 |
| T08 | 提交空内容 | Content 为 "" | 拒绝或默认允许 | 未测试 |
| T09 | 提交不存在的 CircleId | CircleId=999999 | 报错或插入后外键失败 | 未测试 |
| T10 | 提交不存在的 TopicId | TopicId=999999 | 同上 | 未测试 |

## 二、后端排查定位策略（白盒）

### 2.1 日志校验

1. 在Controller层添加了调试日志：
   - 接收到帖子创建请求
   - 调用Post服务创建帖子
   - 捕获可能的panic异常
   - 帖子创建成功/失败日志

2. 在Service层添加了调试日志：
   - 帖子创建输入参数
   - 当前登录用户ID
   - 创建帖子记录
   - 事务执行状态

通过日志发现：
- 请求参数正常
- 用户已登录（管理员用户）
- 执行到调用dao.Posts.Ctx(ctx).TX(tx).Insert()方法时出现nil指针异常

### 2.2 数据库验证

检查数据库表结构和数据：

1. 数据库连接正常，可以成功连接到MySQL数据库
2. 表结构验证
   - Posts表结构正确
   - PostImages表结构正确
   - Circle和Topic表结构正确
3. 手动测试
   - 通过SQL直接插入Posts记录成功
   - 检查外键约束，外键设置正确
   - Circle和Topic表中存在ID=1的记录

### 2.3 事务中断识别

在Create方法中关键位置添加日志，发现事务初始化成功，但在调用dao.Posts.Ctx(ctx).TX(tx).Insert()方法时出现异常，事务未能正常执行。

## 三、结构与类型检查

### 3.1 数据结构匹配检查

1. entity与数据库表结构匹配：
   - Posts表字段与entity.Posts结构体字段匹配
   - 字段类型正确
   
2. do层与数据库表结构匹配：
   - do.Posts结构体的字段全部为interface{}类型
   - 字段名称与数据库表字段名称匹配
   
3. 外键关系：
   - circleId和topicId是外键
   - 允许为NULL

### 3.2 默认值设置

1. 默认值检查：
   - viewCount、likeCount、commentCount、shareCount默认值为0
   - status默认值为1（已发布）
   - isTop默认值为0（不置顶）

## 四、问题发现与解决方案

### 4.1 问题原因

经过排查，发现以下几个可能导致问题的原因：

1. 事务处理中使用了do.Posts结构体，而Insert方法期望的是一个map或结构体
2. 可能存在字段值类型不匹配的问题
3. dao层的Posts对象可能未正确初始化

### 4.2 解决方案

1. 简化事务处理，使用g.Map替代do.Posts结构体：

```go
data := g.Map{
    "userId":       userId,
    "content":      in.Content,
    "circleId":     in.CircleId,
    "topicId":      in.TopicId,
    "viewCount":    0,
    "likeCount":    0,
    "commentCount": 0,
    "shareCount":   0,
    "isTop":        0,
    "status":       1, // 状态：1-已发布
    "createTime":   gtime.Now(),
    "updateTime":   gtime.Now(),
}

result, err := dao.Posts.Ctx(ctx).Data(data).Insert()
```

2. 暂时移除事务处理，先确保最基本的插入功能正常：

```go
// 简化实现，不使用事务
// 创建帖子记录
g.Log().Debug(ctx, "创建帖子记录, 用户ID:", userId)

data := g.Map{
    "userId":       userId,
    "content":      in.Content,
    "circleId":     in.CircleId,
    "topicId":      in.TopicId,
    // ...其他字段
}

result, err := dao.Posts.Ctx(ctx).Data(data).Insert()
```

3. 确保类型正确：

```go
// 确保ID类型匹配
userId := uint(auth.GetLoginUserId(ctx))
```

4. 添加更详细的错误处理和日志记录：

```go
// 在每个可能出错的地方添加详细日志
if err != nil {
    g.Log().Error(ctx, "创建帖子失败:", err)
    return nil, err
}
```

## 五、后续优化建议

### 5.1 代码层面优化

1. 统一使用g.Map进行数据库操作，避免使用do结构体可能存在的类型不匹配问题
2. 添加参数验证，确保所有必填字段都有值
3. 对图片URL进行格式校验

### 5.2 稳定性提升

1. 事务中添加defer恢复机制，防止panic导致事务未回滚：

```go
defer func() {
    if r := recover(); r != nil {
        g.Log().Error(ctx, "帖子创建过程中发生panic:", r)
        err = gerror.New("帖子创建失败，系统内部错误")
    }
}()
```

2. 增加单元测试，模拟各种边界条件下的帖子创建场景

3. 统一错误处理机制，返回标准错误码和错误提示

### 5.3 监控与告警

1. 添加系统监控，记录帖子创建成功率
2. 设置告警机制，当帖子创建失败率超过阈值时自动告警

## 六、测试验证

完成修改后，重新执行测试用例，确认功能是否正常：

1. 再次测试帖子创建功能
2. 检查数据库中是否成功插入记录
3. 验证异常情况下的错误处理是否正确

## 七、参考资料

1. [GoFrame 数据库操作文档](https://goframe.org/pages/viewpage.action?pageId=1114245)
2. [GoFrame 事务处理文档](https://goframe.org/pages/viewpage.action?pageId=1114378)
3. [GoFrame 日志处理文档](https://goframe.org/pages/viewpage.action?pageId=1114388) 