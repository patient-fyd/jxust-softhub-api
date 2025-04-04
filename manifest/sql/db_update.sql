-- 更新数据库结构
USE jxust_softhub;

-- =====================================================
-- 14. 成员表（扩展用户信息，专门用于协会成员管理）
DROP TABLE IF EXISTS members;
CREATE TABLE members (
    memberId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '成员ID，主键，自增',
    userId INT UNSIGNED NOT NULL COMMENT '关联的用户ID',
    grade VARCHAR(20) NOT NULL COMMENT '年级，如2020级',
    joinYear YEAR NOT NULL COMMENT '入会年份',
    department VARCHAR(50) DEFAULT NULL COMMENT '所属部门，如技术部、宣传部等',
    position VARCHAR(50) DEFAULT NULL COMMENT '职位，如部长、副部长、干事等',
    skills TEXT DEFAULT NULL COMMENT '技能描述',
    introduction TEXT DEFAULT NULL COMMENT '个人简介',
    isCore TINYINT(1) DEFAULT 0 COMMENT '是否为核心成员：0-否，1-是',
    status TINYINT(1) DEFAULT 1 COMMENT '成员状态：0-待审核，1-正式成员，2-已退出',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
    CONSTRAINT fk_member_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储协会成员详细信息';

-- =====================================================
-- 15. 入会申请表
DROP TABLE IF EXISTS join_applications;
CREATE TABLE join_applications (
    applicationId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '申请ID，主键，自增',
    name VARCHAR(50) NOT NULL COMMENT '申请人姓名',
    studentId VARCHAR(20) NOT NULL COMMENT '学号',
    grade VARCHAR(20) NOT NULL COMMENT '年级，如2020级',
    college VARCHAR(100) NOT NULL COMMENT '学院',
    major VARCHAR(100) NOT NULL COMMENT '专业',
    phone VARCHAR(20) NOT NULL COMMENT '联系电话',
    email VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
    reason TEXT NOT NULL COMMENT '申请理由',
    skills TEXT DEFAULT NULL COMMENT '技能介绍',
    expectDepartment VARCHAR(50) DEFAULT NULL COMMENT '期望加入的部门',
    status TINYINT(1) DEFAULT 0 COMMENT '申请状态：0-待审核，1-通过，2-拒绝',
    reviewerId INT UNSIGNED DEFAULT NULL COMMENT '审核人ID，关联users表',
    reviewComment TEXT DEFAULT NULL COMMENT '审核意见',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    CONSTRAINT fk_application_reviewer FOREIGN KEY (reviewerId) REFERENCES users(userId) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储入会申请记录';

-- =====================================================
-- 16. 网站访问统计表
DROP TABLE IF EXISTS visit_statistics;
CREATE TABLE visit_statistics (
    statId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '统计ID，主键，自增',
    visitDate DATE NOT NULL COMMENT '访问日期',
    pageView INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '页面浏览量',
    uniqueVisitor INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '独立访客数',
    newVisitor INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '新访客数',
    avgStayTime INT UNSIGNED DEFAULT 0 COMMENT '平均停留时间（秒）',
    bounceRate DECIMAL(5,2) DEFAULT 0 COMMENT '跳出率（百分比）',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
    UNIQUE KEY idx_visit_date (visitDate)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储网站访问统计数据';

-- =====================================================
-- 17. 文件表（用于统一管理上传的文件）
DROP TABLE IF EXISTS files;
CREATE TABLE files (
    fileId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '文件ID，主键，自增',
    fileName VARCHAR(255) NOT NULL COMMENT '文件名称',
    fileSize BIGINT UNSIGNED NOT NULL COMMENT '文件大小（字节）',
    fileType VARCHAR(100) NOT NULL COMMENT '文件类型/MIME类型',
    filePath VARCHAR(255) NOT NULL COMMENT '文件存储路径',
    uploaderId INT UNSIGNED DEFAULT NULL COMMENT '上传者ID，关联users表',
    uploadTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    md5Hash VARCHAR(32) DEFAULT NULL COMMENT '文件MD5哈希值，用于去重',
    relatedType VARCHAR(50) DEFAULT NULL COMMENT '关联类型，如news、activity、resource等',
    relatedId INT UNSIGNED DEFAULT NULL COMMENT '关联ID，指向具体的内容ID',
    CONSTRAINT fk_file_uploader FOREIGN KEY (uploaderId) REFERENCES users(userId) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储上传的文件信息';

-- =====================================================
-- 18. 标签表
DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
    tagId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '标签ID，主键，自增',
    tagName VARCHAR(50) NOT NULL COMMENT '标签名称',
    tagType VARCHAR(50) DEFAULT NULL COMMENT '标签类型，如news、project、resource等',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE KEY idx_tag_name_type (tagName, tagType)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储内容标签';

-- =====================================================
-- 19. 内容标签关联表
DROP TABLE IF EXISTS content_tags;
CREATE TABLE content_tags (
    contentType VARCHAR(50) NOT NULL COMMENT '内容类型，如news、project、resource等',
    contentId INT UNSIGNED NOT NULL COMMENT '内容ID',
    tagId INT UNSIGNED NOT NULL COMMENT '标签ID，关联tags表',
    PRIMARY KEY (contentType, contentId, tagId),
    CONSTRAINT fk_content_tag FOREIGN KEY (tagId) REFERENCES tags(tagId) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关联内容与标签';

-- =====================================================
-- 20. 系统配置表
DROP TABLE IF EXISTS system_configs;
CREATE TABLE system_configs (
    configKey VARCHAR(100) NOT NULL PRIMARY KEY COMMENT '配置键名',
    configValue TEXT NOT NULL COMMENT '配置值',
    description VARCHAR(255) DEFAULT NULL COMMENT '配置说明',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储系统配置参数';

-- =====================================================
-- 21. 修改news表，添加newsType字段
ALTER TABLE news ADD COLUMN newsType TINYINT(1) DEFAULT 1 COMMENT '新闻类型：1-协会通知，2-技术分享' AFTER category;

-- =====================================================
-- 22. 添加新闻发布相关权限
INSERT INTO permissions (permissionId, permissionKey, description) VALUES 
(5, 'publish_tech_article', '发布技术文章的权限'),
(6, 'publish_association_notice', '发布协会通知的权限');

-- 默认将publish_tech_article权限分配给协会成员角色
-- 将publish_association_notice权限分配给内容管理员角色
INSERT INTO role_permissions (roleId, permissionId) VALUES
(2, 6), -- 内容管理员(roleId=2)拥有发布协会通知的权限(permissionId=6)
(3, 5); -- 协会成员(roleId=3)拥有发布技术分享的权限(permissionId=5)

-- 为用户表添加粉丝数量和关注数量字段
ALTER TABLE users
ADD COLUMN followerCount INT UNSIGNED DEFAULT 0 COMMENT '粉丝数量',
ADD COLUMN followingCount INT UNSIGNED DEFAULT 0 COMMENT '关注数量';

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

-- 为评论表添加点赞数字段
ALTER TABLE comments ADD COLUMN likeCount INT UNSIGNED DEFAULT 0 COMMENT '点赞数';

-- 评论表
CREATE TABLE IF NOT EXISTS `comments` (
  `commentId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `contentType` varchar(20) NOT NULL COMMENT '内容类型(post:帖子)',
  `contentId` int unsigned NOT NULL COMMENT '内容ID',
  `userId` int unsigned NOT NULL COMMENT '用户ID',
  `content` text NOT NULL COMMENT '评论内容',
  `createTime` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`commentId`),
  KEY `idx_contentType_contentId` (`contentType`,`contentId`),
  KEY `idx_userId` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

-- 点赞表
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

-- 关注表
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