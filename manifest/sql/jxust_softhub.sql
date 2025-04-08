/*
 Navicat Premium Dump SQL

 Source Server         : mysql8
 Source Server Type    : MySQL
 Source Server Version : 80402 (8.4.2)
 Source Host           : localhost:3306
 Source Schema         : jxust_softhub

 Target Server Type    : MySQL
 Target Server Version : 80402 (8.4.2)
 File Encoding         : 65001

 Date: 08/04/2025 14:44:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activities
-- ----------------------------
DROP TABLE IF EXISTS `activities`;
CREATE TABLE `activities` (
  `activityId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '活动ID，主键，自增',
  `title` varchar(255) NOT NULL COMMENT '活动标题',
  `description` text NOT NULL COMMENT '活动详细描述',
  `startTime` datetime NOT NULL COMMENT '活动开始时间',
  `endTime` datetime NOT NULL COMMENT '活动结束时间',
  `location` varchar(255) NOT NULL COMMENT '活动举办地点',
  `maxParticipants` int DEFAULT '0' COMMENT '最大参与人数限制',
  `status` tinyint(1) DEFAULT '0' COMMENT '活动状态：0-未开始, 1-进行中, 2-已结束',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`activityId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储活动信息';

-- ----------------------------
-- Records of activities
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for activity_registrations
-- ----------------------------
DROP TABLE IF EXISTS `activity_registrations`;
CREATE TABLE `activity_registrations` (
  `registrationId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '报名记录ID，主键，自增',
  `activityId` int unsigned NOT NULL COMMENT '活动ID，关联 activities 表',
  `userId` int unsigned DEFAULT NULL COMMENT '报名用户ID，关联 users 表，如为空表示未登录报名',
  `name` varchar(50) NOT NULL COMMENT '报名者姓名',
  `studentId` varchar(20) NOT NULL COMMENT '报名者学号',
  `contact` varchar(50) NOT NULL COMMENT '报名者联系方式，如电话或邮箱',
  `status` tinyint(1) DEFAULT '0' COMMENT '报名状态：0-待审核, 1-通过, 2-拒绝',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`registrationId`),
  KEY `fk_activity` (`activityId`),
  KEY `fk_registration_user` (`userId`),
  CONSTRAINT `fk_activity` FOREIGN KEY (`activityId`) REFERENCES `activities` (`activityId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_registration_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储活动报名记录';

-- ----------------------------
-- Records of activity_registrations
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for audit_logs
-- ----------------------------
DROP TABLE IF EXISTS `audit_logs`;
CREATE TABLE `audit_logs` (
  `logId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '日志ID，主键，自增',
  `userId` int unsigned DEFAULT NULL COMMENT '操作用户ID，关联 users 表，如为空表示系统自动操作',
  `action` varchar(255) NOT NULL COMMENT '操作名称或类型，如 "login", "update_news" 等',
  `description` text COMMENT '操作详细描述，记录关键信息',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录操作时间',
  PRIMARY KEY (`logId`),
  KEY `fk_audit_user` (`userId`),
  CONSTRAINT `fk_audit_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储操作审计日志，用于追踪系统操作';

-- ----------------------------
-- Records of audit_logs
-- ----------------------------
BEGIN;
INSERT INTO `audit_logs` (`logId`, `userId`, `action`, `description`, `createTime`) VALUES (1, 1, 'approve_join_application', '{\"applicationId\":3,\"name\":\"王五\",\"studentId\":\"202110101020\"}', '2025-04-05 06:38:56');
COMMIT;

-- ----------------------------
-- Table structure for circles
-- ----------------------------
DROP TABLE IF EXISTS `circles`;
CREATE TABLE `circles` (
  `circleId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '圈子ID',
  `name` varchar(50) NOT NULL COMMENT '圈子名称',
  `description` varchar(500) DEFAULT NULL COMMENT '圈子描述',
  `icon` varchar(500) DEFAULT NULL COMMENT '圈子图标',
  `userId` int unsigned NOT NULL COMMENT '创建者ID',
  `postCount` int unsigned DEFAULT '0' COMMENT '帖子数',
  `memberCount` int unsigned DEFAULT '0' COMMENT '成员数',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：0-已删除，1-正常',
  `isOfficial` tinyint(1) DEFAULT '0' COMMENT '是否官方圈子：0-否，1-是',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`circleId`),
  KEY `idx_userId` (`userId`),
  CONSTRAINT `fk_circles_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='圈子表';

-- ----------------------------
-- Records of circles
-- ----------------------------
BEGIN;
INSERT INTO `circles` (`circleId`, `name`, `description`, `icon`, `userId`, `postCount`, `memberCount`, `status`, `isOfficial`, `createTime`, `updateTime`) VALUES (1, '默认圈子', '系统默认圈子', NULL, 1, 0, 0, 1, 1, '2025-04-06 11:22:21', '2025-04-06 11:22:21');
COMMIT;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `commentId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID，主键，自增',
  `contentType` varchar(50) NOT NULL COMMENT '评论关联内容类型，如 "news" 或 "activity"',
  `contentId` int unsigned NOT NULL COMMENT '关联内容ID，指向具体的新闻或活动',
  `userId` int unsigned NOT NULL COMMENT '评论用户ID，关联 users 表',
  `content` text NOT NULL COMMENT '评论内容',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `likeCount` int unsigned DEFAULT '0' COMMENT '点赞数',
  PRIMARY KEY (`commentId`),
  KEY `fk_comment_user` (`userId`),
  CONSTRAINT `fk_comment_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储用户评论记录';

-- ----------------------------
-- Records of comments
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for content_tags
-- ----------------------------
DROP TABLE IF EXISTS `content_tags`;
CREATE TABLE `content_tags` (
  `contentType` varchar(50) NOT NULL COMMENT '内容类型，如news、project、resource等',
  `contentId` int unsigned NOT NULL COMMENT '内容ID',
  `tagId` int unsigned NOT NULL COMMENT '标签ID，关联tags表',
  PRIMARY KEY (`contentType`,`contentId`,`tagId`),
  KEY `fk_content_tag` (`tagId`),
  CONSTRAINT `fk_content_tag` FOREIGN KEY (`tagId`) REFERENCES `tags` (`tagId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='关联内容与标签';

-- ----------------------------
-- Records of content_tags
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `fileId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '文件ID，主键，自增',
  `fileName` varchar(255) NOT NULL COMMENT '文件名称',
  `fileSize` bigint unsigned NOT NULL COMMENT '文件大小（字节）',
  `fileType` varchar(100) NOT NULL COMMENT '文件类型/MIME类型',
  `filePath` varchar(255) NOT NULL COMMENT '文件存储路径',
  `uploaderId` int unsigned DEFAULT NULL COMMENT '上传者ID，关联users表',
  `uploadTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `md5Hash` varchar(32) DEFAULT NULL COMMENT '文件MD5哈希值，用于去重',
  `relatedType` varchar(50) DEFAULT NULL COMMENT '关联类型，如news、activity、resource等',
  `relatedId` int unsigned DEFAULT NULL COMMENT '关联ID，指向具体的内容ID',
  PRIMARY KEY (`fileId`),
  KEY `fk_file_uploader` (`uploaderId`),
  CONSTRAINT `fk_file_uploader` FOREIGN KEY (`uploaderId`) REFERENCES `users` (`userId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储上传的文件信息';

-- ----------------------------
-- Records of files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
  `followId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '关注ID',
  `userId` int unsigned NOT NULL COMMENT '用户ID',
  `followedId` int unsigned NOT NULL COMMENT '被关注对象ID',
  `followType` tinyint(1) NOT NULL COMMENT '关注类型：1-用户，2-圈子，3-话题',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`followId`),
  UNIQUE KEY `idx_user_followed` (`userId`,`followedId`,`followType`),
  KEY `idx_followed` (`followedId`,`followType`),
  CONSTRAINT `fk_follows_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='关注表';

-- ----------------------------
-- Records of follows
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for join_applications
-- ----------------------------
DROP TABLE IF EXISTS `join_applications`;
CREATE TABLE `join_applications` (
  `applicationId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '申请ID，主键，自增',
  `name` varchar(50) NOT NULL COMMENT '申请人姓名',
  `studentId` varchar(20) NOT NULL COMMENT '学号',
  `grade` varchar(20) NOT NULL COMMENT '年级，如2020级',
  `college` varchar(100) NOT NULL COMMENT '学院',
  `major` varchar(100) NOT NULL COMMENT '专业',
  `phone` varchar(20) NOT NULL COMMENT '联系电话',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `reason` text NOT NULL COMMENT '申请理由',
  `skills` text COMMENT '技能介绍',
  `expectDepartment` varchar(50) DEFAULT NULL COMMENT '期望加入的部门',
  `status` tinyint(1) DEFAULT '0' COMMENT '申请状态：0-待审核，1-通过，2-拒绝',
  `reviewerId` int unsigned DEFAULT NULL COMMENT '审核人ID，关联users表',
  `reviewComment` text COMMENT '审核意见',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`applicationId`),
  KEY `fk_application_reviewer` (`reviewerId`),
  CONSTRAINT `fk_application_reviewer` FOREIGN KEY (`reviewerId`) REFERENCES `users` (`userId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储入会申请记录';

-- ----------------------------
-- Records of join_applications
-- ----------------------------
BEGIN;
INSERT INTO `join_applications` (`applicationId`, `name`, `studentId`, `grade`, `college`, `major`, `phone`, `email`, `reason`, `skills`, `expectDepartment`, `status`, `reviewerId`, `reviewComment`, `createTime`, `updateTime`) VALUES (2, '李明', '20250001', '2025级', '计算机学院', '软件工程', '13800138000', 'liming@example.com', '我热爱编程，想加入软件协会提升技能', '熟悉Java, Go语言, 有Web开发经验', '技术部', 0, NULL, NULL, '2025-04-03 10:36:43', '2025-04-03 10:36:43');
INSERT INTO `join_applications` (`applicationId`, `name`, `studentId`, `grade`, `college`, `major`, `phone`, `email`, `reason`, `skills`, `expectDepartment`, `status`, `reviewerId`, `reviewComment`, `createTime`, `updateTime`) VALUES (3, '王五', '202110101020', '2021级', '计算机学院', '软件工程', '13300002222', 'wangwu@test.com', '我热爱编程，希望能加入软件协会提升自己的能力', '熟悉Java、Python编程语言，有一定的前端开发经验', '技术部', 1, 1, '申请资料完善，技术背景良好，通过申请', '2025-04-05 06:35:17', '2025-04-04 22:38:56');
INSERT INTO `join_applications` (`applicationId`, `name`, `studentId`, `grade`, `college`, `major`, `phone`, `email`, `reason`, `skills`, `expectDepartment`, `status`, `reviewerId`, `reviewComment`, `createTime`, `updateTime`) VALUES (4, '赵六', '123456678', '2022级', '计算机学院', '软件工程', '13300006666', 'zhaoliu@gmail.com', '想要加强自己的能力', '精通ps，pr', '设计部', 0, NULL, NULL, '2025-04-05 07:27:53', '2025-04-05 07:27:53');
COMMIT;

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
  `likeId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '点赞ID',
  `userId` int unsigned NOT NULL COMMENT '用户ID',
  `targetId` int unsigned NOT NULL COMMENT '目标ID',
  `targetType` tinyint(1) NOT NULL COMMENT '目标类型：1-帖子，2-评论',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`likeId`),
  UNIQUE KEY `idx_user_target` (`userId`,`targetId`,`targetType`),
  KEY `idx_target` (`targetId`,`targetType`),
  CONSTRAINT `fk_likes_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='点赞表';

-- ----------------------------
-- Records of likes
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for members
-- ----------------------------
DROP TABLE IF EXISTS `members`;
CREATE TABLE `members` (
  `memberId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '成员ID，主键，自增',
  `userId` int unsigned NOT NULL COMMENT '关联的用户ID',
  `grade` varchar(20) NOT NULL COMMENT '年级，如2020级',
  `joinYear` year NOT NULL COMMENT '入会年份',
  `department` varchar(50) DEFAULT NULL COMMENT '所属部门，如技术部、宣传部等',
  `position` varchar(50) DEFAULT NULL COMMENT '职位，如部长、副部长、干事等',
  `skills` text COMMENT '技能描述',
  `introduction` text COMMENT '个人简介',
  `isCore` tinyint(1) DEFAULT '0' COMMENT '是否为核心成员：0-否，1-是',
  `status` tinyint(1) DEFAULT '1' COMMENT '成员状态：0-待审核，1-正式成员，2-已退出',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`memberId`),
  KEY `fk_member_user` (`userId`),
  CONSTRAINT `fk_member_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储协会成员详细信息';

-- ----------------------------
-- Records of members
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `messageId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '消息ID，主键，自增',
  `senderId` int unsigned NOT NULL COMMENT '发送者ID，关联 users 表',
  `receiverId` int unsigned NOT NULL COMMENT '接收者ID，关联 users 表',
  `content` text NOT NULL COMMENT '消息内容',
  `readStatus` tinyint(1) DEFAULT '0' COMMENT '读取状态：0-未读, 1-已读',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录发送时间',
  PRIMARY KEY (`messageId`),
  KEY `fk_message_sender` (`senderId`),
  KEY `fk_message_receiver` (`receiverId`),
  CONSTRAINT `fk_message_receiver` FOREIGN KEY (`receiverId`) REFERENCES `users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_message_sender` FOREIGN KEY (`senderId`) REFERENCES `users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储用户私信消息';

-- ----------------------------
-- Records of messages
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `newsId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '新闻ID，主键，自增',
  `title` varchar(255) NOT NULL COMMENT '新闻标题',
  `content` text NOT NULL COMMENT '新闻内容，支持 Markdown 格式',
  `category` varchar(50) NOT NULL COMMENT '新闻分类，如协会新闻、技术分享、赛事通知等',
  `newsType` tinyint(1) DEFAULT '1' COMMENT '新闻类型：1-协会通知，2-技术分享',
  `coverImage` varchar(255) DEFAULT NULL COMMENT '封面图片的URL',
  `authorId` int unsigned DEFAULT NULL COMMENT '作者ID，关联 users 表',
  `viewCount` int DEFAULT '0' COMMENT '浏览次数',
  `status` tinyint(1) DEFAULT '0' COMMENT '新闻状态，0: 草稿, 1: 发布',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`newsId`),
  KEY `fk_news_author` (`authorId`),
  CONSTRAINT `fk_news_author` FOREIGN KEY (`authorId`) REFERENCES `users` (`userId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储新闻资讯内容';

-- ----------------------------
-- Records of news
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for notifications
-- ----------------------------
DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications` (
  `notificationId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '通知ID，主键，自增',
  `userId` int unsigned NOT NULL COMMENT '接收通知的用户ID，关联 users 表',
  `title` varchar(255) NOT NULL COMMENT '通知标题',
  `content` text NOT NULL COMMENT '通知内容',
  `type` varchar(50) DEFAULT NULL COMMENT '通知类型，如系统通知、活动提醒等',
  `readStatus` tinyint(1) DEFAULT '0' COMMENT '读取状态：0-未读, 1-已读',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录发送时间',
  PRIMARY KEY (`notificationId`),
  KEY `fk_notification_user` (`userId`),
  CONSTRAINT `fk_notification_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储用户通知记录';

-- ----------------------------
-- Records of notifications
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions` (
  `permissionId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '权限ID，主键，自增',
  `permissionKey` varchar(50) NOT NULL COMMENT '权限标识，用于后台校验，如 "news_edit" ',
  `description` varchar(255) DEFAULT NULL COMMENT '权限描述信息',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`permissionId`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储系统各项权限';

-- ----------------------------
-- Records of permissions
-- ----------------------------
BEGIN;
INSERT INTO `permissions` (`permissionId`, `permissionKey`, `description`, `createTime`, `updateTime`) VALUES (5, 'publish_tech_article', '发布技术文章的权限', '2025-04-05 05:51:07', '2025-04-05 05:51:07');
INSERT INTO `permissions` (`permissionId`, `permissionKey`, `description`, `createTime`, `updateTime`) VALUES (6, 'publish_association_notice', '发布协会通知的权限', '2025-04-05 05:51:07', '2025-04-05 05:51:07');
COMMIT;

-- ----------------------------
-- Table structure for postImages
-- ----------------------------
DROP TABLE IF EXISTS `postImages`;
CREATE TABLE `postImages` (
  `imageId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '图片ID',
  `postId` int unsigned NOT NULL COMMENT '帖子ID',
  `imageUrl` varchar(500) NOT NULL COMMENT '图片URL',
  `sortOrder` int unsigned DEFAULT '0' COMMENT '排序顺序',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`imageId`),
  KEY `idx_postId` (`postId`),
  CONSTRAINT `fk_postimages_post` FOREIGN KEY (`postId`) REFERENCES `posts` (`postId`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='帖子图片表';

-- ----------------------------
-- Records of postImages
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `postId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
  `userId` int unsigned NOT NULL COMMENT '发帖用户ID',
  `content` text NOT NULL COMMENT '帖子内容',
  `circleId` int unsigned DEFAULT NULL COMMENT '所属圈子ID',
  `topicId` int unsigned DEFAULT NULL COMMENT '所属话题ID',
  `viewCount` int unsigned DEFAULT '0' COMMENT '浏览量',
  `likeCount` int unsigned DEFAULT '0' COMMENT '点赞数',
  `commentCount` int unsigned DEFAULT '0' COMMENT '评论数',
  `shareCount` int unsigned DEFAULT '0' COMMENT '分享数',
  `isTop` tinyint(1) DEFAULT '0' COMMENT '是否置顶：0-否，1-是',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：0-草稿，1-已发布，2-已删除',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`postId`),
  KEY `idx_userId` (`userId`),
  KEY `idx_circleId` (`circleId`),
  KEY `idx_topicId` (`topicId`),
  CONSTRAINT `fk_posts_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='帖子表';

-- ----------------------------
-- Records of posts
-- ----------------------------
BEGIN;
INSERT INTO `posts` (`postId`, `userId`, `content`, `circleId`, `topicId`, `viewCount`, `likeCount`, `commentCount`, `shareCount`, `isTop`, `status`, `createTime`, `updateTime`) VALUES (1, 1, '这是手动插入的测试帖子', 1, 1, 0, 0, 0, 0, 0, 1, '2025-04-06 11:51:16', '2025-04-06 11:51:16');
INSERT INTO `posts` (`postId`, `userId`, `content`, `circleId`, `topicId`, `viewCount`, `likeCount`, `commentCount`, `shareCount`, `isTop`, `status`, `createTime`, `updateTime`) VALUES (2, 1, '手动插入的测试帖子', 1, NULL, 0, 0, 0, 0, 0, 1, '2025-04-06 15:42:51', '2025-04-06 15:42:51');
INSERT INTO `posts` (`postId`, `userId`, `content`, `circleId`, `topicId`, `viewCount`, `likeCount`, `commentCount`, `shareCount`, `isTop`, `status`, `createTime`, `updateTime`) VALUES (3, 1, '再次手动插入的测试帖子', 1, NULL, 0, 0, 0, 0, 0, 1, '2025-04-06 15:46:43', '2025-04-06 15:46:43');
COMMIT;

-- ----------------------------
-- Table structure for projects
-- ----------------------------
DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects` (
  `projectId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '项目ID，主键，自增',
  `name` varchar(255) NOT NULL COMMENT '项目名称',
  `description` text COMMENT '项目描述信息',
  `techStack` varchar(255) DEFAULT NULL COMMENT '使用的技术栈，如 GoFrame2, Vue3 等',
  `coverImage` varchar(255) DEFAULT NULL COMMENT '项目封面图片URL',
  `link` varchar(255) DEFAULT NULL COMMENT '项目链接或跳转地址',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`projectId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储协会各项目展示';

-- ----------------------------
-- Records of projects
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for resources
-- ----------------------------
DROP TABLE IF EXISTS `resources`;
CREATE TABLE `resources` (
  `resourceId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '资源ID，主键，自增',
  `title` varchar(255) NOT NULL COMMENT '资源标题',
  `description` text COMMENT '资源描述信息',
  `category` varchar(50) NOT NULL COMMENT '资源分类，如编程语言、算法、工具等',
  `filePath` varchar(255) NOT NULL COMMENT '资源文件存储路径或 URL',
  `uploaderId` int unsigned DEFAULT NULL COMMENT '上传者ID，关联 users 表',
  `downloads` int DEFAULT '0' COMMENT '下载次数',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  PRIMARY KEY (`resourceId`),
  KEY `fk_resource_uploader` (`uploaderId`),
  CONSTRAINT `fk_resource_uploader` FOREIGN KEY (`uploaderId`) REFERENCES `users` (`userId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储各类下载资源';

-- ----------------------------
-- Records of resources
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions` (
  `roleId` int unsigned NOT NULL COMMENT '角色ID，关联 roles 表',
  `permissionId` int unsigned NOT NULL COMMENT '权限ID，关联 permissions 表',
  PRIMARY KEY (`roleId`,`permissionId`),
  KEY `fk_permission` (`permissionId`),
  CONSTRAINT `fk_permission` FOREIGN KEY (`permissionId`) REFERENCES `permissions` (`permissionId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_role` FOREIGN KEY (`roleId`) REFERENCES `roles` (`roleId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='关联角色与权限';

-- ----------------------------
-- Records of role_permissions
-- ----------------------------
BEGIN;
INSERT INTO `role_permissions` (`roleId`, `permissionId`) VALUES (3, 5);
INSERT INTO `role_permissions` (`roleId`, `permissionId`) VALUES (2, 6);
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `roleId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID，主键，自增',
  `roleName` varchar(50) NOT NULL COMMENT '角色名称，如超级管理员、内容管理员等',
  `description` varchar(255) DEFAULT NULL COMMENT '角色描述信息',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`roleId`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储角色信息';

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` (`roleId`, `roleName`, `description`, `createTime`, `updateTime`) VALUES (1, '超级管理员', '具有系统最高权限，可以管理所有功能和用户', '2025-04-03 09:30:32', '2025-04-03 09:30:32');
INSERT INTO `roles` (`roleId`, `roleName`, `description`, `createTime`, `updateTime`) VALUES (2, '内容管理员', '负责管理网站内容，包括新闻、活动等', '2025-04-03 09:30:32', '2025-04-03 09:30:32');
INSERT INTO `roles` (`roleId`, `roleName`, `description`, `createTime`, `updateTime`) VALUES (3, '会长', '协会会长，具有管理协会成员和内容的权限', '2025-04-03 09:30:32', '2025-04-03 09:30:32');
INSERT INTO `roles` (`roleId`, `roleName`, `description`, `createTime`, `updateTime`) VALUES (4, '普通用户', '普通注册用户，仅有基本浏览和参与功能', '2025-04-03 09:30:32', '2025-04-03 09:30:32');
INSERT INTO `roles` (`roleId`, `roleName`, `description`, `createTime`, `updateTime`) VALUES (5, '会员', '协会会员，可以发布技术博客和参加活动', '2025-04-05 05:59:59', '2025-04-05 05:59:59');
COMMIT;

-- ----------------------------
-- Table structure for system_configs
-- ----------------------------
DROP TABLE IF EXISTS `system_configs`;
CREATE TABLE `system_configs` (
  `configKey` varchar(100) NOT NULL COMMENT '配置键名',
  `configValue` text NOT NULL COMMENT '配置值',
  `description` varchar(255) DEFAULT NULL COMMENT '配置说明',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`configKey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储系统配置参数';

-- ----------------------------
-- Records of system_configs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `tagId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '标签ID，主键，自增',
  `tagName` varchar(50) NOT NULL COMMENT '标签名称',
  `tagType` varchar(50) DEFAULT NULL COMMENT '标签类型，如news、project、resource等',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`tagId`),
  UNIQUE KEY `idx_tag_name_type` (`tagName`,`tagType`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储内容标签';

-- ----------------------------
-- Records of tags
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for topics
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `topicId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '话题ID',
  `name` varchar(50) NOT NULL COMMENT '话题名称',
  `description` varchar(500) DEFAULT NULL COMMENT '话题描述',
  `icon` varchar(500) DEFAULT NULL COMMENT '话题图标',
  `postCount` int unsigned DEFAULT '0' COMMENT '帖子数',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：0-已删除，1-正常',
  `isHot` tinyint(1) DEFAULT '0' COMMENT '是否热门：0-否，1-是',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`topicId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='话题表';

-- ----------------------------
-- Records of topics
-- ----------------------------
BEGIN;
INSERT INTO `topics` (`topicId`, `name`, `description`, `icon`, `postCount`, `status`, `isHot`, `createTime`, `updateTime`) VALUES (1, '默认话题', '系统默认话题', NULL, 0, 1, 0, '2025-04-06 11:22:21', '2025-04-06 11:22:21');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `userId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID，主键，自增',
  `userName` varchar(50) NOT NULL COMMENT '用户名，登录和显示名称',
  `password` varchar(100) NOT NULL COMMENT '用户密码，存储加密后的密码',
  `name` varchar(50) DEFAULT NULL COMMENT '真实姓名',
  `roleId` int unsigned DEFAULT NULL COMMENT '角色ID，关联 roles 表，标识用户所属角色',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像图片URL',
  `joinYear` year DEFAULT NULL COMMENT '入会年份，格式如2025',
  `email` varchar(100) DEFAULT NULL COMMENT '用户邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '用户联系电话',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  `followerCount` int unsigned DEFAULT '0' COMMENT '粉丝数量',
  `followingCount` int unsigned DEFAULT '0' COMMENT '关注数量',
  PRIMARY KEY (`userId`),
  KEY `fk_user_role` (`roleId`),
  CONSTRAINT `fk_user_role` FOREIGN KEY (`roleId`) REFERENCES `roles` (`roleId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储用户基本信息';

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`userId`, `userName`, `password`, `name`, `roleId`, `avatar`, `joinYear`, `email`, `phone`, `createTime`, `updateTime`, `followerCount`, `followingCount`) VALUES (1, 'admin', '0192023a7bbd73250516f069df18b500', '超级管理员', 1, NULL, NULL, 'admin@example.com', '13800000000', '2025-04-03 09:30:32', '2025-04-03 09:30:32', 0, 0);
INSERT INTO `users` (`userId`, `userName`, `password`, `name`, `roleId`, `avatar`, `joinYear`, `email`, `phone`, `createTime`, `updateTime`, `followerCount`, `followingCount`) VALUES (4, 'student001', 'ad6a280417a0f533d8b670c61667e1a0', '李明', 2, NULL, NULL, 'liming@example.com', '13800138000', '2025-04-03 10:20:02', '2025-04-03 03:06:26', 0, 0);
INSERT INTO `users` (`userId`, `userName`, `password`, `name`, `roleId`, `avatar`, `joinYear`, `email`, `phone`, `createTime`, `updateTime`, `followerCount`, `followingCount`) VALUES (5, 'zhangsan', '482c811da5d5b4bc6d497ffa98491e38', '张三', 4, NULL, NULL, 'zhagnsan@example.com', '13800138111', '2025-04-03 13:01:30', '2025-04-06 08:50:48', 0, 0);
INSERT INTO `users` (`userId`, `userName`, `password`, `name`, `roleId`, `avatar`, `joinYear`, `email`, `phone`, `createTime`, `updateTime`, `followerCount`, `followingCount`) VALUES (6, 'lisi', '482c811da5d5b4bc6d497ffa98491e38', '李四', 4, NULL, NULL, 'lisi@example.com', '13800132222', '2025-04-04 09:08:43', '2025-04-06 08:51:16', 0, 0);
INSERT INTO `users` (`userId`, `userName`, `password`, `name`, `roleId`, `avatar`, `joinYear`, `email`, `phone`, `createTime`, `updateTime`, `followerCount`, `followingCount`) VALUES (7, 'wangwu', '0192023a7bbd73250516f069df18b500', '王五', 4, NULL, NULL, 'wangwu@test.com', '13300002222', '2025-04-04 09:27:25', '2025-04-05 06:34:31', 0, 0);
INSERT INTO `users` (`userId`, `userName`, `password`, `name`, `roleId`, `avatar`, `joinYear`, `email`, `phone`, `createTime`, `updateTime`, `followerCount`, `followingCount`) VALUES (8, 'zhaoliu', '482c811da5d5b4bc6d497ffa98491e38', '赵六', 4, NULL, NULL, 'zhaoliu@gmail.com', '13300006666', '2025-04-05 07:26:39', '2025-04-05 07:26:39', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for visit_statistics
-- ----------------------------
DROP TABLE IF EXISTS `visit_statistics`;
CREATE TABLE `visit_statistics` (
  `statId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '统计ID，主键，自增',
  `visitDate` date NOT NULL COMMENT '访问日期',
  `pageView` int unsigned NOT NULL DEFAULT '0' COMMENT '页面浏览量',
  `uniqueVisitor` int unsigned NOT NULL DEFAULT '0' COMMENT '独立访客数',
  `newVisitor` int unsigned NOT NULL DEFAULT '0' COMMENT '新访客数',
  `avgStayTime` int unsigned DEFAULT '0' COMMENT '平均停留时间（秒）',
  `bounceRate` decimal(5,2) DEFAULT '0.00' COMMENT '跳出率（百分比）',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`statId`),
  UNIQUE KEY `idx_visit_date` (`visitDate`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储网站访问统计数据';

-- ----------------------------
-- Records of visit_statistics
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Triggers structure for table follows
-- ----------------------------
DROP TRIGGER IF EXISTS `after_follow_insert`;
delimiter ;;
CREATE TRIGGER `jxust_softhub`.`after_follow_insert` AFTER INSERT ON `follows` FOR EACH ROW BEGIN
    -- 如果是关注用户类型(followType=1)
    IF NEW.followType = 1 THEN
        -- 增加被关注用户的粉丝数
        UPDATE users SET followerCount = followerCount + 1 WHERE userId = NEW.followedId;
        -- 增加关注者的关注数
        UPDATE users SET followingCount = followingCount + 1 WHERE userId = NEW.userId;
    END IF;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table follows
-- ----------------------------
DROP TRIGGER IF EXISTS `after_follow_delete`;
delimiter ;;
CREATE TRIGGER `jxust_softhub`.`after_follow_delete` AFTER DELETE ON `follows` FOR EACH ROW BEGIN
    -- 如果是关注用户类型(followType=1)
    IF OLD.followType = 1 THEN
        -- 减少被关注用户的粉丝数
        UPDATE users SET followerCount = followerCount - 1 WHERE userId = OLD.followedId AND followerCount > 0;
        -- 减少关注者的关注数
        UPDATE users SET followingCount = followingCount - 1 WHERE userId = OLD.userId AND followingCount > 0;
    END IF;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
