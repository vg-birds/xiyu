CREATE TABLE `i_blog` (
  `id` int(10) NOT NULL auto_increment,
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT  '标题',
  `content` text  COMMENT '内容',
  `author` int(10) NOT NULL DEFAULT 0 COMMENT '作者',
  `viewNum` int (10) NOT NULL DEFAULT 0 COMMENT '浏览量',
  `publishTime` int(10) NOT NULL DEFAULT 0 COMMENT '发布时间',
  `deleted` int(10) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `updateTime` int(10) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY(id),
  KEY `id` (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='博客';

CREATE TABLE `i_comment`(
  `id` int(10) NOT NULL auto_increment ,
  `content` varchar (50) NOT NULL DEFAULT '' COMMENT '评论内容',
  `toUid` int(10) NOT NULL  DEFAULT 0 COMMENT '被评论人',
  `fromUid` int (10) NOT NULL DEFAULT 0 COMMENT '评论人',
  `publishTime` int (10) NOT NULL DEFAULT 0 COMMENT '发布时间',
  `deleted` int (10) NOT NULL  DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (id)
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

CREATE TABLE `i_liked`(
  `id` int(10) NOT NULL auto_increment ,
  `bid` int(10) NOT NULL  DEFAULT 0 COMMENT '文章id',
  `fromUid` int (10) NOT NULL DEFAULT 0 COMMENT '评论人',
  `publishTime` int (10) NOT NULL DEFAULT 0 COMMENT '发布时间',
  `deleted` int (10) NOT NULL  DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (id)
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='评论表';