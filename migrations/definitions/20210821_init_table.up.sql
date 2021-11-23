-- GRANT CREATE,INSERT,SELECT,UPDATE,DELETE,DROP ON gin_study.* to 'docker'@'%';

CREATE DATABASE IF NOT EXISTS gin_study;

CREATE TABLE IF NOT EXISTS importances (
	id int PRIMARY KEY AUTO_INCREMENT COMMENT '重要度ID',
	name CHAR(255) NOT NULL COMMENT '重要度ラベル',
	level int NOT NULL COMMENT '重要度'
) COMMENT '重要度';

INSERT INTO importances
	(name, level)
	VALUES
	("MEDIUM", 1),
	("VERY_HIGH", 4),
	("HIGH", 3),
	("LOW", 2);

CREATE TABLE IF NOT EXISTS tasks (
  id CHAR(36) PRIMARY KEY COMMENT 'タスクID',
	name CHAR(255) NOT NULL COMMENT 'タスク名',
	importance_id int NOT NULL COMMENT '重要度ID',
	details VARCHAR(1000) COMMENT 'タスク詳細',
	registered_at datetime NOT NULL COMMENT '登録日時',
	deadline datetime NOT NULL COMMENT '期限日時',
	updated_at datetime NOT NULL COMMENT '更新日時',
	version INT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'バージョン',
	CONSTRAINT fk_importance FOREIGN KEY (importance_id) REFERENCES importances (id) ON UPDATE CASCADE
) COMMENT 'タスク';