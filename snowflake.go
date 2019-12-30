package snowflake

import (
	"sync/atomic"
	"time"
)

const (
	Mask = 4294959104
)

type Snowflake struct {
	clusterId uint8
	workerId  uint8
	seqId     uint32
	tms       int64
}

// NewNode create a new node
func NewNode(clusterId, workerId uint8) *Snowflake {
	return &Snowflake{clusterId, workerId, 0, 0}
}

// NextId returns next id
func (s *Snowflake) NextId() int64 {
	atomic.AddUint32(&s.seqId, 1)
	atomic.StoreInt64(&s.tms, time.Now().UnixNano()/1000>>1)

	return s.generate()
}

func (s *Snowflake) generate() int64 {
	return s.tms<<22 | int64(s.clusterId)<<17 | int64(s.workerId)<<12 | int64(s.seqId)
}
