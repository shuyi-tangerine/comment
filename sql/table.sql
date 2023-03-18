
-- TODO 需要考虑一下二级评论、多级评论怎么弄
CREATE TABLE `comment` (
                           `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                           `comment_id` BIGINT NOT NULL DEFAULT 0 COMMENT '评论ID',
                           `user_id` BIGINT NOT NULL DEFAULT 0 COMMENT '用户ID',
                           `group_id` BIGINT NOT NULL DEFAULT 0 COMMENT '所属分组ID，即在哪个地方评论，比如视频、文章等',
                           `app_id` INT NOT NULL DEFAULT 0 COMMENT '应用ID',
                           `text` VARCHAR(256) DEFAULT NULL COMMENT '文本',
                           `images` JSON DEFAULT NULL COMMENT '图片',
                           `reply_comment_id` BIGINT NOT NULL DEFAULT 0 COMMENT '回复的评论',
                           `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，也就是通知发送时间',
                           `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
                           `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态，1-正常，2-已删除',
                           `extra` JSON DEFAULT NULL COMMENT '一些额外的信息',
                           PRIMARY KEY(`id`),
                           UNIQUE KEY `u_cid` (`comment_id`)
)  ENGINE = InnoDB CHARSET = utf8mb4 COLLATE = utf8mb4_bin AUTO_INCREMENT = 2023 COMMENT '评论';
