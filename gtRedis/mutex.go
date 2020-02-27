package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

const srcLockScript = `
local result = redis.pcall('SET', KEYS[1], ARGV[1], "NX", "EX", ARGV[2])
if result then
    return 1
else
    return 0
end
`
const srcUnlockScript = `
if redis.pcall("GET", KEYS[1]) == ARGV[1] then
    return redis.pcall("DEL", KEYS[1])
else
    return 0
end
`

var lockScript *redis.Script
var unlockScript *redis.Script

func init() {
	lockScript = redis.NewScript(1, srcLockScript)
	unlockScript = redis.NewScript(1, srcUnlockScript)
}

type Mutex struct {
	key    string
	value  string
	expire int
}

func (param *Mutex) New(key string, expire int) {
	param.key = "gtRedisMutex:" + key
	param.value = "mutex_value"
	param.expire = expire
}

func (param *Mutex) Lock(conn redis.Conn) error {
	args := redis.Args{}
	args = args.Add(param.key)
	args = args.Add(param.value)
	args = args.Add(param.expire)
	result, err := redis.Int(lockScript.Do(conn, args...))
	if err != nil {
		return formatError(err, "do script failed. key: %s expire: %d", param.key, param.expire)
	}
	if result != 1 {
		return ErrLockExist
	}
	return nil
}

func (param *Mutex) Unlock(conn redis.Conn) error {
	args := redis.Args{}
	args = args.Add(param.key)
	args = args.Add(param.value)
	result, err := redis.Int(unlockScript.Do(conn, args...))
	if err != nil {
		return formatError(err, "do script failed. key: %s", param.key)
	}
	if result == 0 {
		return ErrLockNotExist
	}
	return nil
}
