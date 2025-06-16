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
	
	pulseMap := make([]int, 13)
	r := []int{int(math.Floor(rand.Float64() * 4095.99)), int(math.Floor(rand.Float64() * 4095.99))}
	n := int(time.Now().Unix())
	a := strings.ToLower(fmt.Sprintf("%v", mt.UserAgent))
	d := strings.ToLower(fmt.Sprintf("%v", mt.getRef()))
	l := strings.ToLower(fmt.Sprintf("%v", "en-US"))
	
	pulseMap[0] = 0
	pulseMap[1] = 1
	pulseMap[2] = r[0]
	pulseMap[3] = r[1]
	
	pulseMap[4] = (n / 0x400000) % 0x1000
	pulseMap[5] = ((n % 0x400000) >> 0xb) % 0x1000
	pulseMap[6] = (n % 0x400000) % 0x800
	
	div := int(math.Floor(float64(-offset/600) / 10.0))
	pulseMap[7] = int(math.Abs(float64(div % 0x1000)))
	pulseMap[8] = int(math.Abs(float64(hashStr(a)))) % 0x1000
	
	re := regexp.MustCompile(`^(?:https?://)?(?:[^@/\n]+@)?(?:www\.)?([^:/\n]+)`)
	match := re.FindStringSubmatch(d)
	var domain string
	if len(match) > 1 {
		domain = match[1]
	}
	
	pulseMap[9] = int(math.Abs(float64(hashStr(domain)))) % 0x1000
	pulseMap[10] = int(math.Abs(float64(hashStr(l)))) % 0x1000
	pulseMap[11] = 10 % 0x1000
	
	var index int32 = 0
	for i := 0; i < 12; i++ {
		index = index*0x1f + int32(pulseMap[i])
		index = index & index
	}
	index = int32(math.Abs(float64(index)))
	pulseMap[12] = int(index) % 0x1000
	
	for i := 4; i < len(pulseMap); i++ {
		pulseMap[i] ^= r[i%2]
	}
	
	var _0x11126a []string
	for _, val := range pulseMap {
		_0x11126a = append(_0x11126a, URLSafeBase4096IntToChar(val))
	}
	
	return strings.Join(_0x11126a, "")
}
