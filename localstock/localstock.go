package localstock

import (
	"fmt"
)

type Localstock struct {
	//本地的库存
	LocalInstock int64
	//销售额
	LocalSalesVolume int64
}

//减少库存的操作
func (stock *Localstock) LocalDeductionStock() bool {
	stock.LocalSalesVolume = stock.LocalSalesVolume + 1
	return stock.LocalSalesVolume <= stock.LocalInstock
}

func Ttt() {
	fmt.Println("I love HCX")
}
