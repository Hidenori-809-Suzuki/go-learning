FROM golang:1.20.10-alpine3.17

# OSのインストール済みのパッケージをバージョンアップし、必要なパッケージをインストールします
RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev
# boiler-plateディレクトリの作成
RUN mkdir /go/src/boiler-plate
# ワーキングディレクトリの設定
WORKDIR /go/src/boiler-plate
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/boiler-plate
# パッケージのインポート
RUN go install golang.org/x/tools/cmd/goimports@latest
