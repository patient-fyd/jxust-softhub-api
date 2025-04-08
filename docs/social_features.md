# 社交功能设计与实现

本文档详细介绍软件社团系统中社交功能的设计与实现，包括用户关注关系、圈子机制以及点赞系统。

## 1. 数据模型设计

### 1.1 关注表（follows）

```sql
CREATE TABLE IF NOT EXISTS `follows` (
  `followId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '关注ID',
  `userId` int unsigned NOT NULL COMMENT '用户ID',
  `followedId` int unsigned NOT NULL COMMENT '被关注对象ID',
  `followType` tinyint unsigned NOT NULL COMMENT '关注类型(1:用户,2:圈子)',
  `createTime` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`followId`),
  UNIQUE KEY `uk_userId_followedId_followType` (`userId`,`followedId`,`followType`),
  KEY `idx_followedId_followType` (`followedId`,`followType`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关注表';
```

### 1.2 圈子表（circles）

```sql
CREATE TABLE IF NOT EXISTS `circles` (
  `circleId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '圈子ID',
  `circleName` varchar(50) NOT NULL COMMENT '圈子名称',
  `description` varchar(255) DEFAULT NULL COMMENT '圈子描述',
  `icon` varchar(255) DEFAULT NULL COMMENT '圈子图标',
  `userId` int unsigned NOT NULL COMMENT '创建者ID',
  `memberCount` int unsigned DEFAULT 0 COMMENT '成员数量',
  `postCount` int unsigned DEFAULT 0 COMMENT '帖子数量',
  `status` tinyint unsigned DEFAULT 1 COMMENT '状态(1:正常,0:已删除)',
  `isOfficial` int DEFAULT 0 COMMENT '是否官方圈子：0-否，1-是',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`circleId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='圈子表';
```

### 1.3 点赞表（likes）

```sql
CREATE TABLE IF NOT EXISTS `likes` (
  `likeId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '点赞ID',
  `userId` int unsigned NOT NULL COMMENT '用户ID',
  `targetId` int unsigned NOT NULL COMMENT '目标ID',
  `targetType` tinyint unsigned NOT NULL COMMENT '目标类型(1:帖子,2:评论)',
  `createTime` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`likeId`),
  UNIQUE KEY `uk_userId_targetId_targetType` (`userId`,`targetId`,`targetType`),
  KEY `idx_targetId_targetType` (`targetId`,`targetType`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';
```

### 1.4 用户表（users）扩展字段

```sql
ALTER TABLE users
ADD COLUMN followerCount INT UNSIGNED DEFAULT 0 COMMENT '粉丝数量',
ADD COLUMN followingCount INT UNSIGNED DEFAULT 0 COMMENT '关注数量';
```

## 2. 关注关系功能

### 2.1 关注类型

系统支持两种关注类型：
- 关注用户（followType=1）：用户A关注用户B，用户A成为用户B的粉丝
- 关注圈子（followType=2）：用户加入某个圈子，成为圈子成员

### 2.2 用户关注统计

为提高查询效率，在用户表中添加了两个计数字段：
- followerCount：粉丝数量
- followingCount：关注数量

通过MySQL触发器自动维护这两个字段的计数：
```sql
-- 创建更新用户关注计数的触发器
DELIMITER //
CREATE TRIGGER after_follow_insert
AFTER INSERT ON follows
FOR EACH ROW
BEGIN
    -- 如果是关注用户类型(followType=1)
    IF NEW.followType = 1 THEN
        -- 增加被关注用户的粉丝数
        UPDATE users SET followerCount = followerCount + 1 WHERE userId = NEW.followedId;
        -- 增加关注者的关注数
        UPDATE users SET followingCount = followingCount + 1 WHERE userId = NEW.userId;
    END IF;
END//

CREATE TRIGGER after_follow_delete
AFTER DELETE ON follows
FOR EACH ROW
BEGIN
    -- 如果是关注用户类型(followType=1)
    IF OLD.followType = 1 THEN
        -- 减少被关注用户的粉丝数
        UPDATE users SET followerCount = followerCount - 1 WHERE userId = OLD.followedId AND followerCount > 0;
        -- 减少关注者的关注数
        UPDATE users SET followingCount = followingCount - 1 WHERE userId = OLD.userId AND followingCount > 0;
    END IF;
END//
DELIMITER ;
```

## 3. 圈子功能

### 3.1 圈子属性

- 圈子名称和描述：圈子的基本信息
- 成员数量：自动统计关注该圈子的用户数
- 帖子数量：自动统计发布在该圈子下的帖子数
- 官方标识：区分官方圈子和用户创建的圈子

### 3.2 核心接口

- 圈子列表（List）：支持分页、关键词搜索、排序等
- 圈子详情（Detail）：查看圈子的详细信息
- 关注/取消关注（Join）：用户加入或退出圈子

### 3.3 关注圈子实现

```go
// Join 关注/取消关注圈子
func (s *sCircle) Join(ctx context.Context, in model.CircleJoinInput) (*model.CircleJoinOutput, error) {
    // 获取当前登录用户ID
    loginUserId := auth.GetLoginUserId(ctx)
    if loginUserId <= 0 {
        return nil, gerror.New("请先登录")
    }

    // 检查圈子是否存在
    var circle *entity.Circles
    err := dao.Circles.Ctx(ctx).
        Where("circleId", in.CircleId).
        Where("status", 1). // 只查询正常状态的圈子
        Scan(&circle)
    if err != nil {
        return nil, err
    }
    if circle == nil {
        return nil, gerror.New("圈子不存在或已删除")
    }

    // 判断当前用户是否已关注该圈子
    var isFollowed bool
    count, err := dao.Follows.Ctx(ctx).
        Where("userId", loginUserId).
        Where("followedId", in.CircleId).
        Where("followType", 2). // 2表示关注圈子
        Count()
    if err != nil {
        return nil, err
    }
    isFollowed = count > 0

    // 已关注，则取消关注；未关注，则添加关注
    if isFollowed {
        // 取消关注
        _, err = dao.Follows.Ctx(ctx).
            Where("userId", loginUserId).
            Where("followedId", in.CircleId).
            Where("followType", 2).
            Delete()
        if err != nil {
            return nil, err
        }

        // 更新圈子成员数量（减1）
        _, err = dao.Circles.Ctx(ctx).
            Data("memberCount=memberCount-1").
            Where("circleId", in.CircleId).
            Where("memberCount > 0").
            Update()
        if err != nil {
            return nil, err
        }

        return &model.CircleJoinOutput{
            IsFollowed: false,
        }, nil
    } else {
        // 添加关注
        _, err = dao.Follows.Ctx(ctx).Insert(g.Map{
            "userId":     loginUserId,
            "followedId": in.CircleId,
            "followType": 2, // 2表示关注圈子
            "createTime": gtime.Now(),
        })
        if err != nil {
            return nil, err
        }

        // 更新圈子成员数量（加1）
        _, err = dao.Circles.Ctx(ctx).
            Data("memberCount=memberCount+1").
            Where("circleId", in.CircleId).
            Update()
        if err != nil {
            return nil, err
        }

        return &model.CircleJoinOutput{
            IsFollowed: true,
        }, nil
    }
}
```

## 4. 点赞系统

### 4.1 点赞类型

系统支持两种点赞类型：
- 帖子点赞（targetType=1）
- 评论点赞（targetType=2）

### 4.2 点赞计数设计

为了提高查询效率，在帖子表和评论表中添加了点赞计数字段：
- posts表：likeCount字段
- comments表：likeCount字段

这些字段在添加或取消点赞时同步更新。

### 4.3 点赞/取消点赞的实现

点赞操作采用"toggle"模式，即同一接口实现点赞和取消点赞功能：
- 如果用户之前未点赞，则添加点赞记录
- 如果用户之前已点赞，则删除点赞记录

### 4.4 点赞状态查询

在获取帖子列表、评论列表等API中，会返回当前登录用户是否已点赞的状态信息，前端根据该状态显示不同的点赞UI。

## 5. 性能考虑

### 5.1 索引优化

所有关键查询字段都添加了合适的索引，包括：
- follows表：userId、followedId和followType的联合索引
- likes表：userId、targetId和targetType的联合索引

### 5.2 计数缓存

为避免频繁统计，系统在实体表中直接存储了计数字段：
- users表：followerCount, followingCount
- circles表：memberCount, postCount
- posts表：likeCount, commentCount
- comments表：likeCount

### 5.3 批量查询优化

在需要获取大量关联数据的场景（如判断用户是否点赞过多个帖子），使用了批量查询而非循环单条查询，以减少数据库交互次数。 