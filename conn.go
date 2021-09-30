package netpool

import (
	"net"
	"sync/atomic"
	"time"
)

var noDeadline = time.Time{}

type Conn struct {
	usedAt    int64 // atomic
	netConn   net.Conn
	Inited    bool
	pooled    bool // 此链接是否需要返还
	createdAt time.Time
}

func newConn(netConn net.Conn) *Conn {
	cn := &Conn{
		netConn:   netConn,
		createdAt: time.Now(),
	}
	cn.SetUsedAt(time.Now())
	return cn
}

func (cn *Conn) SetUnPooled() {
	cn.pooled = false
}

func (cn *Conn) UsedAt() time.Time {
	unix := atomic.LoadInt64(&cn.usedAt)
	return time.Unix(unix, 0)
}

func (cn *Conn) SetUsedAt(tm time.Time) {
	atomic.StoreInt64(&cn.usedAt, tm.Unix())
}

func (cn *Conn) NetConn() net.Conn {
	return cn.netConn
}

func (cn *Conn) Close() error {
	return cn.netConn.Close()
}
