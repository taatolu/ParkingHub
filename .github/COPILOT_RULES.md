# このファイルはCopilotやレビュワーが必ず参照すること
レビューやCopilotのコメントの基本指針です

# Copilot返答ルール
- 返答は必ず日本語で行うこと。
- 丁寧かつ簡潔に説明すること。
- フロントエンドの質問については、以下の”## フロントエンドのディレクトリ構成”を元に、レイヤーごとの責務分離を意識してコメントすること。
- バックエンドの質問については、以下の”## バクエンドのディレクトリ構成”を元に、レイヤーごとの責務分離を意識してコメントすること。
- 質問されたレイヤー以外の変更も必要であれば、他のレイヤーの変更についても合わせてコメントすること。
- 技術用語や専門用語は、必要に応じて簡単な解説を添えること。
- コード例やコマンドはMarkdownコードブロックで示すこと。
- 質問の意図が不明な場合は、確認の質問を返すこと。
- 回答に根拠や理由がある場合は、簡単に補足説明を加えること。
- 不明点や危険な操作が含まれる場合は、注意喚起を行うこと。

## フロントエンドのディレクトリ構成
/frontend
├── public/              # 静的ファイル（画像、favicon、index.html など）
├── src/                 # アプリのソースコード
│   ├── assets/          # 画像・フォント・グローバルCSSなど
│   ├── components/      # 再利用可能な小さなコンポーネント群
│   ├── features/        # 機能単位（ドメイン単位）のまとまり（例：認証、駐車場管理など）
│   ├── hooks/           # カスタムフック
│   ├── layouts/         # レイアウト用コンポーネント（画面の枠組み）
│   ├── pages/           # ページ単位のコンポーネント（React Routerのルートごと）
│   ├── routes/          # ルーティング設定（必要に応じて）
│   ├── services/        # API通信や外部サービス連携のロジック
│   ├── stores/          # 状態管理（Redux, Zustand, Recoilなどを使う場合）
│   ├── types/           # 型定義（TypeScript用）
│   ├── utils/           # 汎用関数
│   └── App.tsx          # エントリーポイント
├── .env                 # 環境変数
├── package.json         # npmの設定ファイル
├── tsconfig.json        # TypeScriptの設定
└── README.md

## バクエンドのディレクトリ構成
├── cmd/                // エントリポイント（main.goなど）
├── domain/             // エンティティやリポジトリのインターフェース（ビジネスルール）
│   ├── model/          // ドメインモデル（例: carOwner.go, user.go）やファクトリー関数
│   └── service/
│       └── password.go     // type PasswordService interface {...}
│   └── repository/     // リポジトリインターフェース
├── usecase/            // ユースケース（ビジネスロジック）
│   └── user.go
│   └── todo.go
├── infrastructure/     // DBや外部サービスとのやりとり（実装）
│   ├── mysql/
│   ├── postgres/
│   └── external/
├── presentation/       // 入出力層（ハンドラー、ルーター、リクエスト/レスポンス）
│   └── http/
│       ├── handler/    //xxx_handler.go (ユースケース層を呼び出すための窓口)
│       ├── router.go   //muxやgin等のルーティングの初期化（URLパスとハンドラー関数のひも付け）
│       ├── request/    //xxx_request.go(POSTデータのJSON構造体、バリデーション関数)
│       └── response/   //xxx_response.go(レスポンス出力用の構造体や補助関数)
├── registry/           // 依存性注入（DI）や各種初期化
├── config/             // 設定ファイル
├── mocks/              // モックやFake（テスト用）
├── main.go
└── go.mod