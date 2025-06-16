package mtcaptcha

import (
	"strings"
	"time"
)

type Kee struct {
	TailKeeint  int
	lastKeeTS   int64
	ciderb64int []int
	keehist     []int
	keehistPos  int
	prevString  string
	initTS      int64
}

func NewKee(input string) *Kee {
	kee := &Kee{
		TailKeeint: 63,
		lastKeeTS:  -1,
		prevString: "",
		initTS:     time.Now().UnixMilli(),
	}
	kee.Init(input)
	return kee
}

func (k *Kee) Init(input string) {
	if input == "" || len(input) < 64 {
		return
	}
	if len(input) >= 64 {
		input = input[:64]
	}
	arr := make([]int, len(input))
	lastChar := URLSafeBase64CharCode2IntMap[input[len(input)-1]]
	
	for i := 0; i < len(input); i++ {
		char := input[i]
		code := URLSafeBase64CharCode2IntMap[char]
		arr[i] = code ^ lastChar
		lastChar = code
	}
	
	k.ciderb64int = arr
	k.keehist = append([]int{}, arr...)
	k.keehistPos = 0
}

func (k *Kee) GetKey(s string) string {
	result := ""
	prev := k.prevString
	if s != "" && len(prev) <= len(s) {
		for i := range s {
			if i >= len(prev) || s[i] != prev[i] {
				result += string(s[i])
			}
		}
		k.prevString = s
		if len(result) > 0 {
			return string(result[0])
		}
		return ""
	}
	k.prevString = s
	return "Backspace"
}

func (k *Kee) Play(value string) bool {
	if k.ciderb64int == nil {
		return false
	}
	key := k.GetKey(value)
	if k.keehistPos == 0 {
		elapsed := int((time.Now().UnixMilli() - k.initTS) / 500)
		if elapsed > 3900 {
			elapsed = 3900
		}
		a := elapsed / 64
		b := elapsed % 64
		
		k.keehist[0] = k.ciderb64int[0] ^ a
		k.keehist[1] = k.ciderb64int[1] ^ b
		k.keehist[2] = k.ciderb64int[2] ^ k.TailKeeint
		k.keehistPos += 2
	}
	
	if key == "-" || key == string(byte(45)) {
		return false
	}
	if key == "Backspace" || key == string(byte(8)) {
		key = "-"
	}
	if len(key) > 1 {
		return false
	}
	var code int
	if len(key) == 1 {
		code = int(key[0])
	} else {
		code = -1
	}
	mapVal := URLSafeBase64CharCode2IntMap[byte(code)]
	
	now := time.Now().UnixMilli()
	delay := int((now - k.lastKeeTS) / 30)
	if k.lastKeeTS < 0 {
		delay = 0
	}
	if delay > 63 {
		delay = 63
	}
	
	pos := k.keehistPos
	k.keehist[pos] = k.ciderb64int[pos] ^ mapVal
	k.keehist[pos+1] = k.ciderb64int[pos+1] ^ delay
	k.keehist[pos+2] = k.ciderb64int[pos+2] ^ k.TailKeeint
	k.keehistPos = pos + 2
	k.lastKeeTS = now
	return true
}

func (k *Kee) Get() string {
	result := make([]string, len(k.keehist))
	for i, v := range k.keehist {
		result[i] = URLSafeBase64Int2CharMap[v]
	}
	return strings.Join(result, "")
}

func GenerateKee(input, solution string) string {
	k := NewKee(input)
	for i := 1; i <= len(solution); i++ {
		k.Play(solution[:i])
	}
	return k.Get()
}
