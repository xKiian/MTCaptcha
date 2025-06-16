package mtcaptcha

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	tlsclient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type MTCaptcha struct {
	SiteKey   string
	HostName  string
	Client    tlsclient.HttpClient
	UserAgent string
	
	sessionID      string
	resetTS        string
	cookie         string
	challengeToken string
}

func New(siteKey, hostName, proxy string) (*MTCaptcha, error) {
	sessionID := "S1" + uuid.New().String()
	
	options := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(5),
		tlsclient.WithClientProfile(profiles.Chrome_133),
		tlsclient.WithCookieJar(tlsclient.NewCookieJar()),
	}
	if proxy != "" {
		options = append(options, tlsclient.WithProxyUrl(proxy))
		options = append(options, tlsclient.WithInsecureSkipVerify())
	}
	client, err := tlsclient.NewHttpClient(tlsclient.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}
	mt := &MTCaptcha{
		SiteKey:   siteKey,
		HostName:  hostName,
		sessionID: sessionID,
		resetTS:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		Client:    client,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
	}
	mt.cookie = fmt.Sprintf("mtv1ConfSum={v:01|wdsz:std|thm:basic|lan:en|chlg:std|clan:1|cstyl:1|afv:0|afot:0|}; jsV=%s; mtv1Pulse=%s", VERSION, mt.GetPulseData())
	return mt, nil
}

func transactionSignature(value string) string {
	constant := "mtcap@mtcaptcha.com"
	hash := md5.Sum([]byte(constant + value))
	return "TH[" + hex.EncodeToString(hash[:]) + "]"
}

func (mt *MTCaptcha) GetChallenge() (GetChallengeRes, error) {
	params := []param{
		{"sk", mt.SiteKey},
		{"bd", mt.HostName},
		{"rt", strconv.FormatInt(time.Now().UnixMilli(), 10)},
		{"tsh", transactionSignature(mt.SiteKey)},
		{"act", "$"}, // _0x1f77ce.action || '$'
		{"ss", mt.sessionID},
		{"lf", "1"}, // textLength | either 0 or 1
		{"tl", "$"}, // _0x1f77ce.textLength != 0x0 ? _0x1f77ce.textLength : '$'
		{"lg", "en"},
		{"tp", "s"}, // _0x1f77ce.widgetSize == _0xdbc730.constant.standard ? 's' : 'm'
	}
	
	url := fmt.Sprintf("https://service.mtcaptcha.com/mtcv1/api/getchallenge.json?%s", encodeParams(params))
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return GetChallengeRes{}, err
	}
	
	req.Header = mt.getHeaders()
	
	resp, err := mt.Client.Do(req)
	if err != nil {
		return GetChallengeRes{}, err
	}
	defer resp.Body.Close()
	
	var res GetChallengeRes
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return GetChallengeRes{}, err
	}
	
	mt.challengeToken = res.Result.Challenge.Ct
	
	return res, nil
}

func (mt *MTCaptcha) GetImage() (GetImageRes, error) {
	params := []param{
		{"sk", mt.SiteKey},
		{"ct", mt.challengeToken},
		{"fa", "$"},
		{"ss", mt.sessionID},
	}
	
	url := fmt.Sprintf("https://service.mtcaptcha.com/mtcv1/api/getimage.json?%s", encodeParams(params))
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return GetImageRes{}, err
	}
	
	req.Header = mt.getHeaders()
	
	resp, err := mt.Client.Do(req)
	if err != nil {
		return GetImageRes{}, err
	}
	defer resp.Body.Close()
	
	var res GetImageRes
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return GetImageRes{}, err
	}
	
	return res, nil
}

func (mt *MTCaptcha) SolveChallenge(challenge GetChallengeRes, solution string) (GetChallengeRes, error) {
	fold := challenge.Result.Challenge.FoldChlg
	
	foldAnswer := "$"
	if challenge.Result.Challenge.HasFoldChlg {
		
		var err error
		foldAnswer, err = SolveFoldChallenge(fold.Fseed, fold.Fslots, fold.Fdepth)
		if err != nil {
			return GetChallengeRes{}, err
		}
	}
	params := []param{
		{"ct", mt.challengeToken},
		{"sk", mt.SiteKey},
		{"st", solution},
		{"lf", "1"},
		{"bd", mt.HostName},
		{"rt", strconv.FormatInt(time.Now().UnixMilli(), 10)},
		{"tsh", transactionSignature(mt.SiteKey)},
		{"fa", foldAnswer},
		{"qh", "$"}, //TODO look into automatedTestCode
		{"act", "$"},
		{"ss", mt.sessionID},
		{"tl", "$"}, // textLength
		{"lg", "en"},
		{"tp", "s"},
		{"kt", ""},
		{"fs", fold.Fseed},
	}
	
	url := fmt.Sprintf("https://service.mtcaptcha.com/mtcv1/api/getchallenge.json?%s", encodeParams(params))
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return GetChallengeRes{}, err
	}
	
	req.Header = mt.getHeaders()
	
	resp, err := mt.Client.Do(req)
	if err != nil {
		return GetChallengeRes{}, err
	}
	defer resp.Body.Close()
	
	var res GetChallengeRes
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return GetChallengeRes{}, err
	}
	
	mt.challengeToken = res.Result.Challenge.Ct
	
	return res, nil
}
