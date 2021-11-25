# gin-study

Go 言語での勉強用 CRUD(TODO 管理)アプリ  
gin-study ってレポジトリ名だけど gRPC と DB 接続周りの勉強用って感じ

## 構成案

- バックエンド：Go(Gin)
- フロントエンド：TypeScript(React)
- DB：MySQL

## TODO（勉強がてらやってみたいこと含む）

- [x] MySQL コンテナの立ち上げ（my.cnf 準備含む）
- [x] API 設計(gRPC) <https://user-first.ikyu.co.jp/entry/2019/06/17/100000>あたりを参考にしてみる
- [x] ダミーのメソッドでサーバー起動（CatServer）
- [x] 実際の API の定義に従って受け口だけ用意
- [x] リクエストのログ出力設定
- [x] interceptor 以外でも zap でログが出せるようにする
- [x] DI 設定　参考：<https://christina04.hatenablog.com/entry/google-wire>
- [x] 環境変数の読み込み設定
- [x] DB 設定
- [x] GORM 設定
- [x] フォーマットとか lint 系の設定（golangci-lint）
- [x] READ 機能の実装（全取得）
- [x] CREATE 機能の実装
- [x] バリデーション実装する（まずは Create 分）
- [x] gorm の time の警告対応：<https://github.com/go-sql-driver/mysql#timetime-support>に書いてある
- [x] UPDATE 機能の実装（バリデーション含む）
- [x] DELETE 機能の実装（バリデーション含む）
- [x] ID によるタスク取得機能の実装
- [x] タスク作成時の RegisterdAt, updatedAt の時刻が ms まで入ってる。秒だけで十分なので直す。
- [x] タスク名での検索
- [x] 重要度での検索
- [x] 検索結果の戻り値に ImportanceName がないので修正
- [x] 検索条件に期限日時が使えるようにする
- [x] 検索機能の全般的な動作確認
- [x] ~~Lint か何かで英単語と日本語の間にスペースが入るようになってしまってるので直す~~ →Prettier のデフォルト設定だしで一旦スルーすることにする
- [x] 例外処理（grpc_recovery のミドルウェア使いつつ）
- [x] マイグレーションツールの設定
- [x] ~~環境に応じてダミーデータを投入したりしなかったりできるようにする~~ → ダミーデータは make ファイルで適宜入れさせるようにした
- [x] ダミーデータを複数行書いても問題なく投入できるようにする
- [ ] DB 定義変更を試してみる
- [ ] ER 図自動生成の設定(SchemaSpy) <https://dev.classmethod.jp/articles/schemaspy-doc/>あたりを参考にしてみる
- [ ] GitHubAction で push 後自動でテスト実施
- [ ] アプリも Docker コンテナとして動かす（デバッグもできるようにする）
- [ ] ~~AWS にデプロイして動かす（CloudFormation 等利用）~~ 別でやる
- [ ] ローカルで JMeter で負荷をかけてみる
- [ ] Cognito で認証機能をつける（基本知識をまずつけるところから）

# 起動メモ

- 起動時引数でどの環境用の環境変数ファイルを読み込むか指定する（ex. ENV=local とすると.env.local が読み込まれる）
