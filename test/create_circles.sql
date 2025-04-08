-- 在数据库中创建常见的编程圈子
USE jxust_softhub;

-- 清除测试数据（如果需要）
-- DELETE FROM follows WHERE followType = 2;
-- DELETE FROM circles WHERE circleId > 1;

-- 创建圈子数据
INSERT INTO circles 
(circleName, description, icon, userId, memberCount, postCount, status, isOfficial, createTime, updateTime)
VALUES
('Java技术圈', 'Java编程语言学习与分享，包括Spring、SpringBoot等框架讨论', '/uploads/circle/java.png', 1, 0, 0, 1, 1, NOW(), NOW()),
('Python爱好者', 'Python编程、机器学习、数据分析和爬虫技术交流', '/uploads/circle/python.png', 1, 0, 0, 1, 1, NOW(), NOW()),
('前端开发', 'HTML、CSS、JavaScript以及各种前端框架和库的讨论', '/uploads/circle/frontend.png', 1, 0, 0, 1, 1, NOW(), NOW()),
('算法与数据结构', '算法设计、数据结构、编程竞赛和面试题解析', '/uploads/circle/algorithm.png', 1, 0, 0, 1, 1, NOW(), NOW()),
('移动开发', 'Android、iOS开发技术分享和交流', '/uploads/circle/mobile.png', 1, 0, 0, 1, 1, NOW(), NOW()),
('数据库技术', 'MySQL、PostgreSQL、MongoDB等数据库技术学习与经验分享', '/uploads/circle/database.png', 1, 0, 0, 1, 1, NOW(), NOW()),
('云计算与DevOps', '云服务、容器技术、CI/CD和自动化运维讨论', '/uploads/circle/cloud.png', 1, 0, 0, 1, 1, NOW(), NOW()),
('网络安全', '网络安全知识、漏洞分析、渗透测试和安全防护', '/uploads/circle/security.png', 1, 0, 0, 1, 1, NOW(), NOW());

-- 查询结果
SELECT circleId, circleName, memberCount FROM circles ORDER BY circleId; 