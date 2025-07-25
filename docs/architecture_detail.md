# 詳細設計書（駐車場管理）

---

## **1. 画面詳細設計**

## **1.1 ログイン画面**

| **項目** | **内容** |
| --- | --- |
| 入力項目 | ユーザーID, パスワード |
| ボタン | ログイン |
| バリデーション | 必須入力、ID/パスワード誤り時エラー表示 |
| 遷移先 | ダッシュボード（認証成功時） |
| 備考 | セッション管理、パスワードはハッシュ化 |

## **1.2 ダッシュボード画面**

| **項目** | **内容** |
| --- | --- |
| 表示内容 | 利用者名、権限、期限切れアラート、違反アラート |
| 機能 | 各管理画面へのリンク、アラート詳細への遷移 |
| 備考 | アラート：免許証/保険/車検の期限、違反累積数超過等 |

## **1.3 社員・業者管理画面**

| **項目** | **内容** |
| --- | --- |
| 表示項目 | ユーザーID、氏名、所属/業者名、連絡先、権限区分、ログインID |
| 検索 | 氏名・所属・業者名・ログインIDで部分一致検索 |
| 詳細表示 | 利用者情報＋関連車両所有者・車両・免許証情報 |
| 編集・削除 | 管理者のみ可能 |
| 新規登録 | 必須項目：氏名、所属/業者名、連絡先、権限区分、ログインID/パスワード |
| バリデーション | 氏名・連絡先必須、連絡先はメール形式チェック、ログインID重複不可 |

## **1.4 車両管理画面**

| **項目** | **内容** |
| --- | --- |
| 表示項目 | 車両ID、車両所有者ID、車両ナンバー、車種、任意保険期限、車検期限 |
| 検索 | 車両ナンバー・車両所有者名で検索 |
| 詳細表示 | 車両情報＋関連車両所有者 |
| 編集・削除 | 管理者のみ可能 |
| 新規登録 | 必須項目：車両ナンバー、車両所有者ID、車種 |
| バリデーション | ナンバー必須、期限は日付形式、所有者ID存在チェック |

## **1.5 免許証管理（利用者詳細内）**

| **項目** | **内容** |
| --- | --- |
| 表示項目 | 有効期限、車両所有者ID |
| 編集 | 管理者のみ可能 |
| バリデーション | 有効期限は未来日、所有者ID存在チェック |

## **1.6 違反記録管理画面**

| **項目** | **内容** |
| --- | --- |
| 表示項目 | 違反ID、違反種別、発生日、車両所有者ID、車両ID、始末書要否 |
| 検索 | 車両所有者名・車両ナンバー・違反種別・期間 |
| 詳細表示 | 違反内容詳細、累積違反数、関連車両所有者情報 |
| 編集・削除 | 管理者のみ可能 |
| 新規登録 | 必須項目：車両所有者ID、車両ID、違反種別、発生日 |
| バリデーション | 違反種別必須、発生日は日付形式、所有者ID・車両ID存在チェック |
| 始末書判定 | 10月～翌年9月の期間で累計5回違反で自動フラグON |

## **1.7 許可証発行画面**

| **項目** | **内容** |
| --- | --- |
| 発行対象 | 車両所有者選択（社員/協力業者/外注業者） |
| 許可証種別 | 従業員駐車場・工場内・外注業者向け |
| 発行機能 | PDF出力・印刷 |
| 有効期限 | 発行日から1年間（デフォルト） |
| バリデーション | 車両所有者・車両情報が揃っていること |

## **1.8 注意文書作成画面**

| **項目** | **内容** |
| --- | --- |
| 入力 | 利用者選択、違反内容選択 |
| 出力 | メール文書・貼紙文書の自動生成 |
| テンプレート | 違反種別ごとに文面テンプレート設定 |
| 配布方法 | 社内メール or 印刷（靴箱掲示用） |

## **1.9 CSVインポート画面**

| **項目** | **内容** |
| --- | --- |
| 対象 | 利用者・車両・違反記録データ |
| ファイル形式 | UTF-8, ヘッダ行必須 |
| バリデーション | 必須項目チェック、データ型チェック |
| 結果表示 | 成功/失敗件数、エラー内容一覧 |

## **2. API設計（主要エンドポイント例）**

| **メソッド** | **パス** | **概要** | **認可** |
| --- | --- | --- | --- |
| POST | /api/login | ログイン | 公開 |
| GET | /api/users | 利用者一覧取得 | 管理者 |
| POST | /api/users | 利用者新規登録 | 管理者 |
| PUT | /api/users/{id} | 利用者情報更新 | 管理者 |
| DELETE | /api/users/{id} | 利用者削除 | 管理者 |
| GET | /api/car_owners | 車両所有者一覧取得 | 管理者 |
| POST | /api/car_owners | 車両所有者新規登録 | 管理者 |
| PUT | /api/car_owners/{id} | 車両所有者情報更新 | 管理者 |
| DELETE | /api/car_owners/{id} | 車両所有者削除 | 管理者 |
| GET | /api/vehicles | 車両一覧取得 | 管理者 |
| POST | /api/vehicles | 車両新規登録 | 管理者 |
| PUT | /api/vehicles/{id} | 車両情報更新 | 管理者 |
| DELETE | /api/vehicles/{id} | 車両削除 | 管理者 |
| GET | /api/violations | 違反記録一覧取得 | 管理者 |
| POST | /api/violations | 違反記録新規登録 | 管理者 |
| PUT | /api/violations/{id} | 違反記録更新 | 管理者 |
| DELETE | /api/violations/{id} | 違反記録削除 | 管理者 |
| GET | /api/permits | 許可証一覧取得 | 管理者 |
| POST | /api/permits | 許可証発行 | 管理者 |
| POST | /api/import/csv | CSVインポート | 管理者 |



## **3. データベース詳細設計（テーブル定義例）**

## **3.1 users（利用者）**

| **カラム名** | **型** | **制約** | **説明** |
| --- | --- | --- | --- |
| id | SERIAL | PK | 利用者ID（User） |
| name | VARCHAR | NOT NULL | 氏名 |
| company | VARCHAR |  | 所属/業者名 |
| email | VARCHAR | NOT NULL | 連絡先 |
| role | VARCHAR | NOT NULL | 権限区分 |
| login_id | VARCHAR | NOT NULL | ログインID |
| password_hash | VARCHAR | NOT NULL | パスワードHash |
| created_at | TIMESTAMP | NOT NULL | 登録日時 |
| updated_at | TIMESTAMP | NOT NULL | 更新日時 |

## **3.2 vehicles（車両）**

| **カラム名** | **型** | **制約** | **説明** |
| --- | --- | --- | --- |
| id | SERIAL | PK | 車両ID |
| car_owner_id | INTEGER | FK(car_owners.id) | 車両所有者ID |
| number | VARCHAR | NOT NULL | ナンバー |
| type | VARCHAR |  | 車種 |
| insurance_exp | DATE |  | 任意保険期限 |
| inspection_exp | DATE |  | 車検期限 |
| created_at | TIMESTAMP | NOT NULL | 登録日時 |
| updated_at | TIMESTAMP | NOT NULL | 更新日時 |

## **3.3 licenses（免許証）**

| **カラム名** | **型** | **制約** | **説明** |
| --- | --- | --- | --- |
| id | SERIAL | PK | 免許証ID |
| car_owner_id | INTEGER | FK(car_owners.id) | 車両所有者ID |
| license_no | VARCHAR | NOT NULL | 免許証番号 |
| expire_date | DATE | NOT NULL | 有効期限 |
| created_at | TIMESTAMP | NOT NULL | 登録日時 |
| updated_at | TIMESTAMP | NOT NULL | 更新日時 |

## **3.4 violations（違反記録）**

| **カラム名** | **型** | **制約** | **説明** |
| --- | --- | --- | --- |
| id | SERIAL | PK | 違反ID |
| car_owner_id | INTEGER | FK(car_owners.id) | 車両所有者ID |
| vehicle_id | INTEGER | FK(vehicles.id) | 車両ID |
| type | VARCHAR | NOT NULL | 違反種別 |
| date | DATE | NOT NULL | 発生日 |
| need_report | BOOLEAN |  | 始末書要否 |
| created_at | TIMESTAMP | NOT NULL | 登録日時 |
| updated_at | TIMESTAMP | NOT NULL | 更新日時 |

## **3.5 car_owners（車両所有者）**

| **カラム名** | **型** | **制約** | **説明** |
| --- | --- | --- | --- |
| id | SERIAL | PK | 車両所有者ID |
| user_id | INTEGER | FK(users.id) | 利用者ID（User） |
| name | VARCHAR | NOT NULL | 氏名 |
| contact | VARCHAR |  | 連絡先 |
| created_at | TIMESTAMP | NOT NULL | 登録日時 |
| updated_at | TIMESTAMP | NOT NULL | 更新日時 |

## **3.6 permits（許可証）**

| **カラム名** | **型** | **制約** | **説明** |
| --- | --- | --- | --- |
| id | SERIAL | PK | 許可証ID |
| car_owner_id | INTEGER | FK(car_owners.id) | 車両所有者ID |
| type | VARCHAR | NOT NULL | 許可証種別 |
| issue_date | DATE | NOT NULL | 発行日 |
| expire_date | DATE | NOT NULL | 有効期限 |
| created_at | TIMESTAMP | NOT NULL | 登録日時 |
| updated_at | TIMESTAMP | NOT NULL | 更新日時 |
| issue_date | DATE | NOT NULL | 発行日 |
| expire_date | DATE | NOT NULL | 有効期限 |
| created_at | TIMESTAMP | NOT NULL | 登録日時 |
| updated_at | TIMESTAMP | NOT NULL | 更新日時 |

## **4. バリデーション・エラーハンドリング**

- 必須項目未入力時はエラーメッセージ表示
- 日付は未来日チェック（免許証・保険・車検など）
- メールアドレス形式チェック
- CSVインポート時は全件バリデーション、エラー行は詳細表示

## **5. ロール・権限制御**

- 管理者：全機能利用可
- 一般利用者：自分の情報参照、一部編集、違反履歴確認のみ

## **6. セキュリティ**

- パスワードはハッシュ化
- JWT等のトークン認証
- 操作ログ記録（ユーザーID、操作内容、日時）
- SQLインジェクション・XSS等の対策

## **7. バックアップ・運用**

- DB定期バックアップ（夜間自動）
- 操作ログの保存期間：1年以上
- 障害時のリストア手順明記

## **8. テンプレート例（注意文書）**

- 違反種別ごとにテンプレートを用意
- 差し込み項目：利用者名、違反日、違反内容、累積違反数など

## **9. CSVフォーマット例**

- ヘッダ行必須、カンマ区切り、UTF-8
- 利用者：name,company,email,role
- 車両：user_id,number,type,insurance_exp,inspection_exp
- 違反：user_id,vehicle_id,type,date

## **10. その他**

- 画面レイアウトやER図、APIレスポンス例などは別途資料化
- 詳細な運用フローや管理者マニュアルも別途作成

**ご要望に応じて、画面モックアップやAPI仕様書、ER図、バリデーション詳細も作成可能です。**

**共有**

**エクスポート**

**書き直す**

上記詳細設計に画面レイアウトやAPI仕様、ER図、バリデーション仕様も追加して再構築してほしい

**上記詳細設計に画面レイアウトやAPI仕様、ER図、バリデーション仕様も追加して再構築してほしい**

答え

承知しました。

以下に、**画面レイアウト例／API仕様書／ER図（テキスト表現）／バリデーション仕様**を含めた、駐車場管理アプリの詳細設計書を再構築します。

## **1. 画面レイアウト（主要画面ワイヤーフレーム）**

## **1.1 ログイン画面**

```
+-------------------------------------+
|        駐車場管理システム           |
+-------------------------------------+
|  ユーザーID [______________]        |
|  パスワード [______________]        |
|  [ ログイン ]                       |
|                                     |
|  (エラーメッセージ表示欄)           |
+-------------------------------------+
```

## **1.2 ダッシュボード画面**

`text+------------------+------------------+
| [社員・業者管理] | [車両管理]       |
+------------------+------------------+
| [違反記録管理]   | [許可証発行]     |
+------------------+------------------+
| [注意文書作成]   | [CSVインポート]  |
+------------------+------------------+

[アラート]
- 免許証期限切れ: ○○さん（2025/07/15）
- 任意保険期限切れ: △△さん（2025/08/01）
- 違反累積超過: □□さん（5回）`

## **1.3 社員・業者管理画面**

`text+-----------------------------------------------+
| [新規登録] [CSVインポート]                    |
+-----------------------------------------------+
| 氏名 | 所属/業者 | 連絡先 | 権限 | [詳細][編集][削除] |
|------------------------------------------------|
| 山田太郎 | 総務部 | yama@example.com | 管理者 | ...    |
| ...                                            |
+------------------------------------------------+
| [検索] [リセット]                              |
+------------------------------------------------+`

## **1.4 車両管理画面**

`text+------------------------------------------------------+
| [新規登録] [CSVインポート]                            |
+------------------------------------------------------+
| 車両ナンバー | 車種 | 任意保険期限 | 車検期限 | 利用者 | [編集][削除] |
|-----------------------------------------------------------------------|
| 1234 | プリウス | 2025/09/30 | 2026/03/31 | 山田太郎 | ...           |
| ...                                                                   |
+-----------------------------------------------------------------------+
| [検索] [リセット]                                                    |
+-----------------------------------------------------------------------+`

## **1.5 違反記録管理画面**

`text+------------------------------------------------------+
| [新規登録] [CSVインポート]                            |
+------------------------------------------------------+
| 利用者 | 車両ナンバー | 違反種別 | 発生日 | 始末書要否 | [編集][削除] |
|-----------------------------------------------------------------------|
| 山田太郎 | 1234 | 許可証未掲示 | 2025/06/01 | 要         | ...         |
| ...                                                                   |
+-----------------------------------------------------------------------+
| [検索] [リセット]                                                    |
+-----------------------------------------------------------------------+`

## **1.6 許可証発行画面**

`text+----------------------------------------+
| 利用者選択 [▼]                         |
| 許可証種別 [従業員/工場内/外注業者]    |
| 有効期限 [2026/07/01]                  |
| [PDF発行] [印刷]                       |
+----------------------------------------+`

## **1.7 注意文書作成画面**

`text+----------------------------------------+
| 利用者選択 [▼]                         |
| 違反内容選択 [▼]                       |
| [文書生成]                             |
|                                        |
| --- 文書プレビュー ---                 |
| ○○さんへ                               |
| あなたは「許可証未掲示」の違反が累計5回となりました。|
| ...                                    |
| [メール送信] [印刷]                    |
+----------------------------------------+`

## **2. API仕様（主要エンドポイント）**

## **共通事項**

- 認証：JWTトークン（**`Authorization: Bearer <token>`**）
- レスポンス：JSON

## **2.1 認証**

## **POST /api/login**

| **項目** | **内容** |
| --- | --- |
| 説明 | ログイン |
| リクエスト | { "userId": "", "password": "" } |
| レスポンス | { "token": "", "role": "" } |
| エラー | 401 Unauthorized |

## **2.2 利用者**

## **GET /api/users**

| **項目** | **内容** |
| --- | --- |
| 説明 | 利用者一覧取得 |
| パラメータ | 検索条件（任意） |
| レスポンス | [ { "id": 1, "name": "...", ... }, ... ] |

## **POST /api/users**

| **項目** | **内容** |
| --- | --- |
| 説明 | 利用者新規登録 |
| リクエスト | { "name": "", "company": "", ... } |
| レスポンス | { "id": 1, ... } |
| エラー | 400 Bad Request |

## **2.3 車両**

## **GET /api/vehicles**

| **項目** | **内容** |
| --- | --- |
| 説明 | 車両一覧取得 |
| パラメータ | 検索条件（任意） |
| レスポンス | [ { "id": 1, "number": "...", ... }, ... ] |

## **POST /api/vehicles**

| **項目** | **内容** |
| --- | --- |
| 説明 | 車両新規登録 |
| リクエスト | { "user_id": 1, "number": "", ... } |
| レスポンス | { "id": 1, ... } |
| エラー | 400 Bad Request |

## **2.4 違反記録**

## **GET /api/violations**

| **項目** | **内容** |
| --- | --- |
| 説明 | 違反記録一覧取得 |
| パラメータ | 検索条件（任意） |
| レスポンス | [ { "id": 1, "user_id": 1, ... }, ... ] |

## **POST /api/violations**

| **項目** | **内容** |
| --- | --- |
| 説明 | 違反記録新規登録 |
| リクエスト | { "user_id": 1, "vehicle_id": 1, "type": "", "date": "" } |
| レスポンス | { "id": 1, ... } |
| エラー | 400 Bad Request |

## **2.5 許可証**

## **POST /api/permits**

| **項目** | **内容** |
| --- | --- |
| 説明 | 許可証発行 |
| リクエスト | { "user_id": 1, "type": "", "expire_date": "" } |
| レスポンス | { "id": 1, ... } |
| エラー | 400 Bad Request |

## **2.6 注意文書**

## **POST /api/notices**

| **項目** | **内容** |
| --- | --- |
| 説明 | 注意文書生成 |
| リクエスト | { "user_id": 1, "violation_type": "" } |
| レスポンス | { "notice_text": "..." } |

## **2.7 CSVインポート**

## **POST /api/import/csv**

| **項目** | **内容** |
| --- | --- |
| 説明 | CSVインポート |
| リクエスト | multipart/form-data（ファイル） |
| レスポンス | { "success": 10, "failed": 2, "errors": [...] } |

## **3. ER図（テキスト表現）**

`text[users] 1 --- n [vehicles]
[users] 1 --- n [licenses]
[users] 1 --- n [violations]
[users] 1 --- n [permits]
[vehicles] 1 --- n [violations]`

- users（利用者）
    
    ├── vehicles（車両）
    
    ├── licenses（免許証）
    
    ├── violations（違反記録）
    
    └── permits（許可証）
    
- vehicles（車両）
    
    └── violations（違反記録）
    

## **4. バリデーション仕様**

## **4.1 共通**

- 必須項目未入力時はエラー
- 文字列は最大長制限（例：氏名50文字、車両ナンバー20文字）
- 日付はYYYY-MM-DD形式
- メールアドレスはRFC準拠
- IDは整数

## **4.2 利用者**

| **項目** | **バリデーション内容** |
| --- | --- |
| 氏名 | 必須、50文字以内 |
| 所属/業者名 | 100文字以内 |
| 連絡先 | 必須、メール形式 |
| 権限区分 | "admin"または"user" |
| パスワード | 必須、8文字以上、英数字混在 |

## **4.3 車両**

| **項目** | **バリデーション内容** |
| --- | --- |
| 車両ナンバー | 必須、20文字以内 |
| 車種 | 30文字以内 |
| 任意保険期限 | 日付、未来日 |
| 車検期限 | 日付、未来日 |

## **4.4 免許証**

| **項目** | **バリデーション内容** |
| --- | --- |
| 免許証番号 | 必須、20文字以内 |
| 有効期限 | 必須、日付、未来日 |

## **4.5 違反記録**

| **項目** | **バリデーション内容** |
| --- | --- |
| 違反種別 | 必須、定義済み種別のみ |
| 発生日 | 必須、日付、過去日または当日 |
| 利用者ID | 必須、存在チェック |
| 車両ID | 必須、存在チェック |

## **4.6 許可証**

| **項目** | **バリデーション内容** |
| --- | --- |
| 許可証種別 | 必須、定義済み種別のみ |
| 有効期限 | 必須、日付、未来日 |
| 利用者ID | 必須、存在チェック |

## **4.7 CSVインポート**

- ファイル形式：CSV（UTF-8、ヘッダ行必須）
- 各行ごとに上記バリデーションを適用
- エラー行は詳細内容とともにレスポンスで返却

## **5. その他**

- 画面デザインはレスポンシブ対応（PC・タブレット）
- 操作ログは管理者のみ参照可
- APIエラー時はHTTPステータス＋エラー詳細をJSONで返却
- セッションタイムアウト：30分