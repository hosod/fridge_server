FROM golang:latest

# RUN mkdir /go/src/sample
# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u github.com/wakashiyo/compose-sample/users
# RUN go get -u github.com/go-sql-driver/mysql
# RUN go get -u github.com/go-xorm/xorm

# COPY main.go /go/src/sample

# CMD ["go", "run", "/go/src/sample/main.go"]

#################################################################################################################

## 存在しない場合は、ディレクトリ作成される
WORKDIR /go/src/github.com/hosod/fridge_server/app
# RUN mkdir -p /go/src/github.com/hosod/fridge_server/app

# WORKDIRからの相対パスの位置に配置する
COPY . /go/src/github.com/hosod/fridge_server

WORKDIR /go/src/github.com/hosod/fridge_server/app

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
# RUN go get -u github.com/go-xorm/xorm

# go install で /go/bin 配下にバイナリがビルドされる
# RUN go install github.com/hosod/fridge_server/app

# image自体に最初からPATHに/go/binのパスが設定されているため、バイナリファイルを書くだけで動かすことができる
# ENTRYPOINT [ "app" ]