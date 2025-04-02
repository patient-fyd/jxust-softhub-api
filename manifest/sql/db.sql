-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS jxust_softhub 
  DEFAULT CHARACTER SET utf8mb4 
  COLLATE utf8mb4_unicode_ci;
USE jxust_softhub;

-- =====================================================
-- 1. 角色表
DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    roleId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '角色ID，主键，自增',
    roleName VARCHAR(50) NOT NULL COMMENT '角色名称，如超级管理员、内容管理员等',
    description VARCHAR(255) DEFAULT NULL COMMENT '角色描述信息',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储角色信息';

-- =====================================================
-- 2. 权限表
DROP TABLE IF EXISTS permissions;
CREATE TABLE permissions (
    permissionId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '权限ID，主键，自增',
    permissionKey VARCHAR(50) NOT NULL COMMENT '权限标识，用于后台校验，如 "news_edit" ',
    description VARCHAR(255) DEFAULT NULL COMMENT '权限描述信息',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储系统各项权限';

-- =====================================================
-- 3. 角色权限关联表
DROP TABLE IF EXISTS role_permissions;
CREATE TABLE role_permissions (
    roleId INT UNSIGNED NOT NULL COMMENT '角色ID，关联 roles 表',
    permissionId INT UNSIGNED NOT NULL COMMENT '权限ID，关联 permissions 表',
    PRIMARY KEY (roleId, permissionId),
    CONSTRAINT fk_role FOREIGN KEY (roleId) REFERENCES roles(roleId) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_permission FOREIGN KEY (permissionId) REFERENCES permissions(permissionId) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关联角色与权限';

-- =====================================================
-- 4. 用户表
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    userId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID，主键，自增',
    userName VARCHAR(50) NOT NULL COMMENT '用户名，登录和显示名称',
    password VARCHAR(100) NOT NULL COMMENT '用户密码，存储加密后的密码',
    name VARCHAR(50) DEFAULT NULL COMMENT '真实姓名',
    roleId INT UNSIGNED DEFAULT NULL COMMENT '角色ID，关联 roles 表，标识用户所属角色',
    avatar VARCHAR(255) DEFAULT NULL COMMENT '用户头像图片URL',
    joinYear YEAR DEFAULT NULL COMMENT '入会年份，格式如2025',
    email VARCHAR(100) DEFAULT NULL COMMENT '用户邮箱',
    phone VARCHAR(20) DEFAULT NULL COMMENT '用户联系电话',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
    CONSTRAINT fk_user_role FOREIGN KEY (roleId) REFERENCES roles(roleId) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储用户基本信息';

-- =====================================================
-- 5. 新闻表
DROP TABLE IF EXISTS news;
CREATE TABLE news (
    newsId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '新闻ID，主键，自增',
    title VARCHAR(255) NOT NULL COMMENT '新闻标题',
    content TEXT NOT NULL COMMENT '新闻内容，支持 Markdown 格式',
    category VARCHAR(50) NOT NULL COMMENT '新闻分类，如协会新闻、技术分享、赛事通知等',
    coverImage VARCHAR(255) DEFAULT NULL COMMENT '封面图片的URL',
    authorId INT UNSIGNED DEFAULT NULL COMMENT '作者ID，关联 users 表',
    viewCount INT DEFAULT 0 COMMENT '浏览次数',
    status TINYINT(1) DEFAULT 0 COMMENT '新闻状态，0: 草稿, 1: 发布',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
    CONSTRAINT fk_news_author FOREIGN KEY (authorId) REFERENCES users(userId) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储新闻资讯内容';

-- =====================================================
-- 6. 活动表
DROP TABLE IF EXISTS activities;
CREATE TABLE activities (
    activityId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '活动ID，主键，自增',
    title VARCHAR(255) NOT NULL COMMENT '活动标题',
    description TEXT NOT NULL COMMENT '活动详细描述',
    startTime DATETIME NOT NULL COMMENT '活动开始时间',
    endTime DATETIME NOT NULL COMMENT '活动结束时间',
    location VARCHAR(255) NOT NULL COMMENT '活动举办地点',
    maxParticipants INT DEFAULT 0 COMMENT '最大参与人数限制',
    status TINYINT(1) DEFAULT 0 COMMENT '活动状态：0-未开始, 1-进行中, 2-已结束',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储活动信息';

-- =====================================================
-- 7. 活动报名表
DROP TABLE IF EXISTS activity_registrations;
CREATE TABLE activity_registrations (
    registrationId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '报名记录ID，主键，自增',
    activityId INT UNSIGNED NOT NULL COMMENT '活动ID，关联 activities 表',
    userId INT UNSIGNED DEFAULT NULL COMMENT '报名用户ID，关联 users 表，如为空表示未登录报名',
    name VARCHAR(50) NOT NULL COMMENT '报名者姓名',
    studentId VARCHAR(20) NOT NULL COMMENT '报名者学号',
    contact VARCHAR(50) NOT NULL COMMENT '报名者联系方式，如电话或邮箱',
    status TINYINT(1) DEFAULT 0 COMMENT '报名状态：0-待审核, 1-通过, 2-拒绝',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
    CONSTRAINT fk_activity FOREIGN KEY (activityId) REFERENCES activities(activityId) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_registration_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储活动报名记录';

-- =====================================================
-- 8. 资源表
DROP TABLE IF EXISTS resources;
CREATE TABLE resources (
    resourceId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '资源ID，主键，自增',
    title VARCHAR(255) NOT NULL COMMENT '资源标题',
    description TEXT DEFAULT NULL COMMENT '资源描述信息',
    category VARCHAR(50) NOT NULL COMMENT '资源分类，如编程语言、算法、工具等',
    filePath VARCHAR(255) NOT NULL COMMENT '资源文件存储路径或 URL',
    uploaderId INT UNSIGNED DEFAULT NULL COMMENT '上传者ID，关联 users 表',
    downloads INT DEFAULT 0 COMMENT '下载次数',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    CONSTRAINT fk_resource_uploader FOREIGN KEY (uploaderId) REFERENCES users(userId) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储各类下载资源';

-- =====================================================
-- 9. 项目表
DROP TABLE IF EXISTS projects;
CREATE TABLE projects (
    projectId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '项目ID，主键，自增',
    name VARCHAR(255) NOT NULL COMMENT '项目名称',
    description TEXT DEFAULT NULL COMMENT '项目描述信息',
    techStack VARCHAR(255) DEFAULT NULL COMMENT '使用的技术栈，如 GoFrame2, Vue3 等',
    coverImage VARCHAR(255) DEFAULT NULL COMMENT '项目封面图片URL',
    link VARCHAR(255) DEFAULT NULL COMMENT '项目链接或跳转地址',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储协会各项目展示';

-- =====================================================
-- 10. 评论表
DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
    commentId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '评论ID，主键，自增',
    contentType VARCHAR(50) NOT NULL COMMENT '评论关联内容类型，如 "news" 或 "activity"',
    contentId INT UNSIGNED NOT NULL COMMENT '关联内容ID，指向具体的新闻或活动',
    userId INT UNSIGNED NOT NULL COMMENT '评论用户ID，关联 users 表',
    content TEXT NOT NULL COMMENT '评论内容',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    CONSTRAINT fk_comment_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储用户评论记录';

-- =====================================================
-- 11. 消息表
DROP TABLE IF EXISTS messages;
CREATE TABLE messages (
    messageId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '消息ID，主键，自增',
    senderId INT UNSIGNED NOT NULL COMMENT '发送者ID，关联 users 表',
    receiverId INT UNSIGNED NOT NULL COMMENT '接收者ID，关联 users 表',
    content TEXT NOT NULL COMMENT '消息内容',
    readStatus TINYINT(1) DEFAULT 0 COMMENT '读取状态：0-未读, 1-已读',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录发送时间',
    CONSTRAINT fk_message_sender FOREIGN KEY (senderId) REFERENCES users(userId) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_message_receiver FOREIGN KEY (receiverId) REFERENCES users(userId) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储用户私信消息';

-- =====================================================
-- 12. 通知表
DROP TABLE IF EXISTS notifications;
CREATE TABLE notifications (
    notificationId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '通知ID，主键，自增',
    userId INT UNSIGNED NOT NULL COMMENT '接收通知的用户ID，关联 users 表',
    title VARCHAR(255) NOT NULL COMMENT '通知标题',
    content TEXT NOT NULL COMMENT '通知内容',
    type VARCHAR(50) DEFAULT NULL COMMENT '通知类型，如系统通知、活动提醒等',
    readStatus TINYINT(1) DEFAULT 0 COMMENT '读取状态：0-未读, 1-已读',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录发送时间',
    CONSTRAINT fk_notification_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储用户通知记录';

-- =====================================================
-- 13. 审计日志表
DROP TABLE IF EXISTS audit_logs;
CREATE TABLE audit_logs (
    logId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '日志ID，主键，自增',
    userId INT UNSIGNED DEFAULT NULL COMMENT '操作用户ID，关联 users 表，如为空表示系统自动操作',
    action VARCHAR(255) NOT NULL COMMENT '操作名称或类型，如 "login", "update_news" 等',
    description TEXT DEFAULT NULL COMMENT '操作详细描述，记录关键信息',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '记录操作时间',
    CONSTRAINT fk_audit_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储操作审计日志，用于追踪系统操作';