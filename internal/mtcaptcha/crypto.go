package mtcaptcha

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var URLSafeBase64Int2CharMap = []string{
	"0", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", "A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V",
	"W", "X", "Y", "Z", "a", "b", "c", "d",
	"e", "f", "g", "h", "i", "j", "k", "l",
	"m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z", "-", "_",
}

func hashStr(s string) int32 {
	var h int32 = 0
	for _, c := range s {
		h = (h << 5) - h + int32(c)
		h = h & h // force overflow
	}
	return h
}

func URLSafeBase64IntToChar(i int) string {
	if i < 0 || i > 63 {
		panic("arg i must be between 0 .. 63 inclusive")
	}
	return URLSafeBase64Int2CharMap[i%64]
}

func URLSafeBase4096IntToChar(i int) string {
	if i < 0 || i > 4095 {
		panic("arg i must be between 0 .. 4095 inclusive")
	}
	return URLSafeBase64IntToChar(i>>6) + URLSafeBase64IntToChar(i&63)
}

func (mt *MTCaptcha) GetPulseData() string {
	
	_, offset := time.Now().Zone()
	input := map[string]interface{}{
		"v": []int{0, 1},
		"r": []int{int(math.Floor(rand.Float64() * 4095.99)), int(math.Floor(rand.Float64() * 4095.99))},
		"n": int(time.Now().Unix()),
		"z": -offset / 600,
		"a": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
		"c": mt.cookie,
		"d": mt.getRef(),
		"l": "en-US",
		"h": 10,
	}
	
	_0x1fc26c := make([]int, 13)
	v := input["v"].([]int)
	r := input["r"].([]int)
	n := input["n"].(int)
	z := input["z"].(int)
	a := strings.ToLower(fmt.Sprintf("%v", input["a"]))
	d := strings.ToLower(fmt.Sprintf("%v", input["d"]))
	l := strings.ToLower(fmt.Sprintf("%v", input["l"]))
	h := input["h"].(int)
	
	_0x1fc26c[0] = v[0]
	_0x1fc26c[1] = v[1]
	_0x1fc26c[2] = r[0]
	_0x1fc26c[3] = r[1]
	
	_0xfa34f3 := (n / 0x400000) % 0x1000
	_0x172215 := ((n % 0x400000) >> 0xb) % 0x1000
	_0x8201d4 := (n % 0x400000) % 0x800
	
	_0x1fc26c[4] = _0xfa34f3
	_0x1fc26c[5] = _0x172215
	_0x1fc26c[6] = _0x8201d4
	
	div := int(math.Floor(float64(z) / 10.0))
	_0x1fc26c[7] = int(math.Abs(float64(div % 0x1000)))
	_0x1fc26c[8] = int(math.Abs(float64(hashStr(a)))) % 0x1000
	
	re := regexp.MustCompile(`^(?:https?://)?(?:[^@/\n]+@)?(?:www\.)?([^:/\n]+)`)
	match := re.FindStringSubmatch(d)
	var domain string
	if len(match) > 1 {
		domain = match[1]
	}
	_0x1fc26c[9] = int(math.Abs(float64(hashStr(domain)))) % 0x1000
	_0x1fc26c[10] = int(math.Abs(float64(hashStr(l)))) % 0x1000
	_0x1fc26c[11] = h % 0x1000
	
	var _0xaddae8 int32 = 0
	for i := 0; i < 12; i++ {
		_0xaddae8 = _0xaddae8*0x1f + int32(_0x1fc26c[i])
		_0xaddae8 = _0xaddae8 & _0xaddae8
	}
	_0xaddae8 = int32(math.Abs(float64(_0xaddae8)))
	_0x1fc26c[12] = int(_0xaddae8) % 0x1000
	
	for i := 4; i < len(_0x1fc26c); i++ {
		_0x1fc26c[i] ^= r[i%2]
	}
	
	var _0x11126a []string
	for _, val := range _0x1fc26c {
		_0x11126a = append(_0x11126a, URLSafeBase4096IntToChar(val))
	}
	
	return strings.Join(_0x11126a, "")
}
