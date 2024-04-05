# kokkai-api

国会議事録検索用APIのラッパー

<https://kokkai.ndl.go.jp/api.html>

すべての検索用パラメータとすべてのAPIレスポンスの型に対応しています。
(2024/04/05時点)

## Example

```golang
var params Params
 params.Any("国会議事録")
 result := GetKaigi(http.DefaultClient, params)
 fmt.Println(result.Result.NumberOfRecords)
 // output: 総結果件数
```
