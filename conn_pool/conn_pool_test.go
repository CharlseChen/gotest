package conn_pool

import (
	"time"
	"testing"
	"sync"
)

func Test_Connect_Pool(t *testing.T) {
	pc := &PoolConfig{
		InitialCap:  int(10),
		MaxCap:      int(10),
		MaxIdle:     int(10),
		Factory:     &DbClient{}, // producerFactory 实现了工厂的接口 底层为创建 tcp 连接
		IdleTimeout: time.Second * 5,
	}

	pool, err := NewChannelPool(pc)
	if err != nil {
		t.Errorf("err:%v\n", err)
		return
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db, _ := pool.Get()
			t.Logf("conn pool get success db:%v\n", db)
			connect := db.(*DbClient)
			_, err = connect.d.Query("select * from xs_user_profile limit 1")
			err = pool.Put(db)
			err = pool.Close(connect)
		}()
	}
	wg.Wait()
}
