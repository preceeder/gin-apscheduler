//   File Name:  init.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/13 15:52
//    Change Activity:

package Init

func Init() {
	rc := RedisConfig{
		Host: "127.0.0.1",
		//Port: "6379",
		Port:        "56379",
		Password:    "QDjk9UdkoD6cv",
		Db:          0,
		MaxIdle:     2,
		IdleTimeout: 240,
		PoolSize:    10,
	}

	_ = initRedis(rc)
	initAps()
	Scheduler.SetDistributed(RedisClient, "")
}
