# ShortUrlServer

## Overview

**URLをカスタムで設定できるサービス**です。

ここで利用できます。
[https://rk2.uk](https://rk2.uk)

## REST API Usage

REST APIに対応しています。
GUIアプリと機能は同じです。

`"expire_in"`キーは、正規表現`\`[0-9]+[smhd]\``を受け付けます。
`"10m", "30h", "2d"`のように指定してください。
指定しない場合、有効期限は設定されません。

`"base_url"`のみの指定の場合、ランダムIDがセットされます。

1. ランダムIDでURL生成
```JSON
{
    "base_url": "https://example.com",  // リダイレクト先URL, 必須
    "use_uppercase": true,              // ランダムIDに英大文字を含めるか。 Default: false
    "use_lowercase": false,             // ランダムIDに英小文字を含めるか。 Default: true
    "use_numbers": true,                // ランダムIDに数字を含めるか。 Default: true
    "id_length": 5,                     // ランダムIDの文字数。 Default: 6
    "expire_in": "10h",                 // リンクの有効期間。 Default: 無期限
    "sand_cushion": true,               // クッションページを使用するか。 Default: false
}
```

2. カスタムIDでURL生成
```JSON
{
    "base_url": "https://example.com",  // リダイレクト先URL。 必須
    "custom_id": "example",             // 設定したいカスタムID。
    "expire_in": "10h",                 // リンクの有効期間。 Default: 無期限
    "sand_cushion": true,               // クッションページを使用するか。 Default: false
}
```

## Contact

Outtech105k

[Twitter(X)](https://x.com/105techno)
