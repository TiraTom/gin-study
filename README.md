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
- [x] ER 図自動生成の設定(SchemaSpy) <https://dev.classmethod.jp/articles/schemaspy-doc/>あたりを参考にしてみる → schemaspyUser(pass: schemaspyPass)は SELECT 権限のみ付与、mysql_native_password 方式で認証させるようにして useSSL=false・allowPublicKeyRetrieval=true の設定を追加（ローカル実行だし）、schemaspy のイメージに用意されている JDBC ドライバは MySQL6 系なので 8 用の jar ファイルを落としてきてマウントして使わせる　等の作業がいろいろ必要だった・・・。
- [x] DB 定義変更を試してみる
- [x] DB 周辺のテスト記述
- [x] toDTO()は presentation 部分で呼び出した方が、grpc でない時にも presentation 層だけで対応できるしいい気がする
- [x] GitHubAction で push 後自動でテスト実施
- [x] .env.test ファイルをプロジェクトルートに移動（repository_impl のテストでの env ファイル読み込みでエラーが起きないようにパス設定をあれこれ直す必要あり）
- [ ] presentation のテスト記述：
- [ ] infrastructure のテストで isTest フラグを使っている箇所を除去する。設定必須の環境変数 ENV の値で管理して分岐させることにして、通常コードがテストのことを意識しなくて済むようにする
- [ ] usecase のテスト記述　<https://zenn.dev/sanpo_shiho/articles/01da627ead98f5>あたり参考にしてみる
- [ ] usecase のテストは DB 処理まで実際に行わせるので、infra 層のテストと並列実行にならないようにしないとテスト失敗するかも？
- [ ] アプリも Docker コンテナとして動かす（デバッグもできるようにする）
- [x] ~~AWS にデプロイして動かす（CloudFormation 等利用）~~ 別でやる
- [ ] ローカルで JMeter で負荷をかけてみる
- [x] ~~Cognito で認証機能をつける（基本知識をまずつけるところから）~~　 CloudFormation と合わせて別で実施

# 起動メモ

- 起動時引数でどの環境用の環境変数ファイルを読み込むか指定する（ex. ENV=local とすると.env.local が読み込まれる）
