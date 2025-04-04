-- 创建帖子表
DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    postId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '帖子ID',
    userId INT UNSIGNED NOT NULL COMMENT '发帖用户ID',
    content TEXT NOT NULL COMMENT '帖子内容',
    circleId INT UNSIGNED DEFAULT NULL COMMENT '所属圈子ID',
    topicId INT UNSIGNED DEFAULT NULL COMMENT '所属话题ID',
    viewCount INT UNSIGNED DEFAULT 0 COMMENT '浏览量',
    likeCount INT UNSIGNED DEFAULT 0 COMMENT '点赞数',
    commentCount INT UNSIGNED DEFAULT 0 COMMENT '评论数',
    shareCount INT UNSIGNED DEFAULT 0 COMMENT '分享数',
    isTop TINYINT(1) DEFAULT 0 COMMENT '是否置顶：0-否，1-是',
    status TINYINT(1) DEFAULT 1 COMMENT '状态：0-草稿，1-已发布，2-已删除',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_userId (userId),
    INDEX idx_circleId (circleId),
    INDEX idx_topicId (topicId),
    CONSTRAINT fk_posts_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子表';

-- 创建帖子图片表
DROP TABLE IF EXISTS postImages;
CREATE TABLE postImages (
    imageId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '图片ID',
    postId INT UNSIGNED NOT NULL COMMENT '帖子ID',
    imageUrl VARCHAR(500) NOT NULL COMMENT '图片URL',
    sortOrder INT UNSIGNED DEFAULT 0 COMMENT '排序顺序',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    INDEX idx_postId (postId),
    CONSTRAINT fk_postimages_post FOREIGN KEY (postId) REFERENCES posts(postId) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子图片表';

-- 创建圈子表
DROP TABLE IF EXISTS circles;
CREATE TABLE circles (
    circleId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '圈子ID',
    name VARCHAR(50) NOT NULL COMMENT '圈子名称',
    description VARCHAR(500) DEFAULT NULL COMMENT '圈子描述',
    icon VARCHAR(500) DEFAULT NULL COMMENT '圈子图标',
    userId INT UNSIGNED NOT NULL COMMENT '创建者ID',
    postCount INT UNSIGNED DEFAULT 0 COMMENT '帖子数',
    memberCount INT UNSIGNED DEFAULT 0 COMMENT '成员数',
    status TINYINT(1) DEFAULT 1 COMMENT '状态：0-已删除，1-正常',
    isOfficial TINYINT(1) DEFAULT 0 COMMENT '是否官方圈子：0-否，1-是',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_userId (userId),
    CONSTRAINT fk_circles_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='圈子表';

-- 创建话题表
DROP TABLE IF EXISTS topics;
CREATE TABLE topics (
    topicId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '话题ID',
    name VARCHAR(50) NOT NULL COMMENT '话题名称',
    description VARCHAR(500) DEFAULT NULL COMMENT '话题描述',
    icon VARCHAR(500) DEFAULT NULL COMMENT '话题图标',
    postCount INT UNSIGNED DEFAULT 0 COMMENT '帖子数',
    status TINYINT(1) DEFAULT 1 COMMENT '状态：0-已删除，1-正常',
    isHot TINYINT(1) DEFAULT 0 COMMENT '是否热门：0-否，1-是',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updateTime DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='话题表';

-- 创建点赞表
DROP TABLE IF EXISTS likes;
CREATE TABLE likes (
    likeId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '点赞ID',
    userId INT UNSIGNED NOT NULL COMMENT '用户ID',
    targetId INT UNSIGNED NOT NULL COMMENT '目标ID',
    targetType TINYINT(1) NOT NULL COMMENT '目标类型：1-帖子，2-评论',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE INDEX idx_user_target (userId, targetId, targetType),
    INDEX idx_target (targetId, targetType),
    CONSTRAINT fk_likes_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';

-- 创建关注表
DROP TABLE IF EXISTS follows;
CREATE TABLE follows (
    followId INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '关注ID',
    userId INT UNSIGNED NOT NULL COMMENT '用户ID',
    followedId INT UNSIGNED NOT NULL COMMENT '被关注对象ID',
    followType TINYINT(1) NOT NULL COMMENT '关注类型：1-用户，2-圈子，3-话题',
    createTime DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE INDEX idx_user_followed (userId, followedId, followType),
    INDEX idx_followed (followedId, followType),
    CONSTRAINT fk_follows_user FOREIGN KEY (userId) REFERENCES users(userId) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关注表';
