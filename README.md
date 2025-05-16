# ShortUrlServer

## Overview

**URLをカスタムで設定できるサービス**です。

ここで利用できます。
[https://rk2.uk](https://rk2.uk)

## REST API Usage

REST APIに対応しています。
GUIアプリと機能は同じです。

### Requests

`"expire_in"`キーは、正規表現`` `[0-9]+[smhd]` ``を受け付けます。
`"10m", "30h", "2d"`のように指定してください。
指定しない場合、有効期限は設定されません。

`"base_url"`のみの指定の場合、ランダムIDがセットされます。

| key | 説明 | 必須/デフォルト値 | 競合する値 |
| :-- | :-- | :-- | :-- |
| `base_url` | リダイレクト先URL | 必須 | |
| `use_uppercase`| ランダムIDに英大文字を含めるか | `false` | `custom_id` |
| `use_lowercase`| ランダムIDに英小文字を含めるか | `true` | `custom_id` |
| `use_numbers`| ランダムIDに数字を含めるか | `true` | `custom_id` |
| `id_length`| ランダムIDの文字数 | `6` | `custom_id` |
| `custom_id`| 設定するカスタムID<br>(最大文字数100文字) | ランダムIDを採用 | `use_uppercase`, `use_lowercase`, `use_numbers`, `id_length` |
| `expire_in`| リンクの有効期間 | 無期限 | |
| `sand_cushion`| クッションページを使用するか | `false` | |

1. ランダムIDでURL生成する例

```JSON
{
    "base_url": "https://example.com",
    "use_uppercase": true,
    "use_lowercase": false,
    "use_numbers": true,
    "id_length": 5,
    "expire_in": "10h",
    "sand_cushion": true
}
```

2. カスタムIDでURL生成する例

```JSON
{
    "base_url": "https://example.com",
    "custom_id": "example",
    "expire_in": "10h",
    "sand_cushion": true
}
```

### Responces

1. OK Responce

```JSON
{
    "base_url": "https://example.com",
    "custom_id": "example",
    "expire_in": "10h",
    "sand_cushion": true
}
```

2. Conflict Responce
カスタムIDリクエスト時、IDが既存の時に返されます。

```JSON
{
    "error": "custom_id already used."
}
```

3. Other error Responces

```JSON
{
    "error": "Error message"
}
```

## Usage

1. 開発環境では

```bash
docker compose -f compose.dev.yml up -d --build
```

Airを利用してホットリロード開発ができます。

2. デプロイ環境では

```bash
docker compose -f compose.prod.yml up -d --build
```

マルチステージングにより、バイナリにビルドした後に Alpine コンテナで実行されます。

## Contact

Outtech105k

[Twitter(X)](https://x.com/105techno)
