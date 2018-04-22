package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Department
// カラム名の先頭は大文字にすること
type Department struct {
	gorm.Model
	DNo    string `gorm:"size:255"`
	DName  string `gorm:"size:255"`
	Budget int
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Department{})

	// Create
	db.Create(&Department{DNo: "D1", DName: "Marketing", Budget: 10000})

	// Read
	var dept Department
	db.First(&dept, 1)
	db.First(&dept, "d_no = ?", "D1")

	// Update
	db.Model(&dept).Update("Budget", 20000)

	// Delete
	db.Delete(&dept)
}
