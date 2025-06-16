package mtcaptcha

import (
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"net/url"
	"strings"
)

const VERSION = "2024-11-14.21.25.03"

func encodeParams(params []param) string {
	var b strings.Builder
	for i, p := range params {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(url.QueryEscape(p.Key))
		b.WriteByte('=')
		b.WriteString(url.QueryEscape(p.Value))
	}
	
	return b.String()
}

func (mt *MTCaptcha) getRef() string {
	return fmt.Sprintf("https://service.mtcaptcha.com/mtcv1/client/iframe.html?v=%s&sitekey=%s&iframeId=mtcaptcha-iframe-1&widgetSize=standard&custom=false&widgetInstance=mtcaptcha&challengeType=standard&theme=basic&lang=en&action=&autoFadeOuterText=false&host=%s&hostname=%s&serviceDomain=service.mtcaptcha.com&textLength=0&lowFrictionInvisible=&enableMouseFlow=false&resetTS=%s",
		VERSION, mt.SiteKey, url.QueryEscape("https://"+mt.HostName), url.QueryEscape(mt.HostName), mt.resetTS,
	)
}

func (mt *MTCaptcha) getHeaders() http.Header {
	return http.Header{
		"accept":          {"*/*"},
		"accept-encoding": {"gzip, deflate, br, zstd"},
		"accept-language": {"en-US,en;q=0.9"},
		"referer":         {mt.getRef()},
		"cookie":          {mt.cookie},
		"priority":        {"u=1, i"},
		//"sec-ch-ua":          {`"Google Chrome";v="137", "Chromium";v="137", "Not/A)Brand";v="24"`},
		"sec-ch-ua-mobile":         {"?0"},
		"sec-ch-ua-platform":       {`"Windows"`},
		"sec-fetch-dest":           {"empty"},
		"sec-fetch-mode":           {"cors"},
		"sec-fetch-site":           {"same-origin"},
		"sec-fetch-storage-access": {"active"},
		"user-agent":               {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"accept-encoding",
			"referer",
			"cookie",
			"priority",
			//"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"sec-fetch-dest",
			"sec-fetch-mode",
			"sec-fetch-site",
			"sec-fetch-storage-access",
			"user-agent",
		},
	}
}
