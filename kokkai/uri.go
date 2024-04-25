package kokkai

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	KaniURI     = "https://kokkai.ndl.go.jp/api/meeting_list"
	KaigiURI    = "https://kokkai.ndl.go.jp/api/meeting"
	HatsugenURI = "https://kokkai.ndl.go.jp/api/speech"
)

func parseURI(baseURI string, params Params) string {
	return fmt.Sprintf("%s?%s", baseURI, params.encode())
}

// 検索条件
//
// - {検索条件}は「パラメータ名=値」の形式で指定し、UTF-8でURLエンコードしてください。
//
// - 複数のパラメータで検索する場合には、半角の&（U+0026）で接続してください。
//
// - 検索条件部分は全体で2000バイトが上限です。
//
// - 院名、会議名、検索語、発言者名、開会日付／始点、開会日付／終点、発言番号、発言者肩書き、発言者所属会派、発言者役割、発言ID、冊子ID、国会回次From、国会回次To、号数From、号数Toのいずれも指定がなかった場合には、エラーになります。
type Params struct {
	values url.Values
}

func NewParam() Params {
	return Params{values: url.Values{}}
}

func (p *Params) encode() string {
	p.values.Add("recordPacking", "json")
	return p.values.Encode()
}

// 開始位置	検索結果の取得開始位置を「1～検索件数」の範囲で指定可能。
// 省略時のデフォルト値は「1」
func (p *Params) StartRecord(val int) {
	p.values.Add("startRecord", strconv.Itoa(val))
}

// 開始位置 の再設定。
func (p *Params) ResetStartRecord(val int) {
	p.values.Del("startRecord")
	p.values.Add("startRecord", strconv.Itoa(val))
}

// 一回の最大取得件数	一回のリクエストで取得できるレコード数を、会議単位簡易出力、発言単位出力の場合は「1～100」、会議単位出力の場合は「1～10」の範囲で指定可能。
// 省略時のデフォルト値は、会議単位簡易出力、発言単位出力の場合は「30」、会議単位出力の場合は「3」
func (p *Params) MaximumRecords(val int) {
	p.values.Add("maximumRecords", strconv.Itoa(val))
}

// 院名	院名として「衆議院」「参議院」「両院」「両院協議会」のいずれかを指定可能。「両院」と「両院協議会」の結果は同じ。
// 省略可（省略時は検索条件に含めない）。また、指定可能な値以外を指定した場合も、検索条件に含めない。
func (p *Params) NameOfHouse(val string) {
	p.values.Add("nameOfHouse", val)
}

// 会議名	本会議、委員会等の会議名（ひらがな可）を指定可能。部分一致検索。半角スペース（U+0020）を区切り文字として複数指定した場合は、指定した語のOR検索となる。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) NameOfMeeting(val string) {
	p.values.Add("nameOfMeeting", val)
}

// 検索語	発言内容等に含まれる言葉を指定可能。部分一致検索。半角スペース（U+0020）を区切り文字として複数指定した場合は、指定した語のAND検索となる。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) Any(val string) {
	p.values.Add("any", val)
}

// 発言者名	発言者名（議員名はひらがな可）を指定可能。部分一致検索。半角スペース（U+0020）を区切り文字として複数指定した場合は、指定した語のOR検索となる。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) Speaker(val string) {
	p.values.Add("speaker", val)
}

// 開会日付／始点	検索対象とする会議の開催日の始点を「YYYY-MM-DD」の形式で指定可能。
// 省略可（省略時は「0000-01-01」が指定されたものとして検索する）。
func (p *Params) From(val string) {
	p.values.Add("from", val)
}

// 開会日付／終点	検索対象とする会議の開催日の終点を「YYYY-MM-DD」の形式で指定可能。
// 省略可（省略時は「9999-12-31」が指定されたものとして検索する）。
func (p *Params) Until(val string) {
	p.values.Add("until", val)
}

// 追録・附録指定	検索対象を追録・附録に限定するか否かを「true」「false」で指定可能。
// 省略可（省略時は「false」（限定しない）が指定されたものとして検索する）。
func (p *Params) SupplementAndAppendix(val bool) {
	p.values.Add("supplementAndAppendix", fmt.Sprintf("%v", val))
}

// 目次・索引指定	検索対象を目次・索引に限定するか否かを「true」「false」で指定可能。
// 省略可（省略時は「false」（限定しない）が指定されたものとして検索する）。
func (p *Params) ContentsAndIndex(val bool) {
	p.values.Add("contentsAndIndex", fmt.Sprintf("%v", val))
}

// 議事冒頭・本文指定	検索語（パラメータ名：any）を指定して検索する際の検索対象箇所を「冒頭」「本文」「冒頭・本文」のいずれかで指定可能。
// 省略可（省略時は「冒頭・本文」が指定されたものとして検索する）。検索語を指定しなかった時は検索条件には含めない。
func (p *Params) SearchRange(val string) {
	p.values.Add("searchRange", val)
}

// 閉会中指定	検索対象を閉会中の会議録に限定するか否かを「true」「false」で指定可能。
// 省略可（省略時は「false」（限定しない）が指定されたものとして検索する）。
func (p *Params) Closing(val bool) {
	p.values.Add("closing", fmt.Sprintf("%v", val))
}

// 発言番号	発言番号を0以上の整数（例：発言番号10の場合は「speechNumber=10」）で指定可能。完全一致検索。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) SpeechNumber(val int) {
	p.values.Add("speechNumber", strconv.Itoa(val))
}

// 発言者肩書き	発言者の肩書きを指定可能。部分一致検索。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) SpeakerPosition(val string) {
	p.values.Add("speakerPosition", val)
}

// 発言者所属会派	発言者の所属会派を指定可能。部分一致検索（なお、登録されているデータは正式名称のみ）。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) SpeakerGroup(val string) {
	p.values.Add("speakerGroup", val)
}

// 発言者役割	発言者の役割として「証人」「参考人」「公述人」のいずれかを指定可能。
// 省略可（省略時は検索条件に含めない）。指定可能な値以外を指定した場合はエラーになる。
func (p *Params) SpeakerRole(val string) {
	p.values.Add("speakerRole", val)
}

// 発言ID	発言を一意に識別するIDとして、「会議録ID（パラメータ名：issueID。21桁の英数字）_発言番号（会議録テキスト表示画面で表示されている各発言に付されている、先頭に0を埋めて3桁にした数字。4桁の場合は4桁の数字）」の書式で指定可能（例：「100105254X00119470520_000」）。完全一致検索。
// 省略可（省略時は検索条件に含めない）。書式が適切でない場合にはエラーになる。
func (p *Params) SpeechID(val string) {
	p.values.Add("speechID", val)
}

// 会議録ID	会議録（冊子）を一意に識別するIDとして、会議録テキスト表示画面の「会議録テキストURLを表示」リンクで表示される21桁の英数字で指定可能（例：「100105254X00119470520」）。完全一致検索。
// 省略可（省略時は検索条件に含めない）。書式が適切でない場合にはエラーになる。
func (p *Params) IssueID(val string) {
	p.values.Add("issueID", val)
}

// 国会回次From	検索対象とする国会回次の始まり（開始回）を3桁までの自然数で指定可能。国会回次Toと組み合わせて指定した場合には範囲指定検索、国会回次From単独で指定した場合は当該の回次のみを完全一致検索。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) SessionFrom(val int) {
	p.values.Add("sessionFrom", strconv.Itoa(val))
}

// 国会回次To	検索対象とする国会回次の終わり（終了回）を3桁までの自然数で指定可能。国会回次Fromと組み合わせて指定した場合には範囲指定検索、国会回次To単独で指定した場合は当該の回次のみを完全一致検索。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) SessionTo(val int) {
	p.values.Add("sessionTo", strconv.Itoa(val))
}

// 号数From	検索対象とする号数の始まり（開始号）を3桁までの整数で指定可能（目次・索引・附録・追録は0号扱い）。号数Toと組み合わせて指定した場合には範囲指定検索、号数From単独で指定した場合は当該の回次のみを完全一致検索。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) IssueFrom(val int) {
	p.values.Add("issueFrom", strconv.Itoa(val))
}

// 号数To	検索対象とする号数の終わり（終了号）を3桁までの整数で指定可能（目次・索引・附録・追録は0号扱い）。号数Fromと組み合わせて指定した場合には範囲指定検索、号数To単独で指定した場合は当該の回次のみを完全一致検索。
// 省略可（省略時は検索条件に含めない）。
func (p *Params) IssueTo(val int) {
	p.values.Add("issueTo", strconv.Itoa(val))
}

// 応答形式	検索リクエストに対する応答ファイルの形式として、「xml」「json」のいずれかを指定可能。
// 省略可（省略時は「xml」が指定されたものとして検索する）。
func (p *Params) RecordPacking(val string) {
	p.values.Add("recordPacking", val)
}
