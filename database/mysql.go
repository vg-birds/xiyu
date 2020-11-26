package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var SqlDB * gorm.DB

func init()  {
	var err error
	SqlDB, err = gorm.Open("mysql","root:root@/test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
}