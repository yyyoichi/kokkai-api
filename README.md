# kokkai-api

国会議事録検索用APIのラッパー

<https://kokkai.ndl.go.jp/api.html>

## サポート

議事の全件取得に対応しました。

すべての検索用パラメータとすべてのAPIレスポンスの型に対応しています。
(2024/04/25時点)

## Example

```golang
// 国会議事録にヒットする初めの10件の会議情報を取得する。
var params = NewParam()
params.Any("国会議事録")
params.MaximumRecords(100)
result := GetKaigi(http.DefaultClient, params)
if result.Err != nil {
    return result.Err
}
fmt.Println(result.Result.NumberOfReturn) // 返戻件数
```

### 検索条件についてすべて取得する

```golang
// 国会議事録にヒットする発言をすべて取得する。
var params = NewParam()
params.Any("国会議事録")
params.MaximumRecords(100)
client := &Client[*HatsugenResult]{
    HTTPClient: http.DefaultClient,
    Get:        GetHatsugen,
    Interval:   time.Duration(1) * time.Second,
}
for result := range client.GetAll(context.Background(), params) {
    fmt.Println(result.Result.NumberOfReturn) // 返戻件数
}
```
