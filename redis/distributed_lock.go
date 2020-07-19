package redis


const (
	Distributed_Lock = "distributed_lock"
	SETNX_EXPIRED_TIME    = 60 * 60 * 24
)

func DistributedLock(Function func()){
	redisConnection := NewRedis()
	if ok, _ := redisConnection.SetLocker(Distributed_Lock, Distributed_Lock, SETNX_EXPIRED_TIME); ok {
		Function()
	}

	//unlock
	redisConnection.DelValue(Distributed_Lock)
}