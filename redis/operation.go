package redis

import "github.com/garyburd/redigo/redis"

type Redis struct {
	pool *redis.Pool
}

func (this *Redis) SetLocker(key string, value string, expired int) (ok bool, err error) {
	c := this.pool.Get()
	defer c.Close()

	_, err = redis.String(c.Do("SET", key, value, "EX", expired, "NX"))
	if err != nil {
		if err == redis.ErrNil {
			return false, nil
		}
		return
	}

	return true, nil
}

func (this *Redis) DelValue(key string) (err error) {
	c := this.pool.Get()
	defer c.Close()

	_, err = c.Do("DEL", key)
	if err != nil {
		return
	}

	return
}