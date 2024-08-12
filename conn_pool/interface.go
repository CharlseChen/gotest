package conn_pool

import (
	"sync"
	"time"
)

type IConnPool interface {
	// 获取资源
	Get() (interface{}, error)
	// 放回资源
	Put(interface{}) error
	// 关闭资源
	Close(interface{}) error
	// 释放所有资源
	Release()
	// 返回当前池子内有效连接数量
	Len() int
}

type ConnectionFactory interface {
	//生成连接的方法
	Factory() (interface{}, error)
	//关闭连接的方法
	Close(interface{}) error
	//检查连接是否有效的方法
	Ping(interface{}) error
}

type channelPool struct {
	mutex        sync.Mutex
	conns        chan *idleConn
	factory      ConnectionFactory
	idleTimeout  time.Duration
	waitTimeout  time.Duration
	maxActive    int
	openingConns int
	connReqs     []chan connReq
}

type idleConn struct {
	conn interface{}
	t    time.Time
}

type connReq struct {
	idleConn *idleConn
}

type PoolConfig struct {
	InitialCap  int
	MaxCap      int
	MaxIdle     int
	Factory     ConnectionFactory
	IdleTimeout time.Duration
}
