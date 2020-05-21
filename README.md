# fridge_server
情報システム論実習Group1で作るアプリのサーバーです。
最終的には以下のような機能を実装していきたい。

- 冷蔵庫の中身一覧
    - 食材の情報(名前、数量、期限)
    - 食材追加
    - 食材消費
    - 分類
    - ソート(賞味期限順など)
- 他の人との共有(やりたい)
- レシピを推薦(できれば)
## インストール
`サーバーを動かすのにdockerを使っています。まずは各自でdocker環境の構築を行なってください。`

### ローカルに持ってくる
以下のコマンドを実行してリポジトリをローカルの好きな場所にクローンします。
ローカルにgoの環境を作っている場合は`$GO_PATH/src/github.com/hosod/`を作ってその下にcloneしてあげないと、開発の時にプロジェクト内のパッケージを参照できなくなります。注意してください。
> git clone https://github.com/hosod/fridge_server.git\

\
goの環境を作っている場合は多分下のコマンドを叩いてもローカルにダウンロードできます。こっちではGO_PATHの下に自動で展開してくれます。
>go get github.com/hosod/fridge_server

### 外部パッケージ
以下の外部パッケージに依存しています。サーバーを実行する分にはこれらがインストールされている必要はありませんが、ローカルで開発を行う時にエディタの自動補完等を効かせたい場合は`go get`で以下のパッケージをインストールしてください。
- github.com/gin-gonic/gin
- github.com/go-sql-driver/mysql
- github.com/jinzhu/gorm

## サーバーの操作
詳しくないのでわかりませんが以下の操作ではbashスクリプトを使っているのでWindowでは実行できません。
直接`-f`オプションを指定してdocker-composeしてあげてください。
### サーバーの実行
リポジトリの一番上に移動して以下のコマンドを実行してください。開発用のサーバーが実行されます。初回はビルドに時間がかかります。
`docker ps -a`でコンテナがうまく立ち上がったか確認できます。うまくいかなかったらもう一度スクリプトを実行してみてください。
> sh run_docker_compose_dev.sh

本番用のサーバーを実行するには以下のコマンドを実行してください。
> sh run_docker_compose.sh

\\
開発用と本番用のサーバーにはそれぞれ以下のような違いがあります。目的に応じて使い分けてください。
- 開発用のサーバー
    - 起動スクリプトの再実行で実装の変更が反映される
    - 開発していく時は基本的にこっちを使う
- 本番用のサーバー
    - イメージの再ビルドをしないと実装の変更が反映されない
    - イメージとDBの初期化ファイルだけでコンテナを実行できる
    - たぶん動かすだけならこっちの方が楽


### サーバーの停止
以下のコマンドを実行してください。`docker-compose.yml`と`docker-compose_dev.yml`を参考にしてコンテナを停止し、そのコンテナとネットワークを削除するところまでやってくれます。
> sh down_docker_compose.sh

## DB
### usersテーブル
<pre>
+-------+------------------+------+-----+---------+----------------+
| Field | Type             | Null | Key | Default | Extra          |
+-------+------------------+------+-----+---------+----------------+
| id    | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| name  | varchar(255)     | YES  |     | NULL    |                |
| email | varchar(255)     | YES  |     | NULL    |                |
+-------+------------------+------+-----+---------+----------------+
</pre>
### fridgesテーブル
<pre>
+-------+--------------+------+-----+---------+----------------+
| Field | Type         | Null | Key | Default | Extra          |
+-------+--------------+------+-----+---------+----------------+
| id    | int(11)      | NO   | PRI | NULL    | auto_increment |
| name  | varchar(255) | NO   |     | NULL    |                |
+-------+--------------+------+-----+---------+----------------+
</pre>

## web API

以下のURLはlocalで動いているサーバーにリクエストすることを前提にしています。
他の端末からリクエストする場合はサーバーを動かしている端末のIPを調べて、`localhost`の部分を書き換えてください。
- http://localhost:8000/users 
    - GET: usersテーブルの全レコードがjson形式で返ってくる
    - POST: データベースに登録
- http://localhost:8000/users/:id
    - GET: idで指定したレコードがjsonで返ってくる
    - PUT: idで指定したレコードを更新
    - DELETE: idで指定したレコードを削除
`/usersへのGETで返ってくるjsonの例`
>[
>    {
>        "id":1,
>        "name":"Yamada",
>        "email":"yamada@mail.com"
>    },
>    {
>        "id":2,
>        "name":"Tanaka",
>        "email":"tanaka@mail.com"
>    },...
>]
- http://localhost:8000/fridges
    - GET: fridgesテーブルの全レコードが帰ってくる
    - POST: レコードを追加
- http://localhost:8000/fridges/:id
    - GET: idで指定したレコードが返ってくる
    - POST: idで指定したレコードを更新
    - DELETE: odで指定したレコードを削除

## アプリケーションサーバーでの処理
- main
    - serverやdatabase の初期化を行う。
- server
    - routingやらなんやら
- database
    - mysqlとの接続や通信を担う
- entity
    - ORMはdatabaseのレコードを構造体に落とし込んでくれる。その構造体の定義をしてる
- service
    - サービスロジックとか。



