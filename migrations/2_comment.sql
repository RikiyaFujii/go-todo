-- +migrate Up
CREATE TABLE `comments` (
  `comment_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `body`  varchar(256) NOT NULL COMMENT 'コメント',
  `todo_id`  int(11) NOT NULL COMMENT 'TODO_ID',
  PRIMARY KEY (`comment_id`)
);

-- +migrate Down
DROP TABLE comments;