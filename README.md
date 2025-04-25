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

### コマンド一覧

```bash
shortcut-cli [command] [subcommand] [options]
```

利用可能なコマンド：
- `version`: バージョン情報の表示
- `epic`: Epic関連の操作
- `story`: Story関連の操作

### Epic関連のコマンド

#### Epic一覧の取得
```bash
shortcut-cli epic list
```

#### Epic情報の参照
```bash
shortcut-cli epic get <epic-id>
```

#### Epic検索
```bash
shortcut-cli epic search [検索キーワード]
```

オプション：
- `--state`: Epicの状態でフィルタ（unstarted/started/done）
- `--owner`: オーナーでフィルタ
- `--created`: 作成日でフィルタ（YYYY-MM-DD）
- `--updated`: 更新日でフィルタ（YYYY-MM-DD）

### Story関連のコマンド

#### Story情報の参照
```bash
shortcut-cli story get <story-id>
```

#### Story検索
```bash
shortcut-cli story search [検索キーワード]
```

オプション：
- `--type`: ストーリータイプでフィルタ（feature/bug/chore）
- `--state`: ストーリーの状態でフィルタ
- `--epic`: Epic IDでフィルタ
- `--owner`: オーナーでフィルタ
- `--created`: 作成日でフィルタ（YYYY-MM-DD）
- `--updated`: 更新日でフィルタ（YYYY-MM-DD）

### 日付フィルタの書式

日付フィルタ（`--created`、`--updated`）では以下の形式が使用可能です：

- 特定の日付: `YYYY-MM-DD`
- 相対的な日付: `today`、`yesterday`、`tomorrow`
- 日付範囲:
  - `YYYY-MM-DD..YYYY-MM-DD`: 指定期間
  - `YYYY-MM-DD..*`: 指定日以降
  - `*..YYYY-MM-DD`: 指定日以前
  - `today..*`: 今日以降
  - `*..today`: 今日以前

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
