package gtRandom

import (
	"math/rand"
	"sync"
	"time"
)

const (
	KCharSetNum         = 1
	KCharSetUpperLetter = 1 << 1
	KCharSetLowLetter   = 1 << 2
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
func RandomScope(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano() + count())
	random := rand.New(seed)
	return min + random.Intn(max-min)
}

/*
* 生成指定范围内的随机数
* @return 生成的随机数范围[min, max)
**/
func RandomScopeInt64(min int64, max int64) int64 {
	seed := rand.NewSource(time.Now().UnixNano() + count())
	random := rand.New(seed)
	return min + random.Int63n(max-min)
}

func RandomString(length int, charSet int) string {
	if length <= 0 {
		return ""
	}

	numbers := []byte("0123456789")
	upperLetters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowLetters := []byte("abcdefghijklmnopqrstuvwxyz")

	charPool := make([]byte, 0)

	//数字
	if charSet&KCharSetNum == 1 {
		charPool = append(charPool, numbers...)
	}

	//大写字母
	if charSet&KCharSetUpperLetter == 1 {
		charPool = append(charPool, upperLetters...)
	}

	//小写字母
	if charSet&KCharSetLowLetter == 1 {
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
