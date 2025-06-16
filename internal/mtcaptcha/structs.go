package mtcaptcha

type param struct {
	Key   string
	Value string
}

type GetChallengeRes struct {
	Schema string `json:"schema"`
	Code   int    `json:"code"`
	Result struct {
		Context struct {
			DevMode   bool `json:"devMode"`
			HidePowBy bool `json:"hidePowBy"`
			HideTerms bool `json:"hideTerms"`
		} `json:"context"`
		Challenge struct {
			Ct          string `json:"ct"`
			Ctttl       int    `json:"ctttl"`
			HasTextChlg bool   `json:"hasTextChlg"`
			TextChlg    struct {
				Textlen int `json:"textlen"`
			} `json:"textChlg"`
			HasFoldChlg bool `json:"hasFoldChlg"`
			FoldChlg    struct {
				Fseed  string `json:"fseed"`
				Fslots int    `json:"fslots"`
				Fdepth int    `json:"fdepth"`
				PreRes bool   `json:"preRes"`
			} `json:"foldChlg"`
			HasWaitChlg bool `json:"hasWaitChlg"`
			WaitChlg    struct {
				Time string `json:"time"`
			} `json:"waitChlg"`
		} `json:"challenge"`
		Field3 string `json:"_"`
	} `json:"result"`
}

type GetImageRes struct {
	Schema string `json:"schema"`
	Code   int    `json:"code"`
	Result struct {
		Img struct {
			Image64 string `json:"image64"`
		} `json:"img"`
		Field2 string `json:"_"`
	} `json:"result"`
}

type VerifyRes struct {
	Schema string `json:"schema"`
	Code   int    `json:"code"`
	Result struct {
		VerifyResult struct {
			IsVerified    bool `json:"isVerified"`
			VerifiedToken struct {
				Vt    string `json:"vt"`
				Vtttl int    `json:"vtttl"`
			} `json:"verifiedToken"`
		} `json:"verifyResult"`
		Field2 string `json:"_"`
	} `json:"result"`
}
