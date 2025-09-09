FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /go/src/app

# go.modとgo.sumを最初にコピー（キャッシュ効率化）
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# 依存関係を再整理
RUN go mod tidy

# ビルドして動作確認
RUN go build -o main .

# アプリケーション実行
CMD ["go","run","."]