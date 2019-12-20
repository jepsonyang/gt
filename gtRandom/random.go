package gtRandom

import (
	"math/rand"
	"sync"
	"time"
)

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

var counter int64
var mutex sync.Mutex

func count() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	return counter
}
