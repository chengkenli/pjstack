
-- chengken.sr_slow_query_manager definition

CREATE TABLE `sr_slow_query_manager` (
                                         `app` varchar(100) NOT NULL COMMENT '集群名称(英文)',
                                         `feip` varchar(200) NOT NULL COMMENT '集群连接地址(必填)F5,VIP,CLB,FE',
                                         `user` varchar(200) NOT NULL COMMENT '集群登录账号(必填) 建议是管理员角色的账号',
                                         `password` varchar(500) NOT NULL COMMENT '集群登录密码(必填)',
                                         `feport` int(11) NOT NULL DEFAULT '9030' COMMENT '集群登录端口，默认9030',
                                         `address` varchar(500) DEFAULT NULL COMMENT 'MANAGER地址，如果填了MANAGER地址，那么将触发定时检查LICENSE是否过期(企业级)',
                                         `expire` int(11) DEFAULT '30' COMMENT 'LICENSE是否过期(企业级)过期提醒倒计时，单位day',
                                         `status` int(11) NOT NULL DEFAULT '0' COMMENT 'LICENSE是否过期(企业级)开关,0 off, 1 on',
                                         `fe_log_path` varchar(500) NOT NULL COMMENT 'FE 日志目录',
                                         `be_log_path` varchar(500) NOT NULL COMMENT 'BE 日志目录',
                                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='StarRocks登录配置，manager地址,(定期检查license过期日期)';