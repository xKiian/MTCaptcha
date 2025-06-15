package mtcaptcha

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type MTCaptcha struct {
	SiteKey   string
	HostName  string
	SessionID string
}

func New(siteKey, hostName string) *MTCaptcha {
	sessionID := "S1" + uuid.New().String()
	return &MTCaptcha{
		SiteKey:   siteKey,
		HostName:  hostName,
		SessionID: sessionID,
	}
}

func transactionSignature(value string) string {
	constant := "mtcap@mtcaptcha.com"
	hash := md5.Sum([]byte(constant + value))
	return "TH[" + hex.EncodeToString(hash[:]) + "]"
}

func (mt *MTCaptcha) GetChallenge() {
	params := []param{
		{"sk", mt.SiteKey},
		{"bd", mt.HostName},
		{"rt", strconv.FormatInt(time.Now().UnixMilli(), 10)},
		{"tsh", transactionSignature(mt.SiteKey)},
		{"act", "$"}, // _0x1f77ce.action || '$'
		{"ss", mt.SessionID},
		{"lf", "1"}, // textLength | either 0 or 1
		{"tl", "$"}, // _0x1f77ce.textLength != 0x0 ? _0x1f77ce.textLength : '$'
		{"lg", "en"},
		{"tp", "s"}, // _0x1f77ce.widgetSize == _0xdbc730.constant.standard ? 's' : 'm'
	}
	
	fmt.Println(encodeParams(params))
}
