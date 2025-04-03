-- 初始化系统数据
USE jxust_softhub;

-- 清空已有数据
DELETE FROM users WHERE userId > 0;
DELETE FROM roles WHERE roleId > 0;

-- 初始化角色数据
INSERT INTO roles (roleId, roleName, description) VALUES 
(1, '超级管理员', '具有系统最高权限，可以管理所有功能和用户'),
(2, '内容管理员', '负责管理网站内容，包括新闻、活动等'),
(3, '会长', '协会会长，具有管理协会成员和内容的权限'),
(4, '普通用户', '普通注册用户，仅有基本浏览和参与功能');

-- 初始化超级管理员账号
-- 密码: admin123 (已用MD5加密)
INSERT INTO users (userId, userName, password, name, roleId, email, phone) VALUES
(1, 'admin', '0192023a7bbd73250516f069df18b500', '超级管理员', 1, 'admin@example.com', '13800000000'); 