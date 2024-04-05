package kokkai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Get(string) (*http.Response, error)
}

// 会議単位簡易出力をリクエストする。
func GetKani(client HTTPClient, params Params) *KaniResult {
	uri := parseURI(KaniURI, params)
	var result KaniResult
	result.URI = uri
	if err := get(client, uri, &result.Result); err != nil {
		result.Err = err
	}
	return &result
}

// 会議単位出力をリクエストする。
func GetKaigi(client HTTPClient, params Params) *KaigiResult {
	uri := parseURI(KaigiURI, params)
	var result KaigiResult
	result.URI = uri
	if err := get(client, uri, &result.Result); err != nil {
		result.Err = err
	}
	return &result
}

// 発言単位出力をリクエストする。
func GetHatsugen(client HTTPClient, params Params) *HatsugenResult {
	uri := parseURI(HatsugenURI, params)
	var result HatsugenResult
	result.URI = uri
	if err := get(client, uri, &result.Result); err != nil {
		result.Err = err
	}
	return &result
}

func get(client HTTPClient, uri string, result interface{}) error {
	resp, err := client.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, result); err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not ok: %s", resp.Status)
	}
	return nil
}
