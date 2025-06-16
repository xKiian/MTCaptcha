package mtcaptcha

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var URLSafeBase64CharCode2IntMap = [256]int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1, -1, -1, -1, -1,
	-1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, -1, -1, -1, -1, 63,
	-1, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
}

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

func URLSafeBase64CharToInt(c rune) (int, error) {
	if int(c) >= 0 && int(c) < 256 {
		val := URLSafeBase64CharCode2IntMap[c%256]
		if val >= 0 {
			return val, nil
		}
	}
	return -1, errors.New("arg charCode must be within chars [a-zA-Z0-9:;]")
}

func hashStr(s string) int {
	h := 0
	for _, c := range s {
		h = (h << 5) - h + int(c)
		h = int(int32(h))
	}
	return h
}

func URLSafeBase64IntToChar(i int) string {
	if i < 0 || i > 63 {
		panic("arg i must be between 0 .. 63 inclusive")
	}
	return URLSafeBase64Int2CharMap[i%64]
}

func URLSafeBase64Str2IntArray(s string) ([]int, error) {
	res := make([]int, 0, len(s))
	for _, ch := range s {
		i, err := URLSafeBase64CharToInt(ch)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

func URLSafeBase4096IntToChar(i int) string {
	if i < 0 || i > 4095 {
		panic("arg i must be between 0 .. 4095 inclusive")
	}
	
	hi := URLSafeBase64IntToChar(i >> 6)
	lo := URLSafeBase64IntToChar(i & 63)
	
	return hi + lo
}

func foldBase64IntArray(a1 []int, foldCount int) []int {
	a2 := make([]int, len(a1))
	copy(a2, a1)
	reverse(a2)
	a3 := make([]int, len(a1))
	copy(a3, a1)
	
	offset := 0
	x := 0
	y := 0
	z := 0
	for i := 0; i < foldCount; i++ {
		offset++
		for x = 0; x < len(a1); x++ {
			a3[x] = (int(math.Floor(float64(a3[x]+a2[(x+offset)%len(a2)])*73/8)) + y + z) % 64
			z = y / 2
			y = a3[x] / 2
		}
	}
	return a3
}

func hashIntAry(a []int) int {
	h := 0
	for _, v := range a {
		h = (h << 5) - h + v
		h = int(int32(h))
	}
	if h < 0 {
		h *= -1
	}
	return h
}

func SolveFoldChallenge(str string, depth, xor int) (string, error) {
	if str == "" || depth < 1 {
		return "0", nil
	}
	
	result := make([]string, 0, depth)
	
	arr, err := URLSafeBase64Str2IntArray(str)
	if err != nil {
		return "", err
	}
	
	for i := 0; i < depth; i++ {
		arr = foldBase64IntArray(arr, 31)
		hashed := hashIntAry(foldBase64IntArray(arr, xor)) % 4096
		result = append(result, URLSafeBase4096IntToChar(hashed))
	}
	
	return strings.Join(result, ""), nil
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
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
