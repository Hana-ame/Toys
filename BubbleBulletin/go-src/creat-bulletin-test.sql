CREATE TABLE bubbles(  
    id int NOT NULL primary key  AUTO_INCREMENT  COMMENT 'Primary Key',
    create_time TIMESTAMP  DEFAULT CURRENT_TIMESTAMP  COMMENT 'Create Time',
    board_name VARCHAR(48) NOT NULL DEFAULT 'test',

    meta_author varchar(24) NOT NULL DEFAULT '',
    meta_title  VARCHAR(48) NOT NULL DEFAULT '',
    meta_imgsrc VARCHAR(4096) NOT NULL DEFAULT '',

    content TEXT NOT NULL DEFAULT '',
    background VARCHAR(128) NOT NULL DEFAULT '',

    pos_x  int(11) NOT NULL DEFAULT 0,
    pos_y  int(11) NOT NULL DEFAULT 0,
    size_x int(11) NOT NULL DEFAULT 0,
    size_y int(11) NOT NULL DEFAULT 0,

    scale FLOAT NOT NULL DEFAULT 0
    
) default charset utf8mb4 COMMENT '';