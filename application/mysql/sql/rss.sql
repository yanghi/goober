create table if not exists gb_rss_feed (
    id INT NOT NULL AUTO_INCREMENT,
    feed_link VARCHAR(400) UNIQUE COMMENT 'feed link',
    title VARCHAR(100) NOT NULL COMMENT '标题',
    description VARCHAR(500) COMMENT '描述',
    link VARCHAR(400) NOT NULL COMMENT '站点链接',
    authors JSON COMMENT '[]{name,email}数组',
    published VARCHAR(40) NOT NULL COMMENT '发表时间',
    updated VARCHAR(40) COMMENT '更新时间',
    version VARCHAR(5) COMMENT 'version',
    feed_type VARCHAR(4) COMMENT 'feed type',
    language VARCHAR(8) COMMENT 'language',
    defaut_folder_name VARCHAR(20) COMMENT 'folder name',
    rsshub BOOL COMMENT 'is RSShub',
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

create table if not exists gb_rss_feed_items (
    id INT NOT NULL AUTO_INCREMENT,
    guid VARCHAR(255) UNIQUE NOT NULL COMMENT 'guid',
    feed_id INT NOT NULL COMMENT 'feed id',
    title VARCHAR(300) NOT NULL COMMENT '标题',
    description MEDIUMTEXT,
    content MEDIUMTEXT,
    link VARCHAR(300),
    published VARCHAR(40) NOT NULL COMMENT '发表时间',
    updated VARCHAR(40) COMMENT 'updated time',
    authors JSON COMMENT 'name,email数组',
    categories JSON COMMENT '分类或标签',
    rsshub BOOL COMMENT '是否为rsshub源',
    cretor_id INT NOT NULL DEFAULT 0 COMMENT '添加者的用户id',
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists gb_rss_user_folder (
    id INT NOT NULL AUTO_INCREMENT,
    uid INT NOT NULL,
    name VARCHAR(30) NOT NULL COMMENT '名称',
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists gb_rss_user_feed (
    id INT NOT NULL AUTO_INCREMENT,
    uid INT NOT NULL,
    feed_id INT UNIQUE NOT NULL,
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '用户添加的时间',
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

create table if not exists gb_rss_user_feed_item (
    id INT NOT NULL AUTO_INCREMENT,
    uid INT NOT NULL,
    item_id INT UNIQUE NOT NULL,
    read BOOL DEFAULT false COMMENT '阅读状态',
    favorite BOOL DEFAULT false COMMENT '收藏',
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '用户添加的时间',
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;