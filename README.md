# kokkai-api

国会議事録検索用APIのGo言語ラッパー

<https://kokkai.ndl.go.jp/api.html>

## Supports

- 会議単位簡易出力, 会議単位出力, 発言単位出力 に対応しています。
- 任意検索条件の再帰的全件取得に対応しています。
  - `iter`パッケージを利用しています。
- APIコールの時間間隔の調整が可能です。
- すべての検索用パラメータに対応しています。（2024/04/25時点）

## Install

```shell
go get "github.com/yyyoichi/kokkai-api"
```

## Example

```golang
import (
    "fmt"
    "log"

    kokkaiapi "github.com/yyyoichi/kokkai-api"
)

func main() {
    // 国会回次209の科学技術に関する議事録を「会議単位簡易出力」で最大3件ずつ、再帰的に取得する。
    p := kokkaiapi.NewParam()
    p.Any("科学技術")
    p.RecordPacking("json")
    p.SessionFrom(209)
    p.SessionTo(209)
    p.MaximumRecords(3)
    // exp return 7 records
    for result, err := range kokkaiapi.IterKaniResult(p) {
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("numberOfReturn", result.NumberOfReturn)
    }

    // Output:
    // numberOfReturn 3
    // numberOfReturn 3
    // numberOfReturn 1
}
```
