package kokkaiapi_test

import (
	"fmt"
	"log"
	"net/http"
	"time"

	kokkaiapi "github.com/yyyoichi/kokkai-api"
)

func Example() {
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

func ExampleIterResult() {
	// httpクライアントに任意の構造体を利用して、議事録を取得する。

	request := kokkaiapi.DefaultHatsugenRequest
	// `Get(uri string) (*http.Response, error)` メソッドを持つ構造体が必要。
	request.Client.HttpClient = &http.Client{
		// custom http client
		Timeout: time.Duration(1 * time.Second),
	}
	p := kokkaiapi.NewParam()
	p.Any("鬼滅の刃")
	p.RecordPacking("json")
	p.SessionFrom(209) // 議事録は存在しない。
	p.SessionTo(209)   // 210には議事録があったりする。
	p.MaximumRecords(3)

	// return no record
	for result, err := range kokkaiapi.IterResult(p, request) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("numberOfReturn", result.NumberOfReturn)
	}
	// Output:
	// numberOfReturn 0
}

func ExampleGetResult() {
	// 国会回次209の科学技術に関する議事録を「会議単位簡易出力」で一度だけ取得する。
	p := kokkaiapi.NewParam()
	p.Any("科学技術")
	p.RecordPacking("json")
	p.SessionFrom(209)
	p.SessionTo(209)
	p.MaximumRecords(10)
	// exp return 7 records
	for result, err := range kokkaiapi.IterKaniResult(p) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("numberOfReturn", result.NumberOfReturn)
	}

	// Output:
	// numberOfReturn 7
}
