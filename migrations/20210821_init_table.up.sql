-- 以下のDB用意のクエリはマイグレーション対象外（マイグレーション用のDB接続でDB not existsなどのエラーになることを防ぐため）
-- CREATE DATABASE IF NOT EXISTS gin_study;  ※DB名は適宜変更
-- GRANT CREATE,INSERT,SELECT,UPDATE,DELETE,DROP,REFERENCES,ALTER ON gin_study.* to 'docker'@'%';  ※DB名は適宜変更

CREATE TABLE IF NOT EXISTS importances (
	id int PRIMARY KEY AUTO_INCREMENT COMMENT '重要度ID',
	name CHAR(255) NOT NULL COMMENT '重要度ラベル',
	level int NOT NULL COMMENT '重要度'
) COMMENT '重要度';

INSERT INTO importances
	(id, name, level)
	VALUES
	(1, "MEDIUM", 2),
	(2, "VERY_HIGH", 4),
	(3, "HIGH", 3),
	(4, "LOW", 1);

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