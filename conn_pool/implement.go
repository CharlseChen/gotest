package conn_pool

import (
	"errors"
	"fmt"
	"time"
	"database/sql"
	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
)

func NewChannelPool(poolConfig *PoolConfig) (*channelPool, error) {
	if !(poolConfig.InitialCap <= poolConfig.MaxIdle && poolConfig.MaxCap >= poolConfig.MaxIdle && poolConfig.InitialCap >= 0) {
		return nil, errors.New("invalid capacity settings")
	}
	if poolConfig.Factory == nil {
		return nil, errors.New("invalid factory settings")
	}

	c := &channelPool{
		conns:        make(chan *idleConn, poolConfig.MaxIdle),
		factory:      poolConfig.Factory,
		idleTimeout:  poolConfig.IdleTimeout,
		maxActive:    poolConfig.MaxCap,
		openingConns: poolConfig.InitialCap,
	}

	for i := 0; i < poolConfig.InitialCap; i++ {
		conn, err := c.factory.Factory()
		if err != nil {
			c.Release()
			return nil, fmt.Errorf("err:%v", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}
	return c, nil
}

// getConns 获取所有连接
func (c *channelPool) getConns() chan *idleConn {
	c.mutex.Lock()
	conns := c.conns
	c.mutex.Unlock()
	return conns
}

func (c *channelPool) Get() (interface{}, error) {
	conns := c.getConns()
	if conns == nil {
		return nil, errors.New("conns is nil")
	}

	for {
		select {
		case wrapConn := <-conns:
			if wrapConn == nil {
				return nil, errors.New("wrap conn is nil")
			}

			if timeout := c.idleTimeout; timeout > 0 && wrapConn.t.Before(time.Now()) {
				_ = c.Close(wrapConn.conn)
				continue
			}
			if err := c.Ping(wrapConn.conn); err != nil {
				_ = c.Close(wrapConn.conn)
				continue
			}
			return wrapConn.conn, nil
		default:
			c.mutex.Lock()
			if c.openingConns >= c.maxActive {
				req := make(chan connReq, 1)
				c.connReqs = append(c.connReqs, req)
				c.mutex.Unlock()
				ret, ok := <-req
				if !ok {
					return nil, errors.New("conn is closed")
				}
				if timeout := c.idleTimeout; timeout > 0 && ret.idleConn.t.Add(timeout).Before(time.Now()) {
					//_ = c.Close(ret.idleConn.conn)
					continue
				}
				return ret.idleConn.conn, nil
			}
			if c.factory == nil {
				c.mutex.Unlock()
				return nil, errors.New("conn is closed")
			}
			conn, err := c.factory.Factory()
			if err != nil {
				c.mutex.Unlock()
				return nil, err
			}
			c.openingConns++
			c.mutex.Unlock()
			return conn, nil
		}
	}
}
func (c *channelPool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("conn is nil")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.conns == nil {
		return c.Close(conn)
	}
	if l := len(c.connReqs); l > 0 {
		req := c.connReqs[0]
		copy(c.connReqs, c.connReqs[1:])
		c.connReqs = c.connReqs[:l-1]
		req <- connReq{idleConn: &idleConn{conn: conn, t: time.Now()}}
		return nil
	}
	select {
	case c.conns <- &idleConn{conn: conn, t: time.Now()}:
		return nil
	default:
		return c.Close(conn)
	}
}

func (c *channelPool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("conn is nil")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.openingConns--
	fmt.Println("test dafa", conn)
	return c.factory.Close(conn)
}

func (c *channelPool) Release() {
	c.mutex.Lock()
	conns := c.conns
	c.conns = nil
	c.mutex.Unlock()

	defer func() {
		c.factory = nil
	}()

	if conns == nil {
		return
	}

	close(conns)
	for wrapConn := range conns {
		_ = c.factory.Close(wrapConn)
	}
}

func (c *channelPool) Len() int { return len(c.conns) }

func (c *channelPool) Ping(conn interface{}) error {
	if conn == nil {
		return errors.New("conn is nil")
	}
	return c.factory.Ping(conn)
}

type DbClient struct {
	d *sql.DB
}

func (d DbClient) Close(conn interface{}) error {

	db, _ := conn.(*DbClient)
	return db.d.Close()
}

func (d DbClient) Ping(conn interface{}) error {
	db, _ := conn.(*DbClient)
	return db.d.Ping()
}

func (d DbClient) Factory() (interface{}, error) {
	viper.SetConfigFile("config/config.toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s", err))
		return nil, err
	}
	links := viper.Get("database.default")
	var link string
	if items, ok := links.([]interface{}); ok && len(items) > 0 {
		firstItem := items[0].(map[string]interface{})
		link = firstItem["link"].(string)
	}
	db, err := sql.Open("mysql", link)
	if err != nil || db == nil {
		fmt.Printf("link:%v err:%v\n", link, err)
		return nil, err
	}
	return &DbClient{d: db}, nil
}
