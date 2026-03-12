# kotob_server 仕様書

## 1. 概要
CLI翻訳ツール `kotob` のコアロジックをベースとした、REST APIサーバー。
Gemini APIを介したテキストの翻訳・変換処理を、外部アプリケーションからHTTPリクエスト経由で利用可能にすることを目的とする。

## 2. エンドポイント仕様

### POST /api/translate
テキストの翻訳・変換処理を実行する。

#### リクエストボディ (JSON)
| フィールド | 型 | 必須 | 説明 |
| :--- | :--- | :--- | :--- |
| `text` | string | Yes | 変換対象のテキスト |
| `to` | string | No | 出力言語 (未指定時は en ⇔ ja) |
| `from` | string | No | 入力言語 (デフォルト: auto) |
| `model` | string | No | 使用するGeminiモデルの指定 |
| `system` | string | No | システムプロンプトの指定 |
| `stream` | boolean | No | ストリーム出力の有効化 (デフォルト: true) |

#### レスポンス仕様
- **通常レスポンス (stream: false)**
    - Status: `200 OK`
    - Body: 翻訳結果を含むJSON
- **ストリームレスポンス (stream: true)**
    - Status: `200 OK`
    - Content-Type: `text/event-stream` (SSE形式)
- **エラーレスポンス**
    - Status: `4xx` / `5xx`
    - Body: エラーメッセージを含むJSON

## 3. 設定の優先順位
設定値は以下の順序で評価され、上位にある項目が優先される。

1. **HTTPリクエストボディ**: 各リクエストごとの個別設定
2. **OS環境変数**: サーバー起動時に設定 (例: `KOTOB_API_KEY`)
3. **設定ファイル (kotob_server.json)**: 実行ディレクトリ直下
4. **ユーザー設定ファイル**: `~/.config/kotob/kotob_server.json`
5. **アプリケーション・デフォルト値**: (例: `--to ja`)

## 4. 入出力仕様
- **成功時**:
    - `application/json` 形式で構造化データを返却。
- **失敗時**:
    - 適切なHTTPステータスコードを返し、エラー内容をJSONで返却。

- **ステートレス設計**: サーバーはリクエストごとに完結し、複数のクライアントからの並行リクエストに対応する。