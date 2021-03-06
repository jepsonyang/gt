package gtRandom

import (
	"math/rand"
	"sync"
	"time"
)

const (
	KCharSetNum         = 1				//数字
	KCharSetUpperLetter = 1 << 1		//大写字母
	KCharSetLowLetter   = 1 << 2		//小写字母
)

var counter int64
var mutex sync.Mutex

func count() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	return counter
}

/*
* 生成指定范围内的随机数
* @return 生成的随机数范围[min, max)
**/
func ScopeInt(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano() + count())
	random := rand.New(seed)
	return min + random.Intn(max-min)
}

/*
* 生成指定范围内的随机数
* @return 生成的随机数范围[min, max)
**/
func ScopeInt64(min int64, max int64) int64 {
	seed := rand.NewSource(time.Now().UnixNano() + count())
	random := rand.New(seed)
	return min + random.Int63n(max-min)
}

/*
* 生成随机字符串
* @length 生成字符串的长度
* @charSet 随机字符可能值集合，可使用KCharSetNum\KCharSetUpperLetter\KCharSetLowLetter按位与
* @return 生成的随机字符串
**/
func String(length int, charSet int) string {
	if length <= 0 {
		return ""
	}

	numbers := []byte("0123456789")
	upperLetters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowLetters := []byte("abcdefghijklmnopqrstuvwxyz")

	charPool := make([]byte, 0)

	//数字
	if charSet&KCharSetNum != 0 {
		charPool = append(charPool, numbers...)
	}

	//大写字母
	if charSet&KCharSetUpperLetter != 0 {
		charPool = append(charPool, upperLetters...)
	}

	//小写字母
	if charSet&KCharSetLowLetter != 0 {
		charPool = append(charPool, lowLetters...)
	}

	lenPool := len(charPool)
	if lenPool <= 0 {
		return ""
	}

	seed := rand.NewSource(time.Now().UnixNano() + count())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charPool[random.Intn(lenPool)]
	}

	return string(result)
}
