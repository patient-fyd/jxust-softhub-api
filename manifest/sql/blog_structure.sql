-- ----------------------------
-- Table structure for blogs
-- ----------------------------
DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs` (
  `blogId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '博客ID，主键，自增',
  `title` varchar(255) NOT NULL COMMENT '博客标题',
  `content` text NOT NULL COMMENT '博客内容，支持 Markdown 格式',
  `summary` varchar(500) DEFAULT NULL COMMENT '博客摘要',
  `category` varchar(50) NOT NULL COMMENT '博客分类，如前端、后端、移动开发、算法等',
  `tags` varchar(255) DEFAULT NULL COMMENT '博客标签，多个标签用逗号分隔',
  `coverImage` varchar(255) DEFAULT NULL COMMENT '封面图片的URL',
  `authorId` int unsigned DEFAULT NULL COMMENT '作者ID，关联 users 表',
  `viewCount` int DEFAULT '0' COMMENT '浏览次数',
  `likeCount` int DEFAULT '0' COMMENT '点赞次数',
  `commentCount` int DEFAULT '0' COMMENT '评论次数',
  `isRecommend` tinyint(1) DEFAULT '0' COMMENT '是否推荐：0-否，1-是',
  `status` tinyint(1) DEFAULT '1' COMMENT '博客状态，0: 草稿, 1: 发布, 2: 下架',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`blogId`),
  KEY `idx_author` (`authorId`),
  KEY `idx_category` (`category`),
  KEY `idx_status` (`status`),
  KEY `idx_recommend` (`isRecommend`),
  CONSTRAINT `fk_blog_author` FOREIGN KEY (`authorId`) REFERENCES `users` (`userId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储技术博客内容';

-- ----------------------------
-- Records of blogs (示例数据)
-- ----------------------------
BEGIN;
INSERT INTO `blogs` (`blogId`, `title`, `content`, `summary`, `category`, `tags`, `coverImage`, `authorId`, `viewCount`, `likeCount`, `commentCount`, `isRecommend`, `status`, `createTime`, `updateTime`) VALUES 
(1, 'GoFrame框架入门指南', '# GoFrame框架入门指南\n\nGoFrame是一款模块化、高性能、企业级的Go基础开发框架。本文将带你快速入门GoFrame框架。\n\n## 安装\n\n```bash\ngo get -u github.com/gogf/gf/v2\n```\n\n## 创建项目\n\n使用GoFrame CLI工具创建项目：\n\n```bash\ngf init my-project\ncd my-project\ngo mod tidy\n```\n\n## 启动服务\n\n```bash\ngf run main.go\n```\n\n更多详情，请访问[GoFrame官方文档](https://goframe.org/display/gf)', 'GoFrame是一款模块化、高性能、企业级的Go基础开发框架。本文将带你快速入门GoFrame框架，包括安装、创建项目和启动服务等基本操作。', '后端开发', 'Go,GoFrame,Web框架', '/uploads/blog/goframe-cover.jpg', 1, 128, 45, 12, 1, 1, '2025-04-10 08:30:00', '2025-04-10 08:30:00'),
(2, 'Vue3 组合式API完全指南', '# Vue3 组合式API完全指南\n\nVue3引入了组合式API（Composition API），这是一种全新的逻辑复用和代码组织方式。\n\n## 什么是组合式API？\n\n组合式API是Vue3中新增的一组API，允许我们使用函数而不是声明选项的方式书写Vue组件。\n\n## setup函数\n\n```javascript\nimport { ref, onMounted } from \'vue\';\n\nexport default {\n  setup() {\n    const count = ref(0);\n    \n    function increment() {\n      count.value++;\n    }\n    \n    onMounted(() => {\n      console.log(\'Component mounted!\');\n    });\n    \n    return {\n      count,\n      increment\n    };\n  }\n};\n```\n\n## 响应式API\n\n- ref：用于基本类型的响应式\n- reactive：用于对象类型的响应式\n- computed：计算属性\n- watch/watchEffect：侦听器\n\n更多详情，请查看[Vue3官方文档](https://v3.vuejs.org/guide/composition-api-introduction.html)', 'Vue3引入了组合式API（Composition API），这是一种全新的逻辑复用和代码组织方式。本文将深入介绍setup函数和响应式API的使用方法。', '前端开发', 'Vue3,JavaScript,前端框架', '/uploads/blog/vue3-cover.jpg', 4, 256, 78, 23, 1, 1, '2025-04-11 10:15:00', '2025-04-11 10:15:00'),
(3, '深入理解JavaScript异步编程', '# 深入理解JavaScript异步编程\n\n异步编程是JavaScript中非常重要的概念，本文将深入讲解JavaScript中的异步编程模式。\n\n## 回调函数\n\n```javascript\nfunction fetchData(callback) {\n  setTimeout(() => {\n    callback(\'Data received\');\n  }, 1000);\n}\n\nfetchData((data) => {\n  console.log(data); // 输出: Data received\n});\n```\n\n## Promise\n\n```javascript\nfunction fetchData() {\n  return new Promise((resolve, reject) => {\n    setTimeout(() => {\n      resolve(\'Data received\');\n    }, 1000);\n  });\n}\n\nfetchData()\n  .then(data => {\n    console.log(data); // 输出: Data received\n  })\n  .catch(error => {\n    console.error(error);\n  });\n```\n\n## Async/Await\n\n```javascript\nasync function getData() {\n  try {\n    const data = await fetchData();\n    console.log(data); // 输出: Data received\n  } catch (error) {\n    console.error(error);\n  }\n}\n\ngetData();\n```\n\n## 总结\n\n异步编程从回调函数，到Promise，再到Async/Await，JavaScript的异步编程模式不断进化，使得代码更加简洁、可读性更强。', '异步编程是JavaScript中非常重要的概念，本文将从回调函数、Promise到Async/Await深入讲解JavaScript中的异步编程模式演进历程。', '前端开发', 'JavaScript,异步编程,Promise', '/uploads/blog/js-async-cover.jpg', 6, 189, 56, 18, 0, 1, '2025-04-12 14:20:00', '2025-04-12 14:20:00');
COMMIT;

-- ----------------------------
-- Table structure for blog_comments
-- ----------------------------
DROP TABLE IF EXISTS `blog_comments`;
CREATE TABLE `blog_comments` (
  `commentId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID，主键，自增',
  `blogId` int unsigned NOT NULL COMMENT '博客ID，关联blogs表',
  `userId` int unsigned NOT NULL COMMENT '评论用户ID，关联users表',
  `content` text NOT NULL COMMENT '评论内容',
  `parentId` int unsigned DEFAULT NULL COMMENT '父评论ID，用于回复功能，如为NULL则为顶级评论',
  `likeCount` int unsigned DEFAULT '0' COMMENT '点赞数',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：0-已删除，1-正常',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录最后更新时间',
  PRIMARY KEY (`commentId`),
  KEY `idx_blog` (`blogId`),
  KEY `idx_user` (`userId`),
  KEY `idx_parent` (`parentId`),
  CONSTRAINT `fk_blog_comment_blog` FOREIGN KEY (`blogId`) REFERENCES `blogs` (`blogId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_blog_comment_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_blog_comment_parent` FOREIGN KEY (`parentId`) REFERENCES `blog_comments` (`commentId`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储博客评论';

-- ----------------------------
-- Records of blog_comments (示例数据)
-- ----------------------------
BEGIN;
INSERT INTO `blog_comments` (`commentId`, `blogId`, `userId`, `content`, `parentId`, `likeCount`, `status`, `createTime`) VALUES 
(1, 1, 6, '非常实用的GoFrame入门教程，对初学者很友好！', NULL, 5, 1, '2025-04-10 10:30:00'),
(2, 1, 7, '请问GoFrame和Gin相比有什么优势？', NULL, 2, 1, '2025-04-10 11:45:00'),
(3, 1, 1, '主要在于更完整的企业级功能支持和更便捷的开发体验，GoFrame是一站式解决方案。', 2, 3, 1, '2025-04-10 12:20:00'),
(4, 2, 5, 'Vue3的组合式API确实解决了很多Vue2中的问题，特别是大型组件的逻辑组织问题。', NULL, 8, 1, '2025-04-11 16:10:00'),
(5, 3, 8, '通俗易懂的异步编程讲解，终于理解了async/await的工作原理！', NULL, 6, 1, '2025-04-12 20:05:00');
COMMIT;

-- ----------------------------
-- Table structure for blog_likes
-- ----------------------------
DROP TABLE IF EXISTS `blog_likes`;
CREATE TABLE `blog_likes` (
  `likeId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '点赞ID，主键，自增',
  `blogId` int unsigned NOT NULL COMMENT '博客ID，关联blogs表',
  `userId` int unsigned NOT NULL COMMENT '用户ID，关联users表',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '点赞时间',
  PRIMARY KEY (`likeId`),
  UNIQUE KEY `uk_blog_user` (`blogId`,`userId`),
  KEY `idx_user` (`userId`),
  CONSTRAINT `fk_blog_like_blog` FOREIGN KEY (`blogId`) REFERENCES `blogs` (`blogId`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_blog_like_user` FOREIGN KEY (`userId`) REFERENCES `users` (`userId`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='存储博客点赞记录';

-- ----------------------------
-- Triggers structure for table blog_likes
-- ----------------------------
DROP TRIGGER IF EXISTS `after_blog_like_insert`;
delimiter ;;
CREATE TRIGGER `after_blog_like_insert` AFTER INSERT ON `blog_likes` FOR EACH ROW BEGIN
  UPDATE blogs SET likeCount = likeCount + 1 WHERE blogId = NEW.blogId;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table blog_likes
-- ----------------------------
DROP TRIGGER IF EXISTS `after_blog_like_delete`;
delimiter ;;
CREATE TRIGGER `after_blog_like_delete` AFTER DELETE ON `blog_likes` FOR EACH ROW BEGIN
  UPDATE blogs SET likeCount = likeCount - 1 WHERE blogId = OLD.blogId AND likeCount > 0;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table blog_comments
-- ----------------------------
DROP TRIGGER IF EXISTS `after_blog_comment_insert`;
delimiter ;;
CREATE TRIGGER `after_blog_comment_insert` AFTER INSERT ON `blog_comments` FOR EACH ROW BEGIN
  IF NEW.parentId IS NULL THEN
    UPDATE blogs SET commentCount = commentCount + 1 WHERE blogId = NEW.blogId;
  END IF;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table blog_comments
-- ----------------------------
DROP TRIGGER IF EXISTS `after_blog_comment_delete`;
delimiter ;;
CREATE TRIGGER `after_blog_comment_delete` AFTER UPDATE ON `blog_comments` FOR EACH ROW BEGIN
  IF NEW.status = 0 AND OLD.status = 1 AND OLD.parentId IS NULL THEN
    UPDATE blogs SET commentCount = commentCount - 1 WHERE blogId = OLD.blogId AND commentCount > 0;
  END IF;
END
;;
delimiter ; 