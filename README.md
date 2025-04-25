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

または、ソースからビルド：

```bash
make install
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

### バージョン情報の表示

```bash
shortcut-cli version
```

### Story検索

```bash
shortcut-cli story search "検索キーワード"
```

オプション：
- `--type`: ストーリータイプでフィルタ（feature/bug/chore）
- `--state`: ストーリーの状態でフィルタ
- `--epic`: Epic IDでフィルタ
- `--owner`: オーナーIDでフィルタ
- `--created`: 作成日でフィルタ（YYYY-MM-DD）
- `--updated`: 更新日でフィルタ（YYYY-MM-DD）

### Epic情報の参照

```bash
shortcut-cli epic get <epic-id>
```

### Story情報の参照

```bash
shortcut-cli story get <story-id>
```

## 開発

### 必要条件

- Go 1.21以上
- Make
- [golangci-lint](https://golangci-lint.run/) (自動インストール可能)

### ビルドコマンド

以下のMakeコマンドが利用可能です：

```bash
make build      # バイナリのビルド
make test       # テストの実行
make lint       # リントの実行
make clean      # ビルド成果物の削除
make deps       # 依存関係の更新
make all        # lint、test、buildを順番に実行
make install    # バイナリのインストール
```

### バージョン管理

バイナリには以下の情報が含まれます：
- バージョン番号（gitのタグまたはコミットハッシュ）
- ビルド時刻

これらの情報は `shortcut-cli version` コマンドで確認できます。

## ライセンス

MIT License
