CREATE TABLE metadata(  
    id int NOT NULL primary key  AUTO_INCREMENT  COMMENT 'Primary Key',
    create_time TIMESTAMP  DEFAULT CURRENT_TIMESTAMP  COMMENT 'Create Time',
    bulletin_name VARCHAR(128) NOT NULL DEFAULT '',
    bulletin_describe VARCHAR(128) NOT NULL DEFAULT ''
) default charset utf8mb4 COMMENT '';

insert into metadata (bulletin_name) values("sdfsd") ;


