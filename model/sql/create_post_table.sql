create table if not exists gb_post (
     id INT NOT NULL AUTO_INCREMENT,
     title TEXT(0) NOT NULL COMMENT '文章标题',
     content TEXT(0) NOT NULL COMMENT '文章内容',
     description TEXT(0) NOT NULL COMMENT '文章简介',
     author_id INT(10) NOT NULL COMMENT '作者用户id',
     create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '发表时间',
     comments INT NOT NULL DEFAULT '0' COMMENT '文章评论数',
     view INT(10) NOT NULL DEFAULT '0' COMMENT '浏览量',
     tags JSON COMMENT '文章标签,json object集合,key为id',
     PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;