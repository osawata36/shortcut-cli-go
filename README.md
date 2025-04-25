# Shortcut CLI Tool

ShortcutのAPIをコマンドラインから操作するためのCLIツールです。

## 機能

- Shortcut APIの検索機能
- Epic情報の参照
- Story情報の参照
- ページネーション対応
- APIレート制限（429エラー）へのリトライ対応

## インストール

```bash
go install github.com/osawata36/shortcut-cli-go/cmd/shortcut@latest
```

## 設定

1. Shortcut APIトークンを取得します
   - https://app.shortcut.com/settings/account/api-tokens にアクセス
   - 新しいAPIトークンを生成

2. 環境変数に設定します

```bash
export SHORTCUT_API_TOKEN="your-api-token"
```

## 使い方

### Story検索

```bash
shortcut story search "検索キーワード"
```

### Epic情報の参照

```bash
shortcut epic get <epic-id>
```

### Story情報の参照

```bash
shortcut story get <story-id>
```

## 開発

### 必要条件

- Go 1.21以上

### ビルド方法

```bash
go build -o shortcut ./cmd/shortcut
```

### テスト実行

```bash
go test ./...
```

## ライセンス

MIT License
