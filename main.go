package main

import (
	"fmt"
	"net/http"
	"seckill/localstock"
	"seckill/remotestock"
	"seckill/util"

	"github.com/garyburd/redigo/redis"
	//"seckill/mypackage" // 导入同一项目下的mypackage包
)

var (
	//本地库存
	local_stock localstock.Localstock

	//链接池
	pool *redis.Pool

	//链接
	rc redis.Conn

	remote_stock remotestock.RemoteStock
)

func handleReq(w http.ResponseWriter, r *http.Request) {
	//Conn := pool.Get()
	//LogMsq := ""

	if local_stock.LocalDeductionStock() && remotestock.RemoteDeductionStock(rc) {
		util.Resp(w, 1, "抢购成功", nil)
	} else {
		util.Resp(w, -1, "失败", nil)
	}
}
func main() {
	local_stock.LocalInstock = 10000

	local_stock.LocalSalesVolume = 0

	pool = remotestock.NewPool()
	rc = pool.Get()
	http.HandleFunc("/", handleReq)
	http.ListenAndServe(":3005", nil)
	//rc = pool.Get()
	fmt.Println("wefwef")
	///remotestock.RemoteDeductionStock(rc)
	//remotestock.RemoteDeductionStock(rc)
	//fmt.Println("main")
}
