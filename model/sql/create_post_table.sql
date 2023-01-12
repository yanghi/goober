create table if not exists gb_post (
     id INT NOT NULL AUTO_INCREMENT,
     title VARCHAR(100) NOT NULL COMMENT '文章标题',
     content MEDIUMTEXT NOT NULL COMMENT '文章内容',
     description TEXT(0) NOT NULL COMMENT '文章简介',
     author_id INT(10) NOT NULL COMMENT '作者用户id',
     create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '发表时间',
     update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '更新时间',
     comments INT NOT NULL DEFAULT 0 COMMENT '文章评论数',
     view INT(10) NOT NULL DEFAULT 0 COMMENT '浏览量',
     tags JSON COMMENT '文章标签,json object集合,key为id',
     statu INT(2) NOT NULL DEFAULT 0 COMMENT '文章状态,0公共,1私密,2草稿',
     PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;