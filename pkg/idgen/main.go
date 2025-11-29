package idgen

import (
	"sync"
	"time"
)

const maxMessageSequence = 1000 // 单秒1000条消息.

// TimestampAutoID 带时间戳前缀的自增 ID 生成器
type TimestampAutoID struct {
	nodeId  int64
	mu      sync.Mutex
	seq     int64 // 秒级自增序列（0-999）
	lastSec int64 // 上一次生成 ID 的秒级时间戳
}

// NewTimestampAutoID 初始化生成器
func NewTimestampAutoID(nodeId int64) *TimestampAutoID {
	return &TimestampAutoID{
		nodeId: nodeId,
	}
}

// Next 获取下一个 ID（格式：时间戳（秒）* 1000 + 序列，共 16 位左右）
func (t *TimestampAutoID) Next() int64 {
	t.mu.Lock()
	defer t.mu.Unlock()

	nowSec := time.Now().Unix() // 秒级时间戳（避免毫秒/微秒导致 ID 过长）

	// 同一秒内：序列自增；跨秒：序列重置为 0
	if nowSec == t.lastSec {
		t.seq++
		if t.seq > maxMessageSequence {
			time.Sleep(time.Second - time.Duration(time.Now().Nanosecond())/time.Nanosecond)
			nowSec = time.Now().Unix()
			t.seq = 0
		}
	} else {
		t.seq = 0
	}
	t.lastSec = nowSec
	return nowSec*(maxMessageSequence)%10000000000000 + 10000000000000*t.nodeId + t.seq
}
