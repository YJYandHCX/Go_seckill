package main

import (
	"fmt"
	//"seckill/localstock"
	"seckill/consumer"
	"time"
)

func main() {
	u1 := consumer.User{Name: "Jack", Time: time.Now()}
	consumer.WriteDB(u1)
	fmt.Println("Hello")
}
