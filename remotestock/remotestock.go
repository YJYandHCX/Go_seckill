//package remotestock
package remotestock

import (
	"github.com/garyburd/redigo/redis"
)

type RemoteStock struct {
	Stock_Num string
	Total_Num string
	Order_Num string
}

//这里是给新的连接池
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   10000,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

//var (
//	pool *redis.Pool
//	rc   redis.Conn
//)

//这里是远程减少redis中间的内存
func RemoteDeductionStock(conn redis.Conn) bool {
	var rest int
	for {
		b, _ := redis.Int(conn.Do("setnx", "ttt", "sss"))
		if b == 1 {
			break
		}
	}

	rest, _ = redis.Int(conn.Do("DECR", "stock"))
	//fmt.Println(rest)
	conn.Do("DEL", "ttt")
	if rest >= 0 {
		return true
	}
	return false

}

//func main() {
//	pool = NewPool()
//	rc = pool.Get()0
//	var i int = 1
//	for {
//		if i >= 1 {
//			break
//		}
//	}
//	RemoteDeductionStock(rc)
//	fmt.Println("wfewefwefew")
//}
