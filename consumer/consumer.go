package consumer

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	//gorm.Model
	Id   int `gorm:"primary_key"`
	Name string
	Time time.Time
}

func Createtable() {
	db, _ := gorm.Open("mysql", "root:121456aA@/yanjianyu?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()
	db.CreateTable(&User{})
}

func WriteDB(u1 User) {
	db, _ := gorm.Open("mysql", "root:121456aA@/yanjianyu?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()
	db.NewRecord(u1)
	db.Create(&u1)

}

//func main() {

//	db, _ := gorm.Open("mysql", "root:121456aA@/yanjianyu?charset=utf8&parseTime=True&loc=Local")
//	defer db.Close()
//db.CreateTable(&User{})

//	u1 := User{Name: "Bill", Time: time.Now()}

//	db.NewRecord(u1)

//	db.Create(&u1)

//	fmt.Println("Hello ")
//}
