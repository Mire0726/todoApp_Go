# goのバージョン指定
FROM golang:1.19.2-alpine3.16

# アプリケーションのワーキングディレクトリを設定
WORKDIR /app

# go.mod と go.sum ファイルをコピー
COPY ./app/go.mod ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY ./app .

# アプリケーションをビルド
RUN go build -o main .

# アプリケーションがリッスンするポートを指定
EXPOSE 8080

# main.go実行
CMD ["./main"]

