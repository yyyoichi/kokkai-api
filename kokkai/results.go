package kokkai

// 会議単位簡易出力
type KaniResult struct {
	URI    string
	Err    error
	Result struct {
		Message            string   `json:"message"`            // エラーメッセージ
		Details            []string `json:"details"`            // エラーメッセージの詳細（※検索条件の入力誤りの場合のみ、※検索条件の入力誤りが複数ある場合のみ）
		NumberOfRecords    int      `json:"numberOfRecords"`    // 総結果件数
		NumberOfReturn     int      `json:"numberOfReturn"`     // 返戻件数
		StartRecord        int      `json:"startRecord"`        // 開始位置
		NextRecordPosition int      `json:"nextRecordPosition"` // 次開始位置
		MeetingRecord      []struct {
			IssueID       string `json:"issueID"`       // 会議録ID
			ImageKind     string `json:"imageKind"`     // イメージ種別（会議録・目次・索引・附録・追録）
			SearchObject  int    `json:"searchObject"`  // 検索対象箇所（議事冒頭・本文）
			Session       int    `json:"session"`       // 国会回次
			NameOfHouse   string `json:"nameOfHouse"`   // 院名
			NameOfMeeting string `json:"nameOfMeeting"` // 会議名
			Issue         string `json:"issue"`         // 号数
			Date          string `json:"date"`          // 開催日付
			Closing       string `json:"closing"`       // 閉会中フラグ
			SpeechRecord  []struct {
				SpeechOrder int    `json:"speechOrder"` // 発言番号
				Speaker     string `json:"speaker"`     // 発言者名
				SpeechURL   string `json:"speechURL"`   // 発言URL
			} `json:"speechRecord"`
			MeetingURL string `json:"meetingURL"` // 会議録テキスト表示画面のURL ,
			PdfURL     string `json:"pdfURL"`     // 会議録PDF表示画面のURL（※存在する場合のみ） ,
		} `json:"meetingRecord"`
	}
}

func (r *KaniResult) NextRecordPosition() int { return r.Result.NextRecordPosition }

// 会議単位出力
type KaigiResult struct {
	URI    string
	Err    error
	Result struct {
		Message            string   `json:"message"`            // エラーメッセージ
		Details            []string `json:"details"`            // エラーメッセージの詳細（※検索条件の入力誤りの場合のみ、※検索条件の入力誤りが複数ある場合のみ）
		NumberOfRecords    int      `json:"numberOfRecords"`    // 総結果件数
		NumberOfReturn     int      `json:"numberOfReturn"`     // 返戻件数
		StartRecord        int      `json:"startRecord"`        // 開始位置
		NextRecordPosition int      `json:"nextRecordPosition"` // 次開始位置
		MeetingRecord      []struct {
			IssueID       string `json:"issueID"`       // 会議録ID
			ImageKind     string `json:"imageKind"`     // イメージ種別（会議録・目次・索引・附録・追録）
			SearchObject  int    `json:"searchObject"`  // 検索対象箇所（議事冒頭・本文）
			Session       int    `json:"session"`       // 国会回次
			NameOfHouse   string `json:"nameOfHouse"`   // 院名
			NameOfMeeting string `json:"nameOfMeeting"` // 会議名
			Issue         string `json:"issue"`         // 号数
			Date          string `json:"date"`          // 開催日付
			Closing       string `json:"closing"`       // 閉会中フラグ
			SpeechRecord  []struct {
				SpeechID        string `json:"speechID"`        // 発言ID
				SpeechOrder     int    `json:"speechOrder"`     // 発言番号
				Speaker         string `json:"speaker"`         // 発言者名
				SpeakerYomi     string `json:"speakerYomi"`     // 発言者よみ
				SpeakerGroup    string `json:"speakerGroup"`    // 発言者所属会派
				SpeakerPosition string `json:"speakerPosition"` // 発言者肩書き
				SpeakerRole     string `json:"speakerRole"`     // 発言者役割
				Speech          string `json:"speech"`          // 発言
				StartPage       int    `json:"startPage"`       // 発言が掲載されている開始ページ
				CreateTime      string `json:"createTime"`      // レコード登録日時
				UpdateTime      string `json:"updateTime"`      // レコード更新日時
				SpeechURL       string `json:"speechURL"`       // 発言URL
			} `json:"speechRecord"` // 発言情報
			MeetingURL string `json:"meetingURL"` // 会議録テキスト表示画面のURL ,
			PdfURL     string `json:"pdfURL"`     // 会議録PDF表示画面のURL（※存在する場合のみ） ,
		} `json:"meetingRecord"` // 会議録情報
	}
}

func (r *KaigiResult) NextRecordPosition() int { return r.Result.NextRecordPosition }

// 発言単位出力
type HatsugenResult struct {
	URI    string
	Err    error
	Result struct {
		Message            string   `json:"message"`            // エラーメッセージ
		Details            []string `json:"details"`            // エラーメッセージの詳細（※検索条件の入力誤りの場合のみ、※検索条件の入力誤りが複数ある場合のみ）
		NumberOfRecords    int      `json:"numberOfRecords"`    // 総結果件数
		NumberOfReturn     int      `json:"numberOfReturn"`     // 返戻件数
		StartRecord        int      `json:"startRecord"`        // 開始位置
		NextRecordPosition int      `json:"nextRecordPosition"` // 次開始位置
		SpeechRecord       []struct {
			SpeechID        string `json:"speechID"`        // 発言ID
			IssueID         string `json:"issueID"`         // 会議録ID
			ImageKind       string `json:"imageKind"`       // イメージ種別（会議録・目次・索引・附録・追録）
			SearchObject    int    `json:"searchObject"`    // 検索対象箇所（議事冒頭・本文）
			Session         int    `json:"session"`         // 国会回次
			NameOfHouse     string `json:"nameOfHouse"`     // 院名
			NameOfMeeting   string `json:"nameOfMeeting"`   // 会議名
			Issue           string `json:"issue"`           // 号数
			Date            string `json:"date"`            // 開催日付
			Closing         string `json:"closing"`         // 閉会中フラグ
			SpeechOrder     int    `json:"speechOrder"`     // 発言番号
			Speaker         string `json:"speaker"`         // 発言者名
			SpeakerYomi     string `json:"speakerYomi"`     // 発言者よみ
			SpeakerGroup    string `json:"speakerGroup"`    // 発言者所属会派
			SpeakerPosition string `json:"speakerPosition"` // 発言者肩書き
			SpeakerRole     string `json:"speakerRole"`     // 発言者役割
			Speech          string `json:"speech"`          // 発言
			StartPage       int    `json:"startPage"`       // 発言が掲載されている開始ページ
			SpeechURL       string `json:"speechURL"`       // 発言URL
			SeetingURL      string `json:"meetingURL"`      // 会議録テキスト表示画面のURL
			PdfURL          string `json:"pdfURL"`          // 会議録PDF表示画面のURL（※存在する場合のみ）
		} `json:"speechRecord"` // 発言情報
	}
}

func (r *HatsugenResult) NextRecordPosition() int { return r.Result.NextRecordPosition }
