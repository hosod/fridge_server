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
> git clone https://github.com/hosod/fridge_server.git

goの環境を作っている場合は多分下のコマンドを叩いてもローカルにダウンロードできます。こっちではGO_PATHの下に自動で展開してくれます。
>go get github.com/hosod/fridge_server

### 外部パッケージ
以下の外部パッケージに依存しています。dockerコンテナ上で動くのでサーバーを実行するだけならこれらがインストールされている必要はありません。
- github.com/gin-gonic/gin
- github.com/go-sql-driver/mysql
- github.com/jinzhu/gorm

## サーバーの操作

### サーバーの実行
`run_docker_compose_dev.sh`は開発用のコンテナを起動するスリプトファイルです。
サーバーを起動するときは基本的にこれを実行してください。
初回はビルドに時間がかかります。
`docker ps -a`でコンテナがうまく立ち上がったか確認できます。
MySQLのコンテナの起動と認証に時間がかかることがあります。
10秒くらい待ってください。\\
本番用のサーバーを実行するには以下のコマンドを実行してください。
> sh run_docker_compose.sh

\\
開発用と本番用のサーバーにはそれぞれ以下のような違いがあります。目的に応じて使い分けてください。
- 開発用のサーバー
    - 起動スクリプトの再実行で実装の変更が反映される
    - 開発していく時は基本的にこっちを使う
- 本番用のサーバー
    - イメージをビルドし直さないと実装の変更が反映されない
    - イメージとDBの初期化ファイルだけでコンテナを実行できる(デプロイが楽だけど今回はあまり関係ない)


### サーバーの停止
以下のコマンドを実行してください。`docker-compose.yml`と`docker-compose_dev.yml`を参考にしてコンテナを停止し、そのコンテナとネットワークを削除するところまでやってくれます。
> sh down_docker_compose.sh

## DB
データベースには以下のようなテーブルがあります。
各テーブルの定義などを見たい場合はDBのコンテナに入って直接確認してください。
<pre>
+--------------------+
| Tables_in_DB       |
+--------------------+
| contents           |
| food_genres        |
| food_types         |
| fridges            |
| user_follow_fridge |
| users              |
+--------------------+
</pre>


## web API

以下のURLはlocalで動いているサーバーにリクエストすることを前提にしています。
他の端末からリクエストする場合はサーバーを動かしている端末のIPを調べて、ホストの部分を書き換えてください。
- http://localhost:8000/users 
    - POST: データベースに登録
        - format: {"name":"hosod", "email":"hosoda@mail.com"}
- http://localhost:8000/users?uid={user_id}
    - GET: idで指定したレコードがjsonで返ってくる
        - response: {"id":1, "name":"hosod","email":"hosod@mail.com"}
    - PUT: idで指定したレコードを更新
        - format: {"name":"hosod", "email":"hosoda@mail.com"}
    - DELETE: idで指定したレコードを削除

- http://localhost:8000/fridges
    - POST: レコードを追加
        - format: {"name":"hosod's home"}
- http://localhost:8000/fridges?fid={fridge_id}
    - GET: idで指定したレコードが返ってくる
        - response: {"id":1, "name":"hosod's home"}
    - PUT: idで指定したレコードを更新
        - format: {"name":"hosod's home"}
    - DELETE: idで指定したレコードを削除
- http://localhost:8000/fridges/my-fridge?uid={user_id}
    - GET: user_idで指定したユーザーが所持している冷蔵庫を返す
        - response: {"id":1, "name":"hosod's home"}
- http://localhost:8000/fridges/follow-fridges?uid={user_id}
    - GET: user_idで指定したユーザーがフォローしている冷蔵庫を返す
        - response: {"id":1, "name":"hosod's home"}
- http://localhost:8000/contents?uid={user_id}
    - POST: 冷蔵庫に食品を追加。基本的にまとめて登録することを考えるので、リスト形式でまとめて渡す。ただしクエリパラメータでuser_idを指定した場合、そちらを優先してuserが所持している冷蔵庫に登録されます。
        - format: {"foods":[{"expiration_date":"2020/05/29", "quantity":1.5, "user_id":1,"food_type_id":1},...]}
- http://localhost:8000/contents?cid={content_id}
    - GET: idで指定したレコードがjsonで返ってくる
        - response: {"id":1, "name":"apple","expiration_date":"2020/07/28","quantity":2, "image":"http://localhost:8000/image/url", "genre":{"id":1, "name":"fruits","unit":"個"}}
    - PUT: idで指定したレコードの消費期限を更新
        - format: {"quantity":12345}
    - DELETE: idで指定したレコードを削除
- http://localhost:8000/contents/user?uid={user_id}
    - GET: 指定したuser_idのuserが所持している冷蔵庫に入っている食品のリスト
        - response: {"foods":[{"id":1, "name":"apple","expiration_date":"2020/07/28","quantity":2, "image":"http://localhost:8000/image/url", "genre":{"id":1, "name":"fruits","unit":"個"}},...]}
- http://localhost:8000/contents/fridge?fid={fridge_id}
    - GET: 指定したfridge_idの冷蔵庫に入っている食品のリスト
        - response: {"foods":[{"id":1, "name":"apple","expiration_date":"2020/07/28","quantity":2, "image":"http://localhost:8000/image/url", "genre":{"id":1, "name":"fruits","unit":"個"}},...]}
        
- http://localhost:8000/food_genres/list
    - GET: 食品の種類のリストを返します
        - response: [{"food_type_id":1,"name":"りんご","img_url":"http://localhost:8000/food_genres/imgs?iid=vege","genre":{"id":2,"name":"果物","unit":"個"}},...}]

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
    - DBにクエリ投げたり結果を返したり



