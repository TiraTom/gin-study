# gin-study

Go 言語での勉強用 CRUD(TODO 管理)アプリ

## 構成案

- バックエンド：Go(Gin)
- フロントエンド：TypeScript(React)
- DB：MySQL

## TODO（勉強がてらやってみたいこと含む）

- [x] MySQL コンテナの立ち上げ（my.cnf 準備含む）
- [ ] API 設計(gRPC) <https://user-first.ikyu.co.jp/entry/2019/06/17/100000>あたりを参考にしてみる
- [ ] DB 設計
- [ ] ER 図自動生成の設定(SchemaSpy) <https://dev.classmethod.jp/articles/schemaspy-doc/>あたりを参考にしてみる
- [ ] READ 機能の実装
- [ ] CREATE 機能の実装
- [ ] UPDATE 機能の実装
- [ ] DELETE 機能の実装
- [ ] 検索機能の実装
- [ ] フロントエンドの実装（別レポジトリ）
- [ ] Lint か何かで英単語と日本語の間にスペースが入るようになってしまってるので直す
- [ ] アプリも Docker コンテナとして動かす（デバッグもできるようにする）
- [ ] AWS にデプロイして動かす（CloudFormation 利用）
- [ ] ローカルで JMeter で負荷をかけてみる
- [ ] Cognito で認証機能をつける（基本知識をまずつけるところから）
