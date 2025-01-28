# ビルドステージ
FROM golang:1.23 AS builder

WORKDIR /app

# モジュールファイルをコピーして依存関係をダウンロード
COPY go.mod ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY main.go ./
RUN go build -o main .

# ランタイムステージ
FROM ubuntu:24.04

WORKDIR /root/

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

# 実行権限を付与
RUN chmod +x ./main

ENTRYPOINT ["./main"]
