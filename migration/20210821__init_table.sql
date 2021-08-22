CREATE DATABASE gin_study;

CREATE TABLE gin_study.importances (
	id int PRIMARY KEY COMMENT '重要度ID',
	name CHAR(255) NOT NULL COMMENT '重要度ラベル',
	level int NOT NULL COMMENT '重要度'
) COMMENT '重要度';

CREATE TABLE gin_study.tasks (
  id CHAR(32) PRIMARY KEY COMMENT 'タスクID',
	name CHAR(255) NOT NULL COMMENT 'タスク名',
	importance_id int NOT NULL COMMENT '重要度ID',
	details VARCHAR(1000) COMMENT 'タスク詳細',
	registered_at datetime NOT NULL COMMENT '登録日時',
	deadline datetime NOT NULL COMMENT '期限日時',
	updated_at datetime NOT NULL COMMENT '更新日時',
	CONSTRAINT fk_importance FOREIGN KEY (importance_id) REFERENCES importances (id) ON UPDATE CASCADE
) COMMENT 'タスク';