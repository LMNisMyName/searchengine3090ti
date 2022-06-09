package idgen

import (
	"fmt"
	"searchengine3090ti/pkg/constants"
	"sync"
	"time"
)

//雪花算法生成id
//本地生成全局唯一id,支持高并发
type Snowflake struct {
	sync.Mutex       // 锁
	timestamp  int64 // 时间戳 ，毫秒
	workerid   int64 // 机器编号
	sequence   int64 // 序列号
}

const (
	epoch          = int64(1577808000000)
	timestampBits  = uint(41)
	workeridBits   = uint(10)
	sequenceBits   = uint(12)
	timestampMax   = int64(-1 ^ (-1 << timestampBits))
	workeridMax    = int64(-1 ^ (-1 << workeridBits))
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits))
	workeridShift  = sequenceBits
	timestampShift = sequenceBits + workeridBits
)

var s *Snowflake

func Init() {
	s = new(Snowflake)
	s.sequence = epoch
	s.workerid = constants.WorkerID
}

func GetID() int64 {
	s.Lock()
	now := time.Now().UnixNano() / 1000000
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		s.Unlock()
		fmt.Println("epoch must be between 0 and ", timestampMax-1)
		return 0
	}
	s.timestamp = now
	r := int64((t)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))
	s.Unlock()
	return r
}

// 自增算法产生ID
// 问题：已删除过的ID无法复用，不支持并发, 已弃用
// var ID int64

// func Init() {
// 	ret, err := rpc.SearchClient.QueryIDNumber(context.Background(), &searchapi.QueryIDNumberRequest{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	ID = ret.Number
// 	fmt.Println("Current records number: ", ID)
// }

// func GetID() int64 {
// 	ID++
// 	return ID
// }
