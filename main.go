package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var banner string = `
┏┓  ┓        ┏┓   •      ┏┓┓ ┏           
┃┓┏┓┃┏┓┏┓┏┓  ┗┓╋┏┓┓┏┓┏┓  ┃┃┣┓╋┓┏┏┏┏┓╋┏┓┏┓
┗┛┗┛┗┗┻┛┗┗┫  ┗┛┗┛ ┗┛┗┗┫  ┗┛┗┛┛┗┻┛┗┗┻┗┗┛┛ 
          ┛           ┛                  

Golang String Obfuscator                  
version: 1.0.0
made by love: RUSTY J0y
`

var rng *rand.Rand

var symbols []string

func fillSymbols() {
	for i := 32; i <= 126; i++ {
		c := rune(i)
		if !(c >= 'A' && c <= 'Z') && !(c >= 'a' && c <= 'z') {
			symbols = append(symbols, string(c))
		}
	}
}

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	fillSymbols()
}

var words = map[string]string{
	"a1": "apple", "a2": "air", "a3": "animal", "a4": "adventure", "a5": "angel",
	"a6": "artist", "a7": "answer", "a8": "amazing", "a9": "autumn", "a10": "arrow",
	"a11": "angle", "a12": "ancient", "a13": "alarm", "a14": "affair", "a15": "agree",
	"a16": "attempt", "a17": "alert", "a18": "attend", "a19": "arrival", "a20": "attach",

	"b1": "banana", "b2": "book", "b3": "blue", "b4": "bridge", "b5": "bottle",
	"b6": "bread", "b7": "bird", "b8": "bag", "b9": "button", "b10": "butter",
	"b11": "bench", "b12": "bright", "b13": "black", "b14": "bark", "b15": "bubble",
	"b16": "begin", "b17": "believe", "b18": "breeze", "b19": "budget", "b20": "bitter",

	"c1": "cat", "c2": "car", "c3": "cake", "c4": "circle", "c5": "cloud",
	"c6": "chair", "c7": "child", "c8": "clock", "c9": "city", "c10": "computer",
	"c11": "candle", "c12": "crystal", "c13": "color", "c14": "cousin", "c15": "camera",
	"c16": "change", "c17": "climb", "c18": "corner", "c19": "chocolate", "c20": "chance",

	"d1": "dog", "d2": "dance", "d3": "dream", "d4": "door", "d5": "dragon",
	"d6": "desert", "d7": "diamond", "d8": "dress", "d9": "drive", "d10": "dust",
	"d11": "danger", "d12": "decision", "d13": "delight", "d14": "detail", "d15": "doctor",
	"d16": "defend", "d17": "doubt", "d18": "dinner", "d19": "desire", "d20": "discover",

	"e1": "elephant", "e2": "ear", "e3": "earth", "e4": "energy", "e5": "envelope",
	"e6": "engine", "e7": "event", "e8": "eagle", "e9": "empty", "e10": "excite",
	"e11": "effect", "e12": "elegant", "e13": "enemy", "e14": "escape", "e15": "exercise",
	"e16": "experience", "e17": "explore", "e18": "equal", "e19": "eventual", "e20": "edit",

	"f1": "fish", "f2": "flower", "f3": "forest", "f4": "friend", "f5": "family",
	"f6": "fire", "f7": "frame", "f8": "fruit", "f9": "famous", "f10": "factor",
	"f11": "fast", "f12": "field", "f13": "flood", "f14": "finish", "f15": "future",
	"f16": "figure", "f17": "focus", "f18": "fabric", "f19": "funny", "f20": "flight",

	"g1": "garden", "g2": "game", "g3": "gold", "g4": "giant", "g5": "ground",
	"g6": "glass", "g7": "green", "g8": "guide", "g9": "group", "g10": "guitar",
	"g11": "growth", "g12": "ghost", "g13": "glove", "g14": "grain", "g15": "gravity",
	"g16": "galaxy", "g17": "gain", "g18": "gesture", "g19": "garage", "g20": "glimpse",

	"h1": "house", "h2": "horse", "h3": "happy", "h4": "hill", "h5": "heart",
	"h6": "hand", "h7": "honey", "h8": "hotel", "h9": "holiday", "h10": "hunter",
	"h11": "honor", "h12": "humble", "h13": "hero", "h14": "hunger", "h15": "human",
	"h16": "habit", "h17": "hazard", "h18": "helmet", "h19": "history", "h20": "hope",

	"i1": "ice", "i2": "island", "i3": "idea", "i4": "image", "i5": "important",
	"i6": "inside", "i7": "interest", "i8": "increase", "i9": "invisible", "i10": "internet",
	"i11": "identity", "i12": "impact", "i13": "issue", "i14": "imagine", "i15": "item",
	"i16": "intense", "i17": "install", "i18": "invite", "i19": "income", "i20": "instance",

	"j1": "jungle", "j2": "juice", "j3": "jump", "j4": "jewel", "j5": "jacket",
	"j6": "journey", "j7": "joke", "j8": "join", "j9": "judge", "j10": "jazz",
	"j11": "jealous", "j12": "jog", "j13": "jolly", "j14": "job", "j15": "justice",
	"j16": "junior", "j17": "juggle", "j18": "joy", "j19": "journal", "j20": "junction",

	"k1": "king", "k2": "kite", "k3": "key", "k4": "kind", "k5": "kitchen",
	"k6": "knowledge", "k7": "knock", "k8": "knife", "k9": "kick", "k10": "keep",
	"k11": "kit", "k12": "knee", "k13": "kneel", "k14": "kingdom", "k15": "kettle",
	"k16": "keyboard", "k17": "kid", "k18": "keen", "k19": "knight", "k20": "karma",

	"l1": "lion", "l2": "lake", "l3": "light", "l4": "leaf", "l5": "love",
	"l6": "lamp", "l7": "long", "l8": "lucky", "l9": "listen", "l10": "learn",
	"l11": "limit", "l12": "leader", "l13": "letter", "l14": "launch", "l15": "level",
	"l16": "line", "l17": "language", "l18": "land", "l19": "layer", "l20": "library",

	"m1": "moon", "m2": "mountain", "m3": "music", "m4": "magic", "m5": "mirror",
	"m6": "machine", "m7": "morning", "m8": "market", "m9": "memory", "m10": "motion",
	"m11": "moment", "m12": "message", "m13": "mystery", "m14": "metal", "m15": "monkey",
	"m16": "medal", "m17": "museum", "m18": "medium", "m19": "mission", "m20": "model",

	"n1": "night", "n2": "nature", "n3": "name", "n4": "number", "n5": "nest",
	"n6": "needle", "n7": "note", "n8": "novel", "n9": "neutral", "n10": "noise",
	"n11": "narrow", "n12": "nation", "n13": "necessary", "n14": "noble", "n15": "near",
	"n16": "novice", "n17": "nasty", "n18": "notice", "n19": "neighbor", "n20": "november",

	"o1": "ocean", "o2": "orange", "o3": "open", "o4": "owl", "o5": "orbit",
	"o6": "object", "o7": "office", "o8": "offer", "o9": "order", "o10": "opinion",
	"o11": "outer", "o12": "owner", "o13": "outline", "o14": "output", "o15": "over",
	"o16": "opposite", "o17": "occasion", "o18": "observe", "o19": "option", "o20": "organic",

	"p1": "pink", "p2": "paper", "p3": "plant", "p4": "place", "p5": "planet",
	"p6": "people", "p7": "party", "p8": "picture", "p9": "pocket", "p10": "point",
	"p11": "popular", "p12": "puzzle", "p13": "pass", "p14": "peace", "p15": "profit",
	"p16": "practice", "p17": "protect", "p18": "process", "p19": "public", "p20": "pattern",

	"q1": "queen", "q2": "question", "q3": "quick", "q4": "quiet", "q5": "quote",
	"q6": "quality", "q7": "quest", "q8": "queue", "q9": "quilt", "q10": "quirk",
	"q11": "quack", "q12": "quake", "q13": "quarry", "q14": "qualify", "q15": "quench",
	"q16": "quart", "q17": "query", "q18": "quiver", "q19": "quota", "q20": "quoteable",

	"r1": "rain", "r2": "river", "r3": "road", "r4": "rose", "r5": "room",
	"r6": "railway", "r7": "rock", "r8": "race", "r9": "rainbow", "r10": "reach",
	"r11": "rule", "r12": "right", "r13": "reply", "r14": "ride", "r15": "rich",
	"r16": "round", "r17": "rare", "r18": "reveal", "r19": "relax", "r20": "repair",

	"s1": "sun", "s2": "star", "s3": "sky", "s4": "smile", "s5": "snake",
	"s6": "sugar", "s7": "strong", "s8": "sea", "s9": "stone", "s10": "speed",
	"s11": "shape", "s12": "shadow", "s13": "shine", "s14": "shop", "s15": "share",
	"s16": "song", "s17": "soft", "s18": "storm", "s19": "stick", "s20": "sweep",

	"t1": "tree", "t2": "time", "t3": "town", "t4": "table", "t5": "train",
	"t6": "task", "t7": "truth", "t8": "travel", "t9": "tower", "t10": "tiger",
	"t11": "track", "t12": "team", "t13": "taste", "t14": "tool", "t15": "thread",
	"t16": "trail", "t17": "tune", "t18": "tiny", "t19": "trade", "t20": "trend",

	"u1": "under", "u2": "umbrella", "u3": "universe", "u4": "unique", "u5": "unit",
	"u6": "update", "u7": "use", "u8": "urgent", "u9": "usual", "u10": "urban",
	"u11": "union", "u12": "up", "u13": "ultra", "u14": "upset", "u15": "utility",
	"u16": "uncle", "u17": "until", "u18": "unfold", "u19": "upgrade", "u20": "upward",

	"v1": "victory", "v2": "vision", "v3": "valley", "v4": "volunteer", "v5": "vehicle",
	"v6": "vacation", "v7": "value", "v8": "verify", "v9": "virtual", "v10": "venture",
	"v11": "vivid", "v12": "vessel", "v13": "vaccine", "v14": "vocal", "v15": "volume",
	"v16": "variety", "v17": "velocity", "v18": "visitor", "v19": "vote", "v20": "vortex",

	"w1": "water", "w2": "wind", "w3": "world", "w4": "wave", "w5": "wood",
	"w6": "wall", "w7": "wheel", "w8": "window", "w9": "white", "w10": "work",
	"w11": "wonder", "w12": "warm", "w13": "weight", "w14": "wish", "w15": "watch",
	"w16": "wild", "w17": "winter", "w18": "wolf", "w19": "west", "w20": "word",

	"x1": "xray", "x2": "xylophone", "x3": "xerox", "x4": "xenon", "x5": "xylem",
	"x6": "xenial", "x7": "xenolith", "x8": "xiphoid", "x9": "xenagogue", "x10": "xerothermic",
	"x11": "xanthic", "x12": "xenophile", "x13": "xystus", "x14": "xeriscape", "x15": "xylograph",
	"x16": "xeroderma", "x17": "xenograft", "x18": "xenogenesis", "x19": "xenium", "x20": "xenotransplant",

	"y1": "yellow", "y2": "yoga", "y3": "yard", "y4": "youth", "y5": "year",
	"y6": "yesterday", "y7": "yawn", "y8": "yet", "y9": "yield", "y10": "young",
	"y11": "yolk", "y12": "yell", "y13": "yes", "y14": "yearn", "y15": "yonder",
	"y16": "yacht", "y17": "yardstick", "y18": "yaw", "y19": "yip", "y20": "yodel",

	"z1": "zebra", "z2": "zero", "z3": "zone", "z4": "zoom", "z5": "zeal",
	"z6": "zinc", "z7": "zenith", "z8": "zesty", "z9": "zigzag", "z10": "zodiac",
	"z11": "zombie", "z12": "zephyr", "z13": "zen", "z14": "zip", "z15": "zillion",
	"z16": "zig", "z17": "zoo", "z18": "zealous", "z19": "zenana", "z20": "zap",
}

func rand20() int {
	return rng.Intn(20) + 1
}

func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func charToWord(ch string) string {
	chLower := strings.ToLower(ch)
	randomInt := rand20()
	randomStr := strconv.Itoa(randomInt)
	key := chLower + randomStr
	return key
}

func makeFirstUpper(word string) string {
	start := string(word[0])
	start = strings.ToUpper(start)
	ans := start + word[1:]
	return ans
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func WordEnc(arg string) []string {
	ans := make([]string, len(arg))

	for i := 0; i < len(arg); i++ {
		ch := string(arg[i])
		var myWord string
		if contains(symbols, ch) {
			myWord = ch
		} else {
			alreadyUpper := isUpper(ch)
			key := charToWord(ch)
			myWord = words[key]
			if alreadyUpper {
				myWord = makeFirstUpper(myWord)
			}
		}
		ans[i] = myWord
	}

	return ans
}

func MakeStrArray(arg string) {
	slice := WordEnc(arg)
	fmt.Print(`myObfArray := []string{`)
	for i := 0; i < len(slice); i++ {
		fmt.Print("`" + slice[i] + "`")
		if i != len(slice)-1 {
			fmt.Print(`,`)
		}
	}
	fmt.Print(`}`)
	fmt.Println()
}

func MakeOrginal(array []string) string {
	ans := ""
	for _, word := range array {
		ans += string(word[0])
	}
	return ans
}

func main() {
	fmt.Println(banner)
	fmt.Println()

	var joined string
	if len(os.Args) > 1 {
		args := os.Args[1:]
		joined = strings.Join(args, " ")
		//fmt.Println(joined)
	} else {
		fmt.Println("[-] Missing argument")
		return
	}

	fmt.Println()

	MakeStrArray(joined)

}
