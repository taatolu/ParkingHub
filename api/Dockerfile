FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /go/src/app

# カレントディレクトリの内容をコンテナにコピー
COPY . .

# 必要なGoモジュールをインストール
RUN go mod tidy
RUN go mod download

# GORMをインストール
RUN go get gorm.io/gorm@latest && \
    go get gorm.io/driver/postgres@latest

# 本番用なら go build でバイナリを作成しCMDで実行
# 開発運用の場合は go run でOK
CMD ["go", "run", "main.go"]